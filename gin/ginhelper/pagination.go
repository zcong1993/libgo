package ginhelper

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/utils"
	"net/http"
	"strconv"
)

// HEADER_TOTAL_COUNT is total count in header for pagination
const HEADER_TOTAL_COUNT = "X-TOTAL-COUNT"

func parsePaginationArg(c *gin.Context, key string, defaultValue int) int {
	valueStr, ok := c.GetQuery(key)
	if !ok {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		return defaultValue
	}

	if value <= 0 {
		return defaultValue
	}

	return value
}

// OffsetLimitPaginator is struct of OffsetLimit style paginator
type OffsetLimitPaginator struct {
	DefaultNumPerPage int
}

// NewOffsetLimitPaginator init a OffsetLimitPaginator by default num per page
func NewOffsetLimitPaginator(defaultNumPerPage int) *OffsetLimitPaginator {
	return &OffsetLimitPaginator{DefaultNumPerPage: defaultNumPerPage}
}

// ParsePagination get limit and offset from request
func (op *OffsetLimitPaginator) ParsePagination(c *gin.Context) (limit, offset int) {
	return parsePaginationArg(c, "limit", op.DefaultNumPerPage), parsePaginationArg(c, "offset", 0)
}

// DefaultOffsetLimitPaginator is default OffsetLimitPaginator with 100 items per page
var DefaultOffsetLimitPaginator = NewOffsetLimitPaginator(100)

// ResponsePagination response array data with total count in header
func ResponsePagination(ctx *gin.Context, count int, data interface{}) {
	ctx.Header(HEADER_TOTAL_COUNT, utils.Num2String(count))
	ctx.JSON(http.StatusOK, data)
}

// PaginationQuery query with pagination, return total count
func PaginationQuery(db *gorm.DB, t interface{}, limit, offset int) (int, error) {
	var count int
	count = 0
	err := db.Count(&count).Error

	if err != nil {
		return count, err
	}

	err = db.Limit(limit).Offset(offset).Find(t).Error

	if err != nil {
		return count, err
	}

	return count, nil
}
