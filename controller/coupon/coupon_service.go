package coupon

import (
	"fmt"
	"layout/internal/model"
	"time"

	config "layout/_config"

	"github.com/go-redis/redis/v8"
)

// get coupon from redis
func getCouponRedis(resource *config.Resource) (coupon []model.Coupon) {
	key := "coupon"
	err := resource.GetRedis(key, &coupon)
	if err != nil && err != redis.Nil {

		return
	}

	if err == redis.Nil {
		timeSet := time.Duration(24 * int64(time.Hour))
		err = resource.SetRedis(key, timeSet, coupon)
		if err != nil {
			fmt.Println("cannot_set_redis_article:", err)
		}
	}

	return coupon
}

// group coupon by code
func groupCoupon(coupon []model.Coupon) map[string]model.Coupon {
	couponMap := make(map[string]model.Coupon)

	for _, v := range coupon {
		couponMap[v.Code] = v
	}

	return couponMap
}
