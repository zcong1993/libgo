package ginhelper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func BindRouter(r gin.IRoutes, prefix string, rest IRest) {
	withID := fmt.Sprintf("%s/:id", prefix)
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
}
