package middle

import (
	http_error "task/app/interface/http/error"
	"task/app/pkg/status"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type middleware struct {
}

func NewMiddleware() Handler {
	return &middleware{}
}

type Handler interface {
	HandleFunc(func(*Context)) gin.HandlerFunc
}

func (m *middleware) HandleFunc(
	handler func(*Context),
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerContext := &Context{
			ctx,
		}
		handler(customerContext)
	}
}

func (c *Context) Response(status status.Status, data interface{}) {
	c.JSON(status.HttpCode(), data)
}

func (c *Context) ErrorRes(err error) {
	errStatus, ok := err.(status.Status)
	if !ok {
		errStatus = status.InternalServerError
	}

	c.JSON(errStatus.HttpCode(), http_error.NewErrRes(errStatus))
}
