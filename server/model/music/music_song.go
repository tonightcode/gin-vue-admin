package music

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type MusicSong struct {
	global.GVA_MODEL
	Name     string      `json:"name" form:"name" gorm:"comment:名称"`
	Url      string      `json:"url" form:"url" gorm:"comment:地址"`
	Type     uint        `json:"type" form:"type" gorm:"comment:1 song 2 sound"`
	Singerid uint        `json:"singerid" form:"singerid" gorm:"comment:歌手"`
	Lyric    string      `json:"lyric" form:"lyric" gorm:"comment:歌词"`
	Singer   MusicSinger `gorm:"foreignkey:Singerid;comment:歌手"`
}
