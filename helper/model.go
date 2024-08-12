package helper

import (
	"app/routinity/model/domain"
	"app/routinity/model/web"
)

func ToRoutinityResponse(routinity domain.Routinity) web.RoutinityRes {
	return web.RoutinityRes{
		Id:   routinity.Id,
		Activity: routinity.Activity,
	}
}

func ToActivityResponses(routinity []domain.Routinity) []web.RoutinityRes {
	var routinityResponses []web.RoutinityRes
	for _, routinity := range routinity {
		routinityResponses = append(routinityResponses, ToRoutinityResponse(routinity))
	}
	return routinityResponses
}