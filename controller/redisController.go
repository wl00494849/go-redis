package controller

import (
	"go-redis/cache"
	"go-redis/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

var postCache *cache.PostCache

//Select 1 60s
var redis1 = cache.NewPostCache(cache.NewBasicRedisCache("172.28.0.2:6379", 1, 60))

//Select 2 360s
var redis2 = cache.NewPostCache(cache.NewPwdRedisCache("172.28.0.2:6379", 2, "7414", 360))

func TurnPostCache(pattern string) {
	switch pattern {
	case "1":
		postCache = redis1
	case "2":
		postCache = redis2
	default:
		postCache = redis1
	}
}

func SetRedis(ctx *gin.Context) {
	user := &model.User{}
	ctx.ShouldBindJSON(&user)
	postCache.Set(strconv.Itoa(user.Id), user)
	ctx.JSON(200, user)
}

func GetRedis(ctx *gin.Context) {
	input := ctx.Query("id")
	data := postCache.Get(input)

	ctx.JSON(200, data)
}

func PushRedis(ctx *gin.Context) {
	users := &[]model.User{}
	ctx.ShouldBindJSON(&users)
	for _, user := range *users {
		postCache.Push("User", &user)
	}

	ctx.JSON(200, users)
}
func LrangeRedis(ctx *gin.Context) {
	intput := ctx.Query("key")
	data := postCache.Lrange(intput, 0, -1)

	if data == nil {
		ctx.JSON(401, "data not found")
	}

	ctx.JSON(200, data)
}
