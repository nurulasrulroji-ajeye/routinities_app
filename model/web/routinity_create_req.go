package web

type RoutinityCreateReq struct {
	Activity string `validate:"required,min=1,max=100" json:"activity"`
}