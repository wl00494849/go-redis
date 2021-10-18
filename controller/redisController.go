package controller

import (
	"go-redis/cache"
	"go-redis/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Select 1 60s
var postCache1 = cache.NewPostCache(cache.NewRedisCache("172.28.0.2:6379", 1, 60))

//Select 2 360s
var postCache2 = cache.NewPostCache(cache.NewRedisCache("172.28.0.2:6379", 2, 360))

func SetRedis(ctx *gin.Context) {
	user := &model.User{}
	ctx.ShouldBindJSON(&user)
	postCache1.Set(strconv.Itoa(user.Id), user)
	ctx.JSON(200, user)
}

func GetRedis(ctx *gin.Context) {
	input := ctx.Query("id")
	data := postCache1.Get(input)

	ctx.JSON(200, data)
}

func PushRedis(ctx *gin.Context) {
	users := &[]model.User{}
	ctx.ShouldBindJSON(&users)
	for _, user := range *users {
		postCache2.Push("User", &user)
	}

	ctx.JSON(200, users)
}
func LrangeRedis(ctx *gin.Context) {
	intput := ctx.Query("key")
	data := postCache2.Lrange(intput, 0, -1)

	if data == nil {
		ctx.JSON(401, "data not found")
	}

	ctx.JSON(200, data)
}
