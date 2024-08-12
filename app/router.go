package app

import (
	"app/routinity/controller"
	"app/routinity/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(routinityController controller.RoutinityController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/routinities", routinityController.FindAll)
	router.GET("/api/routinities/:routinityId", routinityController.FindById)
	router.POST("/api/routinities", routinityController.Create)
	router.PUT("/api/routinities/:routinityId", routinityController.Update)
	router.DELETE("/api/routinities/:routinityId", routinityController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}