package session_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cholazzzb/amaz_corp_be/internal/app/repository/session"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/session"
)

func TestGetSessionById(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sid, err := msr.GetSessionById(context.Background(), ent.SessionId("not-found"))
	assert.Error(t, err, "when not found a session should return error")
	assert.Empty(t, sid.Id, "when not found should return nil session Id")
}

func TestCreateSession(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sessionId, err := msr.CreateSession(context.Background(), time.Now(), time.Now().Add(60*time.Minute))
	assert.Empty(t, err, "create new session should not return error")
	assert.Greater(t, len(sessionId), 0, "create new session should return uuid length greater than 0")
	ses, err := msr.GetSessionById(context.Background(), sessionId)
	assert.Empty(t, err, "should not return error when the sessionId is created")
	assert.Equal(t, sessionId, ses.Id, "should return the same sessionId from getSessionById")
}

func TestUpdateSessionStatus(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sId, _ := msr.CreateSession(
		context.Background(),
		time.Now(), time.Now().Add(time.Hour),
	)

	err := msr.UpdateSessionStatus(
		context.Background(),
		sId,
		ent.SessionStatus(ent.Invited),
	)
	assert.Empty(t, err, "update session status should not return error")

	ses, err := msr.GetSessionById(context.Background(), sId)
	assert.Empty(t, err, "should not return error when update success")
	assert.Equal(t, ent.SessionStatus(ent.Invited), ses.Status, "the status should be updated to Invited")
}

// Room
func TestGetRoomById(t *testing.T) {
	msr := session.NewMockSessionRepo()
	room, err := msr.GetRoomById(context.Background(), ent.RoomId("not found"))
	assert.Empty(t, room, "should return empty room, because there are no room created")
	assert.Error(t, err, "should return error, because there are no room created")
}

func TestGetRoomsBySessionId(t *testing.T) {
	msr := session.NewMockSessionRepo()
	_, err := msr.CreateSession(context.Background(), time.Now(), time.Now().Add(time.Hour))
	assert.Empty(t, err, "create session should not return error")
}

func TestCreateRoom(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sessionId, err := msr.CreateSession(context.Background(), time.Now(), time.Now().Add(time.Hour))
	assert.Empty(t, err, "create session should not return error")

	roomId1, err := msr.CreateRoom(context.Background(), sessionId, "room 1")
	assert.Empty(t, err, "create room should not return error")
	room1, err := msr.GetRoomById(context.Background(), roomId1)
	assert.NotEmpty(t, room1, "should return room")
	assert.Empty(t, err, "should not return error")
	rooms, _ := msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))
	assert.Equal(t, 1, len(rooms), "should return 1 room in session 1")

	roomId2, err := msr.CreateRoom(context.Background(), sessionId, "room 2")
	assert.Empty(t, err, "create room should not return error")
	room2, err := msr.GetRoomById(context.Background(), roomId2)
	assert.NotEmpty(t, room2, "should return room")
	assert.Empty(t, err, "should not return error")
	rooms, _ = msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))
	assert.Equal(t, 2, len(rooms), "should return 2 rooms in session 1")
}

func TestAddMember(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sessionId, err := msr.CreateSession(context.Background(), time.Now(), time.Now().Add(time.Hour))
	assert.Empty(t, err, "create session should not return error")

	msr.CreateRoom(context.Background(), sessionId, "room 1")
	rooms, err := msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))
	assert.Empty(t, err, "create room should not return error")
	assert.Equal(t, 1, len(rooms), "should return 1 room in 1 session")

	mId, err := msr.AddMember(
		context.Background(),
		ent.RoomId(rooms[0].Id),
		ent.Member{
			Id:     ent.MemberId("member-id-1"),
			Name:   ent.MemberName("member 1"),
			Status: ent.MemberStatus(ent.Invited),
		},
	)
	assert.NotEmpty(t, mId, "memberId from add member should not nil")
	assert.Empty(t, err, "add member should not return error")

	rooms, err = msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))
	assert.Empty(t, err, "should not return error when get rooms")
	member, ok := rooms[0].Members["member-id-1"]
	assert.True(t, ok, "the new member should be in room")
	assert.NotEmpty(t, member, "the new member be in the room")
	assert.Equal(t, ent.MemberId("member-id-1"), member.Id, "member-id-1", "memberId of moved member should be equal")
	assert.Equal(t, ent.MemberName("member 1"), member.Name, "member 1", "memberName of moved member should be equal")
	assert.Equal(t, ent.MemberStatus(ent.Invited), member.Status, "memberStatus of moved member should be equal")
}

