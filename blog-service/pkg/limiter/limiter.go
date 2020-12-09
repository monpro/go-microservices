package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type ILimiter interface {
	Key(context *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...BucketRule) ILimiter
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type BucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
