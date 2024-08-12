package controller

import (
	"app/routinity/helper"
	"app/routinity/model/web"
	"app/routinity/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type RoutinityControllerImpl struct {
	RoutinityService service.RoutinityServices
}

// Create implements RoutinityController.
func (controller *RoutinityControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	routinityCreateReq := web.RoutinityCreateReq{}
	helper.ReadFromRequestBody(request, &routinityCreateReq)

	categoryResponse := controller.RoutinityService.Create(request.Context(), routinityCreateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Delete implements RoutinityController.
func (controller *RoutinityControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	routinityId := params.ByName("routinityId")
	id, err := strconv.Atoi(routinityId)
	helper.PanicIfErr(err)

	controller.RoutinityService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll implements RoutinityController.
func (controller *RoutinityControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	routinityResponses := controller.RoutinityService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   routinityResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindById implements RoutinityController.
func (controller *RoutinityControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	routinityId := params.ByName("routinityId")
	id, err := strconv.Atoi(routinityId)
	helper.PanicIfErr(err)

	routinityResponse := controller.RoutinityService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   routinityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update implements RoutinityController.
func (controller *RoutinityControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	routinityUpdateReq := web.RoutinityUpdateReq{}
	helper.ReadFromRequestBody(request, &routinityUpdateReq)

	routinityIdId := params.ByName("routinityIdId")
	id, err := strconv.Atoi(routinityIdId)
	helper.PanicIfErr(err)

	routinityUpdateReq.Id = id

	routinityResponse := controller.RoutinityService.Update(request.Context(), routinityUpdateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   routinityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func NewRoutinityController(routinityService service.RoutinityServices) RoutinityController {
	return &RoutinityControllerImpl{
		RoutinityService: routinityService,
	}
}
