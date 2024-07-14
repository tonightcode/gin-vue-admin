package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/music"

type SongResponse struct {
	Song music.MusicSong `json:"song"`
}
