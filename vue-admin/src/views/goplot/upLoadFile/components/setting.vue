<template>
  <el-dialog :title="editForm.textMap[editForm.dialogStatus]" :visible.sync="editForm.Visible" style="min-width: 1000px">
    <el-form ref="dataForm" :model="editData" label-position="left" label-width="100px" style="margin-left:10px">
      <el-form-item label="Id" prop="title">
        <el-col :span="4">
          <el-sl-panel>{{ editData.id }}</el-sl-panel>
        </el-col>
      </el-form-item>
      <el-form-item label="任务名称" prop="title">
        <el-col :span="4">
          <el-input v-model="editData.task_name" />
        </el-col>
        <el-col class="line" :span="2">服务器名称</el-col>
        <el-col :span="4">
          <el-input v-model="editData.server_name" />
        </el-col>
        <el-col class="line" :span="2">开始时间</el-col>
        <el-col :span="4">
          <el-input v-model="editData.start_time" />
        </el-col>
        <el-col class="line" :span="2">结束时间</el-col>
        <el-col :span="4">
          <el-input v-model="editData.stop_time" />
        </el-col>
      </el-form-item>
      <el-form-item label="等级需求" prop="title">
        <el-col :span="4">
          <el-input v-model="editData.need_lev" />
        </el-col>
        <el-col class="line" :span="2">职业需求</el-col>
        <el-col :span="4">
          <el-input v-model="editData.need_occupation" />
        </el-col>
        <el-col class="line" :span="2">VIP等级需求</el-col>
        <el-col :span="4">
          <el-input v-model="editData.need_vip" />
        </el-col>
        <el-col class="line" :span="2">战斗力需求</el-col>
        <el-col :span="4">
          <el-input v-model="editData.need_zdl" />
        </el-col>
      </el-form-item>
      <el-form-item label="卖杂货规则">
        <el-input v-model="editData.sell_rule_text" />
      </el-form-item>
      <el-form-item label="交易规则">
        <el-input v-model="editData.trading_rule_text" />
      </el-form-item>
      <el-form-item label="捡物规则">
        <el-input v-model="editData.pick_up_rule_text" />
      </el-form-item>
      <el-form-item label="存仓规则">
        <el-input v-model="editData.ware_house_rule_text" />
      </el-form-item>
      <el-form-item label="自动打开">
        <el-input v-model="editData.use_item_rule_text" />
      </el-form-item>
      <el-form-item label="自动丢弃">
        <el-input v-model="editData.drop_rule_text" />
      </el-form-item>
      <el-form-item label="任务参数名">
        <el-input v-model="editData.param_names" />
      </el-form-item>
      <el-form-item label="任务参数值">
        <el-input v-model="editData.param_values" type="textarea" />
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="editData.descs" type="textarea" />
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
import Command from '@/api/plot/upLoad'
import { Message } from 'element-ui'

export default {
  name: 'Setting',
  props: {
    editData: Object,
    editForm: Object
  },
  data () {
    return {
      serverName: ''
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
