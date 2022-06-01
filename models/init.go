package models

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接
var DB = Init()

// Redis 连接
var Redis = InitRedis()
var RedisCtx = context.Background()

func Init() *gorm.DB {
	dsn := "gin_pj:iA6hrNdNyYp2tNTb@tcp(42.51.5.91:3306)/gin_pj?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
