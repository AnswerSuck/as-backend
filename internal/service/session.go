package service

import (
	"context"
	"fmt"
	"github.com/answersuck/vault/internal/dto"

	"github.com/answersuck/vault/internal/config"
	"github.com/answersuck/vault/internal/domain"
)

type sessionService struct {
	cfg  *config.Session
	repo SessionRepo
}

func NewSessionService(cfg *config.Session, s SessionRepo) *sessionService {
	return &sessionService{
		cfg:  cfg,
		repo: s,
	}
}

func (s *sessionService) Create(ctx context.Context, aid string, d dto.Device) (*domain.Session, error) {
	sess, err := domain.NewSession(aid, d.UserAgent, d.IP, s.cfg.Expiration)
	if err != nil {
		return nil, fmt.Errorf("sessionService - Create - domain.NewSession: %w", err)
	}

	sess, err = s.repo.Create(ctx, sess)
	if err != nil {
		return nil, fmt.Errorf("sessionService - Create - s.repo.Create: %w", err)
	}

	return sess, nil
}

func (s *sessionService) GetById(ctx context.Context, sid string) (*domain.Session, error) {
	sess, err := s.repo.FindById(ctx, sid)
	if err != nil {
		return nil, fmt.Errorf("sessionService - Get - s.repo.FindByID: %w", err)
	}

	return sess, nil
}

func (s *sessionService) GetAll(ctx context.Context, aid string) ([]*domain.Session, error) {
	panic("implement")

	return nil, nil
}

func (s *sessionService) Terminate(ctx context.Context, sid string) error {
	if err := s.repo.Delete(ctx, sid); err != nil {
		return fmt.Errorf("sessionService - Terminate - s.repo.Delete: %w", err)
	}

	return nil
}

func (s *sessionService) TerminateAll(ctx context.Context, aid, sid string) error {
	panic("implement")

	return nil
}