func TestRemoveMember(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sessionId, _ := msr.CreateSession(context.Background(), time.Now(), time.Now().Add(time.Hour))

	msr.CreateRoom(context.Background(), sessionId, "room1")
	rooms, _ := msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))

	mId, _ := msr.AddMember(
		context.Background(),
		ent.RoomId(rooms[0].Id),
		ent.Member{
			Id:     ent.MemberId("member-id-1"),
			Name:   ent.MemberName("member 1"),
			Status: ent.MemberStatus(ent.Invited),
		},
	)

	err := msr.RemoveMember(
		context.Background(),
		ent.RoomId(rooms[0].Id),
		mId,
	)

	assert.Empty(t, err, "remove the valid member Id should not return error ")
	rooms, _ = msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))
	totalMembers := len(rooms[0].Members)
	assert.Empty(t, totalMembers, "should successfully remove member")
}

func TestMoveMember(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sessionId, _ := msr.CreateSession(context.Background(), time.Now(), time.Now().Add(time.Hour))

	msr.CreateRoom(context.Background(), sessionId, "room 1")
	msr.CreateRoom(context.Background(), sessionId, "room 2")
	rooms, _ := msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))

	mId, _ := msr.AddMember(
		context.Background(),
		ent.RoomId(rooms[0].Id),
		ent.Member{
			Id:     ent.MemberId("member-id-1"),
			Name:   ent.MemberName("member 1"),
			Status: ent.MemberStatus(ent.Invited),
		},
	)

	err := msr.MoveMember(
		context.Background(),
		mId,
		ent.RoomId(rooms[0].Id),
		ent.RoomId(rooms[1].Id),
	)

	assert.Empty(t, err, "move the valid parameter should not return error ")
	rooms, _ = msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))
	totalMembers1 := len(rooms[0].Members)
	totalMembers2 := len(rooms[1].Members)
	assert.Empty(t, totalMembers1, "should successfully move member from room 1")
	assert.Equal(t, 1, totalMembers2, "room 1 should not be empty")

	movedmember, ok := rooms[1].Members[mId]
	assert.True(t, ok, "moved member should be in room 2")
	assert.Equal(t, ent.MemberId("member-id-1"), movedmember.Id, "member-id-1", "memberId of moved member should be equal")
	assert.Equal(t, ent.MemberName("member 1"), movedmember.Name, "member 1", "memberName of moved member should be equal")
	assert.Equal(t, ent.MemberStatus(ent.Invited), movedmember.Status, "memberStatus of moved member should be equal")
}

func TestUpdateMemberStatus(t *testing.T) {
	msr := session.NewMockSessionRepo()
	sessionId, _ := msr.CreateSession(context.Background(), time.Now(), time.Now().Add(time.Hour))

	msr.CreateRoom(context.Background(), sessionId, "room 1")
	msr.CreateRoom(context.Background(), sessionId, "room 2")
	rooms, _ := msr.GetRoomsBySessionId(context.Background(), ent.SessionId(sessionId))

	mId, _ := msr.AddMember(
		context.Background(),
		ent.RoomId(rooms[0].Id),
		ent.Member{
			Id:     ent.MemberId("member-id-1"),
			Name:   ent.MemberName("member 1"),
			Status: ent.MemberStatus(ent.Invited),
		},
	)

	err := msr.UpdateMemberStatus(
		context.Background(),
		ent.RoomId(rooms[0].Id),
		mId,
		ent.MemberStatus(ent.Declined),
	)
	assert.Empty(t, err, "update member should not return error, when the params is true")
	rooms, _ = msr.GetRoomsBySessionId(context.Background(), sessionId)
	member, ok := rooms[0].Members["member-id-1"]
	assert.True(t, ok, "the member should on room 1")
	assert.Equal(t, ent.MemberStatus(ent.Declined), member.Status, "the member status should change to declined")
}
