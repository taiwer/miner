<template>
  <el-dialog :title="editForm.textMap[editForm.dialogStatus]" :visible.sync="editForm.Visible" style="min-width: 1000px">
    <el-form ref="dataForm" :model="editData" label-position="left" label-width="100px" style="margin-left:10px">
      <el-form-item label="Id" prop="title">
        <el-col :span="4">
          <span>{{ editData.id }}</span>
        </el-col>
      </el-form-item>
      <el-form-item label="名称" prop="title">
        <el-col :span="4">
          <el-input v-model="editData.name" />
        </el-col>
        <el-col class="line" :span="2"> 值</el-col>
        <el-col :span="4">
          <el-input v-model="editData.val" />
        </el-col>
        <el-col class="line" :span="2">文本值</el-col>
        <el-col :span="4">
          <el-input v-model="editData.text" />
        </el-col>
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="editData.descs" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="editForm.Visible = false">
        Cancel
      </el-button>
      <el-button type="primary" @click="editForm.dialogStatus==='create'?createData():updateData()">
        Confirm
      </el-button>
    </div>
  </el-dialog>
</template>

<script>
import Command from '@/api/plot/globalSetting'
import { Message } from 'element-ui'

export default {
  name: 'Setting',
  props: {
    editData: Object,
    editForm: Object,
    instanceNameSource: Array,
    serverBigSource: Array,
    serverSource: Array,
    nodeSource: Array
  },
  data () {
    return {
      serverName: ''
    }
  },
  computed: {
    serverGroupSourceData: {
      get () {
        const source = []
        this.serverBigSource.forEach(element => {
          // source.push({ value: element.name, label: element.name })
          source.push(element)
        })
        return source
      }
    },
    serverNameSourceData: {
      get () {
        const source = []
        this.serverSource.forEach(element => {
          if (this.editData.server_name_big === element.group_name) {
            source.push(element)
          }
        })
        return source
      }
    }
  },
  created () {
  },
  methods: {
    createData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          Command.create(this.editData).then(() => {
            Message.info('add success')
            this.$emit('reLoadList')
            // this.editForm.Visible = false
          }).catch(error => {
            console.log('add error:', error)
            Message.error({ message: error })
            this.editForm.Visible = false
          })
        }
      })
    },
    updateData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          Command.update(this.editData).then(() => {
            Message.info('add success')
            this.$emit('reLoadList')
            this.editForm.Visible = false
          }).catch(error => {
            Message.error({ message: error })
            console.log('add error:', error)
            Message.error({ message: error })
            this.editForm.Visible = false
          })
        }
      })
    }
  }
}
</script>

<style>
div .el-dialog {
  width: 70%;
}
.el-form-item {
  margin-bottom: 5px;
}
</style>
