package session

import "time"

type Enum int

const (
	// Session Status
	Idle Enum = iota
	Active
	Expired

	// Member Status
	Invited
	Attending
	Declined
)

type SessionStatus Enum
type MemberStatus Enum
type MemberName string
type MemberId string
type Member struct {
	Id     MemberId
	Name   MemberName
	Status MemberStatus
}

type SessionId string
type Session struct {
	Id        SessionId
	StartTime time.Time
	EndTime   time.Time
	Status    SessionStatus
}

type RoomId string
type Room struct {
	Id           RoomId
	Name         string
	SessionId    SessionId
	ParentRoomId RoomId
	Members      map[MemberId]Member
}
