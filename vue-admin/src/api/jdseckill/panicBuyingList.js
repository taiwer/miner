import request from '@/utils/request'

const commandObj = {
  command: function (data) {
    return request({
      url: '/api/v1/jd/panic_buying_list/command',
      method: 'post',
      data
    })
  },
  create: function (data) {
    return request({
      url: '/api/v1/jd/panic_buying_list/create',
      method: 'post',
      data
    })
  },
  update: function (data) {
    return request({
      url: '/api/v1/jd/panic_buying_list/update',
      method: 'put',
      data
    })
  },
  del: function (data) {
    return request({
      url: '/api/v1/jd/panic_buying_list/delete' + data,
      method: 'delete',
      data
    })
  },
  get_list: function (params) {
    return request({
      url: '/api/v1/jd/panic_buying_list/list',
      method: 'get',
      params
    })
  }
}

export default commandObj

