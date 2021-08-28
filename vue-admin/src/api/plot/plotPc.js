import request from '@/utils/request'

const commandObj = {
  command: function (data) {
    return request({
      url: '/api/v1/ploter/command',
      method: 'post',
      data
    })
  },
  create: function (data) {
    return request({
      url: '/api/v1/ploter/create',
      method: 'post',
      data
    })
  },
  update: function (data) {
    return request({
      url: '/api/v1/ploter/update',
      method: 'put',
      data
    })
  },
  del: function (data) {
    return request({
      url: '/api/v1/ploter/delete' + data,
      method: 'delete',
      data
    })
  },
  get_list: function (params) {
    return request({
      url: '/api/v1/ploter/list',
      method: 'get',
      params
    })
  },
  get_list_select: function (params) {
    return request({
      url: '/api/v1/ploter/list_select',
      method: 'get',
      params
    })
  },
  get_data_list_text: function (data) {
    data.cmd = 'get_data_list_text'
    return this.command(data)
  },
  get_data_list: function (data, data_name) {
    data.cmd = 'get_data_list'
    data.data_name = data_name
    return this.command(data)
  },
  get_job_list: function (data) {
    // return this.get_data_list(data, 'getPetBagList')
    return this.get_data_list(data, 'get_job_list')
  },
  send_plot: function (data) {
    data.cmd = 'send_plot'
    return this.command(data)
  },
  send_kill: function (data, plot_id) {
    data.cmd = 'send_kill'
    data.plot_id = plot_id
    return this.command(data)
  },
  send_run_shell: function (data, shell_command) {
    data.cmd = 'send_run_shell'
    data.shell_command = shell_command
    return this.command(data)
  }
}

export default commandObj

