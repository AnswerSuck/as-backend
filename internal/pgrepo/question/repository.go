package question

import "github.com/ysomad/answersuck/internal/pkg/pgclient"

const (
	questionTable = "questions"
)

type Repository struct {
	*pgclient.Client
}

func NewRepository(c *pgclient.Client) *Repository {
	return &Repository{c}
}
