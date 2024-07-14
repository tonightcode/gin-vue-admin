import service from '@/utils/request'
// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Singer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /singer/singer [post]
export const createSinger = (data) => {
  return service({
    url: '/singer/createSinger',
    method: 'post',
    data
  })
}

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Singer true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /singer/singer [put]
export const updateSinger = (data) => {
  return service({
    url: '/singer/singer',
    method: 'put',
    data
  })
}

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Singer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /singer/singer [delete]
export const deleteSinger = (data) => {
  return service({
    url: '/singer/singer',
    method: 'delete',
    data
  })
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Singer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /singer/singer [get]
export const getSinger = (params) => {
  return service({
    url: '/singer/getSinger',
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
// @Router /singer/singerList [get]
export const getSingerList = (params) => {
  return service({
    url: '/singer/getSingerList',
    method: 'get',
    params
  })
}
