package player

import (
	"context"
	"fmt"
)

type (
	repository interface {
		FindByNickname(ctx context.Context, nickname string) (Player, error)
	}
)

type service struct {
	repo repository
}

func NewService(r repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) GetByNickname(ctx context.Context, nickname string) (Player, error) {
	p, err := s.repo.FindByNickname(ctx, nickname)
	if err != nil {
		return Player{}, fmt.Errorf("playerService - GetByNickname - s.repo.FindByNickname: %w", err)
	}

	return p, nil
}
