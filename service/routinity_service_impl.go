package service

import (
	"app/routinity/exception"
	"app/routinity/helper"
	"app/routinity/model/domain"
	"app/routinity/model/web"
	"app/routinity/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type RoutinityServiceImpl struct {
	RoutinityRepo repository.RoutinityRepo
	DB            *sql.DB
	Validate      *validator.Validate
}

// Create / Save implements RoutinityServices.
func (service *RoutinityServiceImpl) Create(ctx context.Context, req web.RoutinityCreateReq) web.RoutinityRes {
	err := service.Validate.Struct(req)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	routinity := domain.Routinity{
		Activity: req.Activity,
	}

	routinity = service.RoutinityRepo.Save(ctx, tx, routinity)
	return helper.ToRoutinityResponse(routinity)
}

// Delete implements RoutinityServices.
func (service *RoutinityServiceImpl) Delete(ctx context.Context, routinityId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	routinity, err := service.RoutinityRepo.FindById(ctx, tx, routinityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RoutinityRepo.Delete(ctx, tx, routinity)
}

// FindAll implements RoutinityServices.
func (service *RoutinityServiceImpl) FindAll(ctx context.Context) []web.RoutinityRes {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	categories := service.RoutinityRepo.FindAll(ctx, tx)

	return helper.ToActivityResponses(categories)
}

// FindById implements RoutinityServices.
func (service *RoutinityServiceImpl) FindById(ctx context.Context, routinityId int) web.RoutinityRes {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	routinity, err := service.RoutinityRepo.FindById(ctx, tx, routinityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRoutinityResponse(routinity)
}

// Update implements RoutinityServices.
func (service *RoutinityServiceImpl) Update(ctx context.Context, req web.RoutinityUpdateReq) web.RoutinityRes {
	err := service.Validate.Struct(req)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	routinity, err := service.RoutinityRepo.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	routinity.Activity = req.Activity

	routinity = service.RoutinityRepo.Update(ctx, tx, routinity)

	return helper.ToRoutinityResponse(routinity)
}

func NewRoutinityService(routinityRepository repository.RoutinityRepo, DB *sql.DB, validate *validator.Validate) RoutinityServices {
	return &RoutinityServiceImpl{
		RoutinityRepo: routinityRepository,
		DB:            DB,
		Validate:      validate,
	}
}
