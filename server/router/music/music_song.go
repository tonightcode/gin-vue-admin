package music

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SongRouter struct{}

func (s *SongRouter) InitSongRouter(Router *gin.RouterGroup) {
	songRouter := Router.Group("song").Use(middleware.OperationRecord())
	songApi := v1.ApiGroupApp.MusicApiGroup.SongApi
	{
		songRouter.GET("getSong", songApi.GetSong)         // 获取
		songRouter.GET("getSongList", songApi.GetSongList) // 获取列表
		songRouter.POST("createSong", songApi.CreateSong)  // 创建
		songRouter.POST("deleteSong", songApi.DeleteSong)  // 删除
		songRouter.PUT("updateSong", songApi.UpdateSong)   // 更新
	}
}
