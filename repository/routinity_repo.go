package repository

import (
	"app/routinity/model/domain"
	"context"
	"database/sql"
)

type RoutinityRepo interface {
	Save(ctx context.Context, tx *sql.Tx, routinity domain.Routinity) domain.Routinity
	Update(ctx context.Context, tx *sql.Tx, routinity domain.Routinity) domain.Routinity
	Delete(ctx context.Context, tx *sql.Tx, routinity domain.Routinity)
	FindById(ctx context.Context, tx *sql.Tx, routinityId int) (domain.Routinity, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Routinity
}