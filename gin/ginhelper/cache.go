package ginhelper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func WithCacheControl(ctx *gin.Context, d time.Duration) {
	ctx.Header("Cache-Control", fmt.Sprintf("max-age=%d", int(d.Seconds())))
}
