// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CountPessoas(ctx context.Context) (int64, error)
	CreatePessoa(ctx context.Context, arg CreatePessoaParams) (uuid.UUID, error)
	GetPessoa(ctx context.Context, id uuid.UUID) (Pessoa, error)
	GetPessoas(ctx context.Context, dollar_1 sql.NullString) ([]Pessoa, error)
}

var _ Querier = (*Queries)(nil)
