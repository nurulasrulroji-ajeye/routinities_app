package web

type RoutinityUpdateReq struct {
	Id       int    `validate:"required"`
	Activity string `validate:"required,max=200,min=1" json:"activity"`
}