package ginhelper

import (
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/utils"
	"github.com/zcong1993/libgo/validator"
	"net/http"
)

const (
	// ERROR_DUPLICATE is duplicate error code
	ERROR_DUPLICATE = "ERROR_DUPLICATE"
)

const (
	// List is list method
	List = iota
	Create
	Retrieve
	Update
	Destroy
)

// ReadOnly is enum of read only methods
var ReadOnly = []uint{List, Retrieve}

// IRestView is rest view set interface
type IRestView interface {
	GetQuerySet() *gorm.DB
	GetModel(isMany bool) interface{}
	GetSerializer(isMany bool) interface{}
	GetCreateSerializer() interface{}
	SaveData(interface{}) (interface{}, error)
	UpdateData(interface{}, string) (interface{}, error)
	GetOrderBy() string
	LookupField() string
}

// IRest is rest interface
type IRest interface {
	List(ctx *gin.Context, restView IRestView)
	Create(ctx *gin.Context, restView IRestView)
	Retrieve(ctx *gin.Context, restView IRestView, id string)
	Update(ctx *gin.Context, restView IRestView, id string)
	Destroy(ctx *gin.Context, restView IRestView, id string)
}

func createInvalidErr(errors interface{}) *ErrResp {
	return &ErrResp{Code: "INVALID_PARAMS", Message: "INVALID_PARAMS", Errors: errors}
}

func mustCopy(toValue, fromValue interface{}) {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		panic("copy error")
	}
}

// Rest is struct impl IRest interface
type Rest struct{}

var _ IRest = &Rest{}

func getCreateSerializer(restView IRestView) interface{} {
	cs := restView.GetCreateSerializer()
	if cs == nil {
		return restView.GetModel(false)
	}
	return cs
}

// List impl IRest's List
func (r *Rest) List(ctx *gin.Context, restView IRestView) {
	limit, offset := DefaultOffsetLimitPaginator.ParsePagination(ctx)
	q := restView.GetQuerySet().Order(restView.GetOrderBy())

	data := restView.GetModel(true)

	count, err := PaginationQuery(q, data, limit, offset)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	out := data
	se := restView.GetSerializer(true)

	if se != nil {
		out = se
		mustCopy(out, data)
	}

	ResponsePagination(ctx, count, out)
}

// Create impl IRest's Create
func (r *Rest) Create(ctx *gin.Context, restView IRestView) {
	input := getCreateSerializer(restView)
	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, createInvalidErr(validator.NormalizeErr(err)))
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

	out := res

	se := restView.GetSerializer(false)

	if se != nil {
		out = se
		mustCopy(out, res)
	}

	ctx.JSON(http.StatusCreated, out)
}

// Retrieve impl IRest's Retrieve
func (r *Rest) Retrieve(ctx *gin.Context, restView IRestView, id string) {
	data := restView.GetModel(false)

	err := restView.GetQuerySet().First(data, fmt.Sprintf("%s = ?", restView.LookupField()), id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	out := data
	se := restView.GetSerializer(false)

	if se != nil {
		out = se
		mustCopy(out, data)
	}

	ctx.JSON(http.StatusOK, out)
}

// Update impl IRest's Update
func (r *Rest) Update(ctx *gin.Context, restView IRestView, id string) {
	data := restView.GetModel(false)
	if restView.GetQuerySet().First(data, fmt.Sprintf("%s = ?", restView.LookupField()), id).RecordNotFound() {
		r.Create(ctx, restView)
		return
	}

	input := getCreateSerializer(restView)
	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, createInvalidErr(validator.NormalizeErr(err)))
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

// Destroy impl IRest's Destroy
func (r *Rest) Destroy(ctx *gin.Context, restView IRestView, id string) {
	err := restView.GetQuerySet().Delete(restView.GetModel(false), fmt.Sprintf("%s = ?", restView.LookupField()), id).Error
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

// RestView is struct impl IRestView interface
type RestView struct{}

var _ IRestView = &RestView{}

// GetQuerySet impl IRestView's GetQuerySet
func (r *RestView) GetQuerySet() *gorm.DB {
	panic("not implement")
}

// GetModel impl IRestView's GetModel
func (r *RestView) GetModel(isMany bool) interface{} {
	panic("not implement")
}

func (r *RestView) GetSerializer(isMany bool) interface{} {
	return nil
}

// GetCreateSerializer impl IRestView's GetCreateSerializer
func (r *RestView) GetCreateSerializer() interface{} {
	panic("not implement")
}

// GetOrderBy impl IRestView's GetOrderBy
func (r *RestView) GetOrderBy() string {
	return "created_at DESC"
}

// LookupField impl IRestView's LookupField
func (r *RestView) LookupField() string {
	return "id"
}

// SaveData impl IRestView's SaveData
func (r *RestView) SaveData(in interface{}) (interface{}, error) {
	panic("not implement")
}

// UpdateData impl IRestView's UpdateData
func (r *RestView) UpdateData(in interface{}, id string) (interface{}, error) {
	panic("not implement")
}

// BindRouter bind rest group routers
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

	mds := utils.NewUintSet(methods...)
	methods = mds.ToSlice()

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
