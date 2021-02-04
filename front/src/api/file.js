import request from '@/utils/request'
export function uploadSingleFile(formData) {
  return request({
    url: '/uploadFile',
    method: 'POST',
    data: formData
  })
}