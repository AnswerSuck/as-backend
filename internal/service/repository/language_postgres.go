package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/answersuck/vault/internal/domain"
	"github.com/answersuck/vault/pkg/postgres"
)

const (
	languageTable = "language"
)

type languageRepository struct {
	*postgres.Client
}

func NewLanguageRepository(pg *postgres.Client) *languageRepository {
	return &languageRepository{pg}
}

func (r *languageRepository) FindAll(ctx context.Context) ([]*domain.Language, error) {
	sql := fmt.Sprintf(`SELECT id, name FROM %s`, languageTable)

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("r.Pool.Query: %w", ErrNotFound)
		}

		return nil, fmt.Errorf("r.Pool.QueryRow.Scan: %w", err)
	}

	defer rows.Close()

	var languages []*domain.Language

	for rows.Next() {
		var l domain.Language

		if err = rows.Scan(&l.Id, &l.Name); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", ErrNotFound)
		}

		languages = append(languages, &l)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return languages, nil
}
