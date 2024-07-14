package music

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/music"
)

type SongService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSong
//@description: 创建客户
//@param: e model.Song
//@return: err error

func (exa *SongService) CreateSong(e music.MusicSong) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.Song
//@return: err error

func (exa *SongService) DeleteSong(e music.MusicSong) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSong
//@description: 更新客户
//@param: e *model.Song
//@return: err error

func (exa *SongService) UpdateSong(e *music.MusicSong) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSong
//@description: 获取客户信息
//@param: id uint
//@return: song model.Song, err error

func (exa *SongService) GetSong(id uint) (song music.MusicSong, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&song).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSongInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *SongService) GetSongInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&music.MusicSong{})
	var SongList []music.MusicSong
	err = db.Count(&total).Error
	if err != nil {
		return SongList, total, err
	} else {
		err = db.Preload("Singer").Limit(limit).Offset(offset).Find(&SongList).Error
	}
	return SongList, total, err
}
