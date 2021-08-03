package test

import (
	"github.com/gin-gonic/gin"
	"os"
	"server/cmd/server/models"
	"server/cmd/server/redis"
	"server/cmd/server/routers"
	"testing"
)

var Router *gin.Engine

func TestMain(m *testing.M) {
	//_ = os.Setenv("TESTING", "LOCAL")
	//_ = os.Setenv("ENV_PATH", "../.env")
	Router = routers.RunGinRouter(gin.Default())
	redis.InitRedisClient()
	models.InitMongoDb()
	code := m.Run()
	os.Exit(code)
}
