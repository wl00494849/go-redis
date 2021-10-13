package controller

import (
	"go-redis/cache"
	"go-redis/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

var postCache cache.PostCache = cache.NewRedisCache("localhost:6379", 1, 10000)

func SetRedis(ctx *gin.Context) {
	user := &model.User{}
	ctx.ShouldBindJSON(&user)
	postCache.Set(strconv.Itoa(user.Id), user)
	ctx.JSON(200, user)
}

func GetRedis(ctx *gin.Context) {

}
