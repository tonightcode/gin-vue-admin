package music

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SingerRouter struct{}

func (s *SingerRouter) InitSingerRouter(Router *gin.RouterGroup) {
	singerRouter := Router.Group("singer").Use(middleware.OperationRecord())
	singerApi := v1.ApiGroupApp.MusicApiGroup.SingerApi
	{
		singerRouter.GET("getSinger", singerApi.GetSinger)         // 获取
		singerRouter.GET("getSingerList", singerApi.GetSingerList) // 获取列表
		singerRouter.POST("createSinger", singerApi.CreateSinger)  // 创建
		singerRouter.POST("deleteSinger", singerApi.DeleteSinger)  // 删除
		singerRouter.PUT("updateSinger", singerApi.UpdateSinger)   // 更新
	}
}
