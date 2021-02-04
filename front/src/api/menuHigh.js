import request from '@/utils/request'
export function queryMenuHighList(parameter = {}) {
  return request({
    url: '/queryMenuHighList',
    method: 'POST',
    data: parameter
  })
}

export function menuHighAdd(parameter = {}) {
  return request({
    url: '/menuHighAdd',
    method: 'POST',
    data: parameter
  })
}


export function menuHighDel(parameter = {}) {
  return request({
    url: '/menuHighDel',
    method: 'POST',
    data: parameter
  })
}