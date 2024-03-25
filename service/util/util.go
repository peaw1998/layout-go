package util

import (
	"context"
	"time"
)

func InitContext(connectTimeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	return ctx, cancel
}
