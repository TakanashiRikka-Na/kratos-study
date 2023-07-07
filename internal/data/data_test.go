package data

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestNewDB(t *testing.T) {

	_, err := gorm.Open(mysql.Open("admin:zwn123456@tcp(43.143.227.115:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestNewRedisConn(t *testing.T) {
	_, err := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}).Ping().Result()
	if err != nil {
		t.Error(err)
		return
	}
}
