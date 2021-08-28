import request from '@/utils/request'

const commandObj = {
  command: function (data) {
    return request({
      url: '/api/v1/up_load_file/command',
      method: 'post',
      data
    })
  },
  create: function (data) {
    return request({
      url: '/api/v1/up_load_file/create',
      method: 'post',
      data
    })
  },
  update: function (data) {
    return request({
      url: '/api/v1/up_load_file/update',
      method: 'put',
      data
    })
  },
  del: function (data) {
    return request({
      url: '/api/v1/up_load_file/delete' + data,
      method: 'delete',
      data
    })
  },
  get_list: function (params) {
    return request({
      url: '/api/v1/up_load_file/list',
      method: 'get',
      params
    })
  },
  get_list_select: function (params) {
    return request({
      url: '/api/v1/up_load_file/list_select',
      method: 'get',
      params
    })
  },
  get_game_server_version_select: function () {
    const data = {}
    data.cmd = 'get_game_server_version_select'
    return this.command(data)
  },
  up_load: function (data) {
    const req = request({
      timeout: 60000,
      url: '/api/v1/up_load_file/up_load',
      method: 'post',
      data
    })
    return req
  }
}

export default commandObj

