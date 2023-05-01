package member

import (
	"context"
	"fmt"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/member/mysql"
)

type Name string

type MockMemberRepository struct {
	BiggestId int64
	Members   map[Name]mysql.Member
}

func NewMockMemberRepository() *MockMemberRepository {
	return &MockMemberRepository{
		BiggestId: 0,
		Members:   map[Name]mysql.Member{},
	}
}

func (mmr *MockMemberRepository) GetMemberByName(ctx context.Context, memberName string) (member.Member, error) {
	m, ok := mmr.Members[Name(memberName)]
	if !ok {
		return member.Member{}, fmt.Errorf("member not found")
	}
	return member.Member{
		Name:   m.Name,
		Status: m.Status,
	}, nil
}

func (mmr *MockMemberRepository) CreateMember(ctx context.Context, newMember member.Member, userID int64) (member.Member, error) {
	ID := mmr.BiggestId + 1

	mmr.BiggestId = ID
	mmr.Members[Name(newMember.Name)] = mysql.Member{
		ID:     ID,
		Name:   newMember.Name,
		Status: newMember.Status,
		UserID: userID,
	}
	return newMember, nil
}
