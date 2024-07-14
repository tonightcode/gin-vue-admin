package music

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type MusicSinger struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"comment:名称"`
	Head string `json:"url" form:"head" gorm:"comment:头像"`
}
