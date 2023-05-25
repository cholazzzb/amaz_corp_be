package session

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/session"
)

type MockSessionRepo struct {
	msr *MockSessionRepository
	mrr *MockRoomRepository
}

type SessionRoom struct {
	ses   session.Session
	rooms []session.RoomId
}

type MockSessionRepository struct {
	Sessions map[session.SessionId]*SessionRoom
}

func NewMockSessionRepository() *MockSessionRepository {
	return &MockSessionRepository{
		Sessions: map[session.SessionId]*SessionRoom{},
	}
}

type MockRoomRepository struct {
	Rooms map[session.RoomId]*session.Room
}

func NewMockRoomRepository() *MockRoomRepository {
	return &MockRoomRepository{
		Rooms: map[session.RoomId]*session.Room{},
	}
}

func NewMockSessionRepo() *MockSessionRepo {
	msr := NewMockSessionRepository()
	mrr := NewMockRoomRepository()

	return &MockSessionRepo{
		msr,
		mrr,
	}
}

func (repo *MockSessionRepo) GetSessionById(
	ctx context.Context,
	sessionId session.SessionId,
) (session.Session, error) {
	ses, ok := repo.msr.Sessions[sessionId]
	if !ok {
		return session.Session{}, fmt.Errorf("session not found")
	}
	return ses.ses, nil
}

func (repo *MockSessionRepo) CreateSession(
	ctx context.Context,
	startTime time.Time,
	endTime time.Time,
) (session.SessionId, error) {
	uuidByte, err := exec.Command("uuidgen").Output()
	if err != nil {
		return session.SessionId(""), err
	}
	uuid := string(uuidByte[:])
	s := session.Session{
		Id:        session.SessionId(uuid),
		StartTime: startTime,
		EndTime:   endTime,
		Status:    session.SessionStatus(session.Idle),
	}
	repo.msr.Sessions[s.Id] = &SessionRoom{
		ses:   s,
		rooms: []session.RoomId{},
	}
	return session.SessionId(uuid), nil
}

func (repo *MockSessionRepo) UpdateSessionStatus(
	ctx context.Context,
	sessionId session.SessionId,
	newStatus session.SessionStatus,
) error {
	_, err := repo.GetSessionById(ctx, sessionId)
	if err != nil {
		return err
	}
	updated := repo.msr.Sessions[sessionId]
	updated.ses.Status = newStatus
	repo.msr.Sessions[sessionId] = updated
	return nil
}

// Rooms
func (repo *MockSessionRepo) GetRoomById(
	ctx context.Context,
	roomId session.RoomId,
) (session.Room, error) {
	room, ok := repo.mrr.Rooms[roomId]
	if !ok {
		return session.Room{}, fmt.Errorf("room not found")
	}
	return *room, nil
}

func (repo *MockSessionRepo) GetRoomsBySessionId(
	ctx context.Context,
	sessionId session.SessionId,
) ([]session.Room, error) {
	ses, ok := repo.msr.Sessions[sessionId]
	if !ok {
		return nil, nil
	}
	rooms := []session.Room{}
	for _, roomId := range ses.rooms {
		room, err := repo.GetRoomById(ctx, roomId)
		if err != nil {
			panic("the roomId and room map are not sync")
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (repo *MockSessionRepo) CreateRoom(
	ctx context.Context,
	sessionId session.SessionId,
	roomName string,
) (session.RoomId, error) {
	_, err := repo.GetSessionById(ctx, sessionId)
	if err != nil {
		return session.RoomId(""), err
	}

	uuidByte, err := exec.Command("uuidgen").Output()
	if err != nil {
		return session.RoomId(""), err
	}
	uuid := string(uuidByte[:])
	roomId := session.RoomId(uuid)
	newRoom := session.Room{
		Id:           roomId,
		Name:         roomName,
		SessionId:    sessionId,
		ParentRoomId: "",
		Members:      map[session.MemberId]session.Member{},
	}

	repo.msr.Sessions[sessionId].rooms = append(
		repo.msr.Sessions[sessionId].rooms,
		roomId,
	)

	repo.mrr.Rooms[roomId] = &newRoom

	return session.RoomId(uuid), nil
}

func (repo *MockSessionRepo) AddMember(
	ctx context.Context,
	roomId session.RoomId,
	member session.Member,
) (session.MemberId, error) {
	room, err := repo.GetRoomById(ctx, roomId)
	if err != nil {
		return session.MemberId(""), fmt.Errorf("room not found")
	}
	room.Members[member.Id] = member
	repo.mrr.Rooms[roomId] = &room
	return member.Id, nil
}

func (repo *MockSessionRepo) RemoveMember(
	ctx context.Context,
	roomId session.RoomId,
	memberId session.MemberId,
) error {
	_, err := repo.GetRoomById(ctx, roomId)
	if err != nil {
		return fmt.Errorf("room not found")
	}
	delete(repo.mrr.Rooms[roomId].Members, memberId)
	return nil
}

func (repo *MockSessionRepo) MoveMember(
	ctx context.Context,
	memberId session.MemberId,
	sourceId session.RoomId,
	targetId session.RoomId,
) error {
	member, ok := repo.mrr.Rooms[sourceId].Members[memberId]
	if !ok {
		return fmt.Errorf("the member is not found in source Id room")
	}
	_, ok = repo.mrr.Rooms[targetId].Members[memberId]
	if ok {
		return fmt.Errorf("the member already in the targetId")
	}

	repo.AddMember(ctx, targetId, member)
	repo.RemoveMember(ctx, sourceId, memberId)

	return nil
}

func (repo *MockSessionRepo) UpdateMemberStatus(
	ctx context.Context,
	roomId session.RoomId,
	memberId session.MemberId,
	memberStatus session.MemberStatus,
) error {
	member, ok := repo.mrr.Rooms[roomId].Members[memberId]
	if !ok {
		return fmt.Errorf("memberId or roomId not found")
	}
	member.Status = memberStatus
	repo.mrr.Rooms[roomId].Members[memberId] = member
	return nil
}
