package ginhelper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	List = iota
	Create
	Retrieve
	Update
	Destroy
)

var ReadOnly = []uint{List, Retrieve}

type IRest interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context, id string)
	Update(ctx *gin.Context, id string)
	Destroy(ctx *gin.Context, id string)
}

type Rest struct{}

func (r *Rest) List(ctx *gin.Context) {
	ctx.Status(http.StatusMethodNotAllowed)
}

func (r *Rest) Create(ctx *gin.Context) {
	ctx.Status(http.StatusMethodNotAllowed)
}

func (r *Rest) Retrieve(ctx *gin.Context, id string) {
	ctx.Status(http.StatusMethodNotAllowed)
}

func (r *Rest) Update(ctx *gin.Context, id string) {
	ctx.Status(http.StatusMethodNotAllowed)
}

func (r *Rest) Destroy(ctx *gin.Context, id string) {
	ctx.Status(http.StatusMethodNotAllowed)
}

func BindRouter(r gin.IRoutes, prefix string, rest IRest, methods ...uint) {
	withID := fmt.Sprintf("%s/:id", prefix)
	if len(methods) == 0 {
		// list
		r.GET(prefix, func(c *gin.Context) {
			rest.List(c)
		})

		// post
		r.POST(prefix, func(c *gin.Context) {
			rest.Create(c)
		})

		// get one
		r.GET(withID, func(c *gin.Context) {
			id := c.Param("id")
			rest.Retrieve(c, id)
		})

		// delete one
		r.DELETE(withID, func(c *gin.Context) {
			id := c.Param("id")
			rest.Destroy(c, id)
		})

		// update one
		r.PATCH(withID, func(c *gin.Context) {
			id := c.Param("id")
			rest.Update(c, id)
		})

		return
	}

	for _, v := range methods {
		switch v {
		case List:
			// list
			r.GET(prefix, func(c *gin.Context) {
				rest.List(c)
			})
		case Retrieve:
			// get one
			r.GET(withID, func(c *gin.Context) {
				id := c.Param("id")
				rest.Retrieve(c, id)
			})
		case Create:
			// post
			r.POST(prefix, func(c *gin.Context) {
				rest.Create(c)
			})
		case Update:
			// update one
			r.PATCH(withID, func(c *gin.Context) {
				id := c.Param("id")
				rest.Update(c, id)
			})
		case Destroy:
			// delete one
			r.DELETE(withID, func(c *gin.Context) {
				id := c.Param("id")
				rest.Destroy(c, id)
			})
		}
	}
}
