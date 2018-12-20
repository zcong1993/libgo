package ginhelper

import (
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

const (
	ERROR_DUPLICATE = "ERROR_DUPLICATE"
)

const (
	List = iota
	Create
	Retrieve
	Update
	Destroy
)

var ReadOnly = []uint{List, Retrieve}

type IRestView interface {
	GetQuerySet() *gorm.DB
	GetSerializers() interface{}
	GetSerializer() interface{}
	GetCreateSerializer() interface{}
	SaveData(interface{}) (interface{}, error)
	UpdateData(interface{}, string) (interface{}, error)
	GetOrderBy() string
	LookupField() string
}

type IRest interface {
	List(ctx *gin.Context, restView IRestView)
	Create(ctx *gin.Context, restView IRestView)
	Retrieve(ctx *gin.Context, restView IRestView, id string)
	Update(ctx *gin.Context, restView IRestView, id string)
	Destroy(ctx *gin.Context, restView IRestView, id string)
}

type Rest struct{}

var _ IRest = &Rest{}

func (r *Rest) List(ctx *gin.Context, restView IRestView) {
	limit, offset := DefaultOffsetLimitPaginator.ParsePagination(ctx)
	q := restView.GetQuerySet().Order(restView.GetOrderBy())

	data := restView.GetSerializers()

	count, err := PaginationQuery(q, data, limit, offset)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ResponsePagination(ctx, count, data)
}

func (r *Rest) Create(ctx *gin.Context, restView IRestView) {
	input := restView.GetCreateSerializer()
	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	res, err := restView.SaveData(input)
	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok && err.Number == mysqlerr.ER_DUP_ENTRY {
			ctx.JSON(http.StatusBadRequest, &ErrResp{Code: ERROR_DUPLICATE, Message: ERROR_DUPLICATE})
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (r *Rest) Retrieve(ctx *gin.Context, restView IRestView, id string) {
	data := restView.GetSerializer()

	err := restView.GetQuerySet().First(data, fmt.Sprintf("%s = ?", restView.LookupField()), id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (r *Rest) Update(ctx *gin.Context, restView IRestView, id string) {
	data := restView.GetSerializer()
	if restView.GetQuerySet().First(data, fmt.Sprintf("%s = ?", restView.LookupField()), id).RecordNotFound() {
		r.Create(ctx, restView)
		return
	}

	input := restView.GetCreateSerializer()
	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = restView.UpdateData(input, id)
	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok && err.Number == mysqlerr.ER_DUP_ENTRY {
			ctx.JSON(http.StatusBadRequest, &ErrResp{Code: ERROR_DUPLICATE, Message: ERROR_DUPLICATE})
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (r *Rest) Destroy(ctx *gin.Context, restView IRestView, id string) {
	err := restView.GetQuerySet().Delete(restView.GetSerializer(), fmt.Sprintf("%s = ?", restView.LookupField()), id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Status(http.StatusNoContent)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusNoContent)
}

type RestView struct{}

var _ IRestView = &RestView{}

func (r *RestView) GetQuerySet() *gorm.DB {
	panic("not implement")
}

func (r *RestView) GetSerializers() interface{} {
	panic("not implement")
}

func (r *RestView) GetSerializer() interface{} {
	panic("not implement")
}

func (r *RestView) GetCreateSerializer() interface{} {
	panic("not implement")
}

func (r *RestView) GetOrderBy() string {
	return "created_at DESC"
}

func (r *RestView) LookupField() string {
	return "id"
}

func (r *RestView) SaveData(in interface{}) (interface{}, error) {
	panic("not implement")
}

func (r *RestView) UpdateData(in interface{}, id string) (interface{}, error) {
	panic("not implement")
}

func BindRouter(r gin.IRoutes, prefix string, restView IRestView, rest IRest, methods ...uint) {
	withID := fmt.Sprintf("%s/:id", prefix)

	if len(methods) == 0 {
		// list
		r.GET(prefix, func(ctx *gin.Context) {
			rest.List(ctx, restView)
		})

		// post
		r.POST(prefix, func(ctx *gin.Context) {
			rest.Create(ctx, restView)
		})

		// get one
		r.GET(withID, func(ctx *gin.Context) {
			id := ctx.Param("id")

			rest.Retrieve(ctx, restView, id)
		})

		// delete one
		r.DELETE(withID, func(ctx *gin.Context) {
			id := ctx.Param("id")

			rest.Destroy(ctx, restView, id)
		})

		// update one
		r.PUT(withID, func(ctx *gin.Context) {
			id := ctx.Param("id")
			rest.Update(ctx, restView, id)
		})

		return
	}

	for _, v := range methods {
		switch v {
		case List:
			// list
			r.GET(prefix, func(ctx *gin.Context) {
				rest.List(ctx, restView)
			})
		case Retrieve:
			// get one
			r.GET(withID, func(ctx *gin.Context) {
				id := ctx.Param("id")

				rest.Retrieve(ctx, restView, id)
			})
		case Create:
			// post
			r.POST(prefix, func(ctx *gin.Context) {
				rest.Create(ctx, restView)
			})
		case Update:
			// update one
			r.PUT(withID, func(ctx *gin.Context) {
				id := ctx.Param("id")
				rest.Update(ctx, restView, id)
			})
		case Destroy:
			// delete one
			r.DELETE(withID, func(ctx *gin.Context) {
				id := ctx.Param("id")

				rest.Destroy(ctx, restView, id)
			})
		}
	}
}
