package router

import (
	couponRoute "layout/controller/coupon"
	userRoute "layout/controller/user"

	config "layout/_config"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the HTTP routes for the API
func Router(r *gin.Engine, resource *config.Resource) *gin.Engine {
	user := r.Group("/user")
	{
		user.GET("/", userRoute.UserFunc())
	}

	coupon := r.Group("/coupon")
	{
		coupon.GET("/", couponRoute.CouponFunc(resource))
	}

	return r
}
