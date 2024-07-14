package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/music"

type SingerResponse struct {
	Singer music.MusicSinger `json:"singer"`
}
