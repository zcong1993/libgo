package ginhelper

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

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

type OffsetLimitPaginator struct {
	DefaultNumPerPage int
}

func NewOffsetLimitPaginator(defaultNumPerPage int) *OffsetLimitPaginator {
	return &OffsetLimitPaginator{DefaultNumPerPage: defaultNumPerPage}
}

func (op *OffsetLimitPaginator) ParsePagination(c *gin.Context) (limit, offset int) {
	return parsePaginationArg(c, "limit", op.DefaultNumPerPage), parsePaginationArg(c, "offset", 0)
}

var DefaultOffsetLimitPaginator = NewOffsetLimitPaginator(100)
