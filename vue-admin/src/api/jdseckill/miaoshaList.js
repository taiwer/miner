import request from '@/utils/request'

const commandObj = {
  command: function (data) {
    return request({
      url: '/api/v1/jd/miaosha_list/command',
      method: 'post',
      data
    })
  },
  create: function (data) {
    return request({
      url: '/api/v1/jd/miaosha_list/create',
      method: 'post',
      data
    })
  },
  update: function (data) {
    return request({
      url: '/api/v1/jd/miaosha_list/update',
      method: 'put',
      data
    })
  },
  del: function (data) {
    return request({
      url: '/api/v1/jd/miaosha_list/delete' + data,
      method: 'delete',
      data
    })
  },
  get_list: function (params) {
    return request({
      url: '/api/v1/jd/miaosha_list/list',
      method: 'get',
      params
    })
  },
  send_get_miaosha_list: function (data) {
    data.cmd = 'get_miaosha_list'
    return this.command(data)
  },
  send_get_item_info: function (data,skuId) {
    data.cmd = 'get_item_info'
    data.skuId = skuId
    return this.command(data)
  },
  send_add_item_cart: function (data,skuId) {
    data.cmd = 'add_item_cart'
    data.skuId = skuId
    return this.command(data)
  },
  send_get_cart_list: function (data) {
    data.cmd = 'get_cart_list'
    return this.command(data)
  },
  send_cart_check_single: function (data, id ,num) {
    data.cmd = 'cart_check_single'
    data.id = id
    data.num = num
    return this.command(data)
  },
  send_cart_uncheck_single: function (data, id ,num) {
    data.cmd = 'cart_uncheck_single'
    data.id = id
    data.num = num
    return this.command(data)
  },
  send_cart_uncheck_all: function (data) {
    data.cmd = 'cart_uncheck_all'
    return this.command(data)
  },
  send_get_order_info: function (data) {
    data.cmd = 'get_order_info'
    return this.command(data)
  },
  send_submit_order: function (data) {
    data.cmd = 'submit_order'
    return this.command(data)
  }
}

export default commandObj

