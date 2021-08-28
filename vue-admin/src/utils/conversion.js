// eslint-disable-next-line no-unused-vars

const commandObj = {
  conver: function (limit) {
    var size = ''
    if (limit < 0.1 * 1024) { // 如果小于0.1KB转化成B
      size = limit.toFixed(1) + 'B'
    } else if (limit < 0.1 * 1024 * 1024) { // 如果小于0.1MB转化成KB
      size = (limit / 1024).toFixed(1) + 'K'
    } else if (limit < 0.1 * 1024 * 1024 * 1024) { // 如果小于0.1GB转化成MB
      size = (limit / (1024 * 1024)).toFixed(1) + 'M'
    } else if (limit < 0.1 * 1024 * 1024 * 1024 * 1024) { // 其他转化成GB
      size = (limit / (1024 * 1024 * 1024)).toFixed(1) + 'G'
    } else {
      size = (limit / (1024 * 1024 * 1024 * 1024)).toFixed(1) + 'T'
    }

    var sizestr = size + ''
    var len = sizestr.indexOf('\.')
    var dec = sizestr.substr(len + 1, 2)
    // eslint-disable-next-line eqeqeq
    if (dec == '00') { // 当小数点后为00时 去掉小数部分
      return sizestr.substring(0, len) + sizestr.substr(len + 3, 2)
    }
    return sizestr
  }
}

export default commandObj
