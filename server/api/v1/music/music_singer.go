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

type SingerApi struct{}

// CreateSinger
// @Tags      Singer
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      music.MusicSinger            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *SingerApi) CreateSinger(c *gin.Context) {
	var singer music.MusicSinger
	err := c.ShouldBindJSON(&singer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(singer, utils.CreateSingerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = singerService.CreateSinger(singer)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSinger
// @Tags      Singer
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      music.MusicSinger            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /customer/customer [delete]
func (e *SingerApi) DeleteSinger(c *gin.Context) {
	var singer music.MusicSinger
	err := c.ShouldBindJSON(&singer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(singer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = singerService.DeleteSinger(singer)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSinger
// @Tags      Singer
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      music.MusicSinger            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [put]
func (e *SingerApi) UpdateSinger(c *gin.Context) {
	var singer music.MusicSinger
	err := c.ShouldBindJSON(&singer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(singer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(singer, utils.CreateSingerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = singerService.UpdateSinger(&singer)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetSinger
// @Tags      Singer
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     music.MusicSinger                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=musicRes.SingerResponse,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /customer/customer [get]
func (e *SingerApi) GetSinger(c *gin.Context) {
	var singer music.MusicSinger
	err := c.ShouldBindQuery(&singer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(singer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := singerService.GetSinger(singer.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(musicRes.SingerResponse{Singer: data}, "获取成功", c)
}

// GetSingerList
// @Tags      Singer
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *SingerApi) GetSingerList(c *gin.Context) {
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
	singerList, total, err := singerService.GetSingerInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     singerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
