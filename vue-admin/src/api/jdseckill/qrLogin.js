import request from '@/utils/request'

const commandObj = {
    command: function (data) {
        return request({
            url: '/user/qr_login/command',
            method: 'post',
            data
        })
    },
    create: function (data) {
        return request({
            url: '/user/qr_login/create',
            method: 'post',
            data
        })
    },
    update: function (data) {
        return request({
            url: '/user/qr_login/update',
            method: 'put',
            data
        })
    },
    del: function (data) {
        return request({
            url: '/user/qr_login/delete' + data,
            method: 'delete',
            data
        })
    },
    get_list: function (params) {
        return request({
            url: '/user/qr_login/list',
            method: 'get',
            params
        })
    },
    send_show: function (data) {
        data.cmd = 'show'
        return this.command(data)
    },
    send_get_tick: function (data,wlfstkSmdl) {
        data.cmd = 'get_tick'
        data.wlfstkSmdl = wlfstkSmdl
        return this.command(data)
    },
    send_get_token: function (data,tick) {
        data.cmd = 'get_token'
        data.tick = tick
        return this.command(data)
    }
}

export default commandObj

