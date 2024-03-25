package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func (info *Resource) DelRedis(key string) {
	var ctx = context.Background()

	info.RDB.Del(ctx, key)
	return
}

func (info *Resource) GetRedis(key string, obj interface{}) error {
	var ctx = context.Background()

	byteObj, err := info.RDB.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil && strings.Contains(key, "decimal") == false {
		}
		return err
	}
	json.Unmarshal([]byte(byteObj), obj)
	return nil
}

func (info *Resource) SetRedis(key string, timeSet time.Duration, obj interface{}) error {
	var ctx = context.Background()
	byteObj, _ := json.Marshal(obj)
	err := info.RDB.SetEX(ctx, key, byteObj, timeSet).Err()
	if err != nil {
		return err
	}
	return nil
}

func (info *Resource) RedisConnect() error {
	var err error
	redisDB := 0
	if os.Getenv("REDIS_DB") != "" {
		redisDB, err = strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			fmt.Println("err_REDIS_DB : ", err)
			return err
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_URL"),
		Password:    "",      // no password set
		DB:          redisDB, // use default DB
		IdleTimeout: 1 * time.Minute,
	})
	info.RDB = rdb
	return nil
}

func (info *Resource) CloseRedis() {
	info.RDB.Close()
}

func (info *Resource) Game() *Resource {
	redisDB := 2
	rdb := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_URL"),
		Password:    "",      // no password set
		DB:          redisDB, // use default DB
		IdleTimeout: 1 * time.Minute,
	})
	return &Resource{RDB: rdb}
}

func (info *Resource) DelGame() {
	var ctx = context.Background()
	info.RDB.FlushDB(ctx)
	return
}

func (info *Resource) GetTTLRedis(key string, remainingTime *time.Duration) error {
	var ctx = context.Background()
	byteObj, err := info.RDB.TTL(ctx, key).Result()
	if err != nil {
		return err
	}
	*remainingTime = byteObj
	return nil
}
