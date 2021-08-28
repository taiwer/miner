import request from '@/utils/request'

const commandObj = {
  command: function (data) {
    return request({
      url: '/api/v1/job/command',
      method: 'post',
      data
    })
  },
  create: function (data) {
    return request({
      url: '/api/v1/job/create',
      method: 'post',
      data
    })
  },
  update: function (data) {
    return request({
      url: '/api/v1/job/update',
      method: 'put',
      data
    })
  },
  del: function (data) {
    return request({
      url: '/api/v1/job/delete' + data,
      method: 'delete',
      data
    })
  },
  get_list: function (params) {
    return request({
      url: '/api/v1/job/list',
      method: 'get',
      params
    })
  },
  get_list_select: function (params) {
    return request({
      url: '/api/v1/job/list_select',
      method: 'get',
      params
    })
  }
}

export default commandObj

