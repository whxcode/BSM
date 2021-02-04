import request from '@/utils/request'

const api = {
  user: '/user',
  role: '/role',
  service: '/service',
  permission: '/permission',
  permissionNoPager: '/permission/no-pager',
  orgTree: '/org/tree'
}

export default api


// 菜单列表查询
export function queryMenusList(parameter = {}) {
  return request({
    url: '/queryMenusList',
    method: 'POST',
    data: parameter
  })
}
// 删除菜单
export function menusDel(parameter) {
  return request({
    url: '/menusDel',
    method: 'POST',
    data: parameter
  })
}
// 提交@更新菜单
export function menuAdd(parameter) {
  return request({
    url: '/menuAdd',
    method: 'POST',
    data: parameter
  })
}
// 获取菜单详情
export function menuGet(parameter) {
  return request({
    url: '/menuGet',
    method: 'POST',
    data: parameter
  })
}
// 用户列表查询
export function queryUsersList(parameter = {}) {
  return request({
    url: '/queryUsersList',
    method: 'POST',
    data: parameter
  })
}
// 获取用户详情
export function getUserSingleInfo(parameter = {}) {
  return request({
    url: '/getUserSingleInfo',
    method: 'POST',
    data: parameter
  })
}
// 添加@更新用户
export function userAdd(parameter) {
  return request({
    url: '/userAdd',
    method: 'POST',
    data: parameter
  })
}
// 获取角色列表
export function queryRolesList(parameter = {}) {
  return request({
    url: '/queryRolesList',
    method: 'POST',
    data: parameter
  })
}


// 修改角色下对某一个菜单的操作权限
export  function roleModifyOperatorMenu(parameter = {}) {
  return request({
    url: '/roleModifyOperatorMenu',
    method: 'POST',
    data: parameter
  })
}

// 获取某个角色下面有那些菜单项
export function roleGetQueryMenus(parameter) {
  return request({
    url: '/roleGetQueryMenus',
    method: 'POST',
    data: parameter
  })
}

// 更新&添加角色
export function roleAdd(parameter) {
  return request({
    url: '/roleAdd',
    method: 'POST',
    data: parameter
  })
}

// 修改用户的其他信息
export function modifyBaseUser(parameter) {
  return request({
    url: '/modifyBaseUser',
    method: 'POST',
    data: parameter
  })
}

// 实时获取用户信息
export function realUser(parameter) {
  return request({
    url: '/realUser',
    method: 'POST',
    data: parameter
  })
}