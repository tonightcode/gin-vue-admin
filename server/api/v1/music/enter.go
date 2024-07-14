package music

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SongApi
	SingerApi
}

var (
	songService   = service.ServiceGroupApp.MusicServiceGroup.SongService
	singerService = service.ServiceGroupApp.MusicServiceGroup.SingerService
)
