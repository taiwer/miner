<template>
  <el-upload
    ref="newupload"
    class="upload-file"
    drag
    :action="doUpload"
    :before-upload="beforeUpload"
    multiple
    :auto-upload="true"
  >
    <i class="el-icon-upload" />
    <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
  </el-upload>
</template>

<script>
import CommandUpLoad from '@/api/plot/upLoad'
import { Message } from 'element-ui'

export default {
  name: 'UpLoad',
  data () {
    return {
      doUpload: '/api/up/file'
    }
  },
  methods: {
    beforeUpload (file) {
      const fd = new FormData()
      fd.append('file', file)// 传文件
      fd.append('srid', 'filename')// 传其他参数
      CommandUpLoad.up_load(fd).then(response => {
        Message.info('success')
        console.log('created response:', response)
        this.$emit('reLoadList')
      }).catch(error => {
        Message.error({ message: error })
        console.log('upLoadFile err:', error)
      })
    },
    newSubmitForm () { // 确定上传
      this.$refs.newupload.submit()
    }
  }
}
</script>

<style scoped>

</style>
