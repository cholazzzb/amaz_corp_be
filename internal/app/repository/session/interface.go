package session

import (
	"context"
	"time"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/session"
)

type SessionRepo interface {
	SessionRepository
	RoomRepository
}

type SessionRepository interface {
	GetSessionById(
		ctx context.Context,
		sessionId session.SessionId,
		roomName string,
	) (session.Session, error)
	CreateSession(
		ctx context.Context,
		startTime time.Time,
		endTime time.Time,
	) (session.SessionId, error)
	UpdateSessionStatus(
		ctx context.Context,
		sessionId session.SessionId,
		newStatus session.SessionStatus,
	) error
}

type RoomRepository interface {
	GetRoomById(
		ctx context.Context,
		roomId session.RoomId,
	) (session.Room, error)
	GetRoomsBySessionId(
		ctx context.Context,
		sessionId session.SessionId,
	) ([]session.Room, error)
	CreateRoom(
		ctx context.Context,
		sessionId session.SessionId,
	) (session.RoomId, error)
	AddMember(
		ctx context.Context,
		roomId session.RoomId,
		member session.Member,
	) (session.MemberId, error)
	RemoveMember(
		ctx context.Context,
		roomId session.RoomId,
		memberId session.MemberId,
	) error
	MoveMember(
		ctx context.Context,
		memberId session.MemberId,
		sourceId session.RoomId,
		targetId session.RoomId,
	) error
	UpdateMemberStatus(
		ctx context.Context,
		roomId session.RoomId,
		memberId session.MemberId,
		memberStatus session.MemberStatus,
	) error
}
