package repository

import (
	"context"
	"fmt"

	"github.com/answersuck/vault/internal/domain"
	"github.com/answersuck/vault/pkg/postgres"
)

const (
	sessionTable = "session"
)

type sessionRepository struct {
	*postgres.Client
}

func NewSessionRepository(pg *postgres.Client) *sessionRepository {
	return &sessionRepository{pg}
}

func (r *sessionRepository) Create(ctx context.Context, s *domain.Session) (*domain.Session, error) {
	sql := fmt.Sprintf(`
		INSERT INTO %s (account_id, max_age, user_agent, ip, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, sessionTable)

	if err := r.Pool.QueryRow(ctx, sql,
		s.AccountId,
		s.MaxAge,
		s.UserAgent,
		s.IP,
		s.ExpiresAt,
		s.CreatedAt,
	).Scan(&s.Id); err != nil {
		if err = isUniqueViolation(err); err != nil {
			return nil, fmt.Errorf("r.Pool.QueryRow.Scan: %w", err)
		}

		return nil, fmt.Errorf("r.Pool.QueryRow.Scan: %w", err)
	}

	return s, nil
}

func (r *sessionRepository) FindById(ctx context.Context, sid string) (*domain.Session, error) {
	panic("implement")

	return nil, nil
}

func (r *sessionRepository) FindAll(ctx context.Context, aid string) ([]*domain.Session, error) {
	panic("implement")

	return nil, nil
}

func (r *sessionRepository) Delete(ctx context.Context, sid string) error {
	panic("implement")

	return nil
}

func (r *sessionRepository) DeleteAll(ctx context.Context, aid, sid string) error {
	panic("implement")

	return nil
}
