package music

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/music"
	musicRes "github.com/flipped-aurora/gin-vue-admin/server/model/music/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SongApi struct{}

// CreateSong
// @Tags      Song
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      music.MusicSong            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *SongApi) CreateSong(c *gin.Context) {
	var song music.MusicSong
	err := c.ShouldBindJSON(&song)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	type songV struct {
		Name string
		Type uint
	}
	song1 := songV{
		Name: song.Name,
		Type: song.Type,
	}
	err = utils.Verify(song1, utils.CreateSongVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = songService.CreateSong(song)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSong
// @Tags      Song
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      music.MusicSong            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /customer/customer [delete]
func (e *SongApi) DeleteSong(c *gin.Context) {
	var song music.MusicSong
	err := c.ShouldBindJSON(&song)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(song.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = songService.DeleteSong(song)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSong
// @Tags      Song
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      music.MusicSong            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [put]
func (e *SongApi) UpdateSong(c *gin.Context) {
	var song music.MusicSong
	err := c.ShouldBindJSON(&song)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	type songV struct {
		Name string
		Type uint
	}
	song1 := songV{
		Name: song.Name,
		Type: song.Type,
	}
	err = utils.Verify(song1, utils.CreateSongVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(song, utils.CreateSongVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = songService.UpdateSong(&song)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetSong
// @Tags      Song
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     music.MusicSong                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=musicRes.SongResponse,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /customer/customer [get]
func (e *SongApi) GetSong(c *gin.Context) {
	var song music.MusicSong
	err := c.ShouldBindQuery(&song)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(song.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := songService.GetSong(song.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(musicRes.SongResponse{Song: data}, "获取成功", c)
}

// GetSongList
// @Tags      Song
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *SongApi) GetSongList(c *gin.Context) {
	var pageInfo request.PageInfo
	fmt.Println(pageInfo)
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	songList, total, err := songService.GetSongInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     songList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
