package controller

type Context interface {
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	Param(string) string
}
