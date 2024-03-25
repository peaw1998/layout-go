package coupon

import (
	"fmt"

	config "layout/_config"

	"github.com/gin-gonic/gin"
)

func CouponFunc(resource *config.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("CouponFunc")

		// call from coupon_service.go
		coupon := getCouponRedis(resource)
		fmt.Println("coupon", coupon)

		couponMap := groupCoupon(coupon)
		fmt.Println("couponMap", couponMap)

		// use model from coupon_model.go
		testCoupon := []couponAgg{}
		fmt.Println("testCoupon", testCoupon)

		c.JSON(200, gin.H{"message": "CouponFunc executed"})
	}
}
