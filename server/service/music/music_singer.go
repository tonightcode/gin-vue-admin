package music

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/music"
)

type SingerService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSinger
//@description: 创建客户
//@param: e model.Singer
//@return: err error

func (exa *SingerService) CreateSinger(e music.MusicSinger) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.Singer
//@return: err error

func (exa *SingerService) DeleteSinger(e music.MusicSinger) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSinger
//@description: 更新客户
//@param: e *model.Singer
//@return: err error

func (exa *SingerService) UpdateSinger(e *music.MusicSinger) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSinger
//@description: 获取客户信息
//@param: id uint
//@return: singer model.Singer, err error

func (exa *SingerService) GetSinger(id uint) (singer music.MusicSinger, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&singer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSingerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *SingerService) GetSingerInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&music.MusicSinger{})
	var SingerList []music.MusicSinger
	err = db.Count(&total).Error
	if err != nil {
		return SingerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&SingerList).Error
	}
	return SingerList, total, err
}
