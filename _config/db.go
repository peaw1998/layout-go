package config

import (
	"context"
	"layout/service/util"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 30
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

type Resource struct {
	DB      *mongo.Database
	DBLog   *mongo.Database
	MQ      *amqp.Connection
	MQCh    *amqp.Channel
	MQError chan *amqp.Error
	RDB     *redis.Client
}

// create db, mq, redis connection
func CreateResource() (*Resource, error) {
	_ = godotenv.Load()
	var err error
	var client *mongo.Client
	var dbName string
	var connectionURI string
	dbName = os.Getenv("MONGODB_DB_NAME")
	connectionURI = os.Getenv("MONGODB_ENDPOINT")

	client, err = mongo.NewClient(
		options.Client().ApplyURI(connectionURI),
		options.Client().SetMinPoolSize(3),
		options.Client().SetMaxPoolSize(10),
		options.Client().SetMaxConnIdleTime(5*time.Minute),
	)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	_ = client.Connect(ctx)
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	color.Green("Connect database successfully")
	color.Green(connectionURI)

	return &Resource{DB: client.Database(dbName)}, nil
}

func (r *Resource) Close() {
	ctx, cancel := util.InitContext(connectTimeout)
	defer cancel()

	if err := r.DB.Client().Disconnect(ctx); err != nil {
		color.Red("Close connection falure, Something wrong...")
		return
	}

	color.Cyan("Close connection successfully")
}
