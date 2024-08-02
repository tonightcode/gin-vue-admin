import service from '@/utils/request'
// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Song true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /song/song [post]
export const createSong = (data) => {
  return service({
    url: '/song/createSong',
    method: 'post',
    data
  })
}

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Song true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /song/song [put]
export const updateSong = (data) => {
  return service({
    url: '/song/updateSong',
    method: 'put',
    data
  })
}

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Song true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /song/song [delete]
export const deleteSong = (data) => {
  return service({
    url: '/song/song',
    method: 'delete',
    data
  })
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Song true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /song/song [get]
export const getSong = (params) => {
  return service({
    url: '/song/getSong',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /song/songList [get]
export const getSongList = (params) => {
  return service({
    url: '/song/getSongList',
    method: 'get',
    params
  })
}
