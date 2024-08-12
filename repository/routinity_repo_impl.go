package repository

import (
	"app/routinity/helper"
	"app/routinity/model/domain"
	"context"
	"database/sql"
	"errors"
)

type RoutinityRepoImpl struct {
}

func NewRoutintityRepo() RoutinityRepo {
	return &RoutinityRepoImpl{}
}

// Save implements RoutinityRepo.
func (r *RoutinityRepoImpl) Save(ctx context.Context, tx *sql.Tx, routinity domain.Routinity) domain.Routinity {
	SQL := "INSERT INTO routinity(activity) values(?)"

	res, err := tx.ExecContext(ctx, SQL, routinity.Activity)
	helper.PanicIfErr(err)

	id, err := res.LastInsertId()
	helper.PanicIfErr(err)

	routinity.Id = int(id)

	return routinity
}

// FindAll implements RoutinityRepo.
func (r *RoutinityRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Routinity {
	SQL := "SELECT id, activity FROM routinity"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	
	var routinities []domain.Routinity
	
	for rows.Next() {
		routinity := domain.Routinity{}
		err := rows.Scan(&routinity.Id, &routinity.Activity)
		helper.PanicIfErr(err)
		routinities = append(routinities, routinity)
	}

	return routinities

}

// FindById implements RoutinityRepo.
func (r *RoutinityRepoImpl) FindById(ctx context.Context, tx *sql.Tx, routinityId int) (domain.Routinity, error) {
	SQL := "SELECT id, activity FROM routinity where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, routinityId)
	helper.PanicIfErr(err)

	routinity := domain.Routinity{}
	if rows.Next() {
		err := rows.Scan(&routinity.Id, routinity.Activity)
		helper.PanicIfErr(err)
		return routinity, nil
	} else {
		return routinity, errors.New("routinity not found")
	}
}

// Update implements RoutinityRepo.
func (r *RoutinityRepoImpl) Update(ctx context.Context, tx *sql.Tx, routinity domain.Routinity) domain.Routinity {
	SQL := "UPDATE routinity set activity = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, routinity.Activity, routinity.Id)

	helper.PanicIfErr(err)

	return routinity
}

// Delete implements RoutinityRepo.
func (r *RoutinityRepoImpl) Delete(ctx context.Context, tx *sql.Tx, routinity domain.Routinity) {
	SQL := "DELETE from routinity where id = ?"
	_, err := tx.ExecContext(ctx, SQL, routinity.Id)
	helper.PanicIfErr(err)
}