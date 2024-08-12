package service

import (
	"app/routinity/model/web"
	"context"
)

type RoutinityServices interface {
	Create(ctx context.Context, req web.RoutinityCreateReq) web.RoutinityRes
	Update(ctx context.Context, req web.RoutinityUpdateReq) web.RoutinityRes
	Delete(ctx context.Context, routinityId int)
	FindById(ctx context.Context, routinityId int) web.RoutinityRes
	FindAll(ctx context.Context) []web.RoutinityRes
}