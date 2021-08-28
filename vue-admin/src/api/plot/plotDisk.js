import request from '@/utils/request'

const commandObj = {
  command: function (data) {
    return request({
      url: '/api/v1/plot_disk/command',
      method: 'post',
      data
    })
  },
  create: function (data) {
    return request({
      url: '/api/v1/plot_disk/create',
      method: 'post',
      data
    })
  },
  update: function (data) {
    return request({
      url: '/api/v1/plot_disk/update',
      method: 'put',
      data
    })
  },
  del: function (data) {
    return request({
      url: '/api/v1/plot_disk/delete' + data,
      method: 'delete',
      data
    })
  },
  get_list: function (params) {
    return request({
      url: '/api/v1/plot_disk/list',
      method: 'get',
      params
    })
  },
  get_list_select: function (params) {
    return request({
      url: '/api/v1/plot_disk/list_select',
      method: 'get',
      params
    })
  }
}

export default commandObj

