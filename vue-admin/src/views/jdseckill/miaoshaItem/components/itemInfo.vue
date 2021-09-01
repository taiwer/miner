<template>
  <div>
    <el-form ref="dataForm" :model="editData" label-position="left" label-width="100px" style="margin-left:10px">
      <el-form-item label="Id" prop="title">
        <el-col :span="12">
          <el-input v-model="itemId" />
        </el-col>
        <el-col :span="4">
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="loadData">
            loadData
          </el-button>
        </el-col>
          <el-col :span="4">
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="addItemCart">
            加入购物车
          </el-button>
        </el-col>
      </el-form-item>
      <el-form-item label="Id" prop="title">
        <el-col :span="4">
          <span>{{ itemId}}</span>
        </el-col>
      </el-form-item>
      <el-form-item v-if="editData.price" label="价格打折" prop="title">
        <el-col :span="4">
          <el-input v-model="editData.price.discount" />
        </el-col>
        <el-col class="line" :span="2">价格</el-col>
        <el-col :span="4">
          <el-input v-model="editData.price.p" />
        </el-col>
        <el-col class="line" :span="2">原价</el-col>
        <el-col :span="4">
          <el-input v-model="editData.price.op" />
        </el-col>
      </el-form-item>
      <template v-if="editData.yuyueInfo" >

      <el-form-item label="预约" v-if="editData.yuyueInfo" >
        <el-col :span="4">
          <el-input v-model="editData.yuyueInfo.btnText" />
        </el-col>
        <el-col class="line" :span="2">buyTime</el-col>
        <el-col :span="4">
          <el-input v-model="editData.yuyueInfo.buyTime" />
        </el-col>
        <el-col class="line" :span="2">已预约</el-col>
        <el-col :span="4">
          <el-input v-model="editData.yuyueInfo.yuyue" />
        </el-col>
        <el-col class="line" :span="2">url</el-col>
        <el-col :span="4">
          <el-input v-model="editData.yuyueInfo.url" />
        </el-col>
      </el-form-item>
      <el-form-item label="预约时间" v-if="editData.yuyueInfo" >
        <el-col :span="4">
          <el-input v-model="editData.yuyueInfo.yuyueTime" />
        </el-col>
        <el-col class="line" :span="2">用户类型</el-col>
        <el-col :span="4">
          <el-input v-model="editData.yuyueInfo.userType" />
        </el-col>
      </el-form-item>
      </template>
      <el-form-item label="秒杀" v-if="editData.miaoshaInfo" >
        <el-col :span="4">
          <span>{{ editData.miaoshaInfo.miaosha}}</span>
        </el-col>
        <el-col class="line" :span="2">结束时间</el-col>
        <el-col :span="4">
          <span>{{ editData.miaoshaInfo.endTime}}</span>
        </el-col>
        <el-col class="line" :span="2">开始时间</el-col>
        <el-col :span="4">
          <span>{{ editData.miaoshaInfo.startTime}}</span>
        </el-col>
        <el-col class="line" :span="2">title</el-col>
        <el-col :span="4">
          <span>{{ editData.miaoshaInfo.title}}</span>
        </el-col>
        <el-col class="line" :span="2">msTrailer</el-col>
        <el-col :span="4">
          <span>{{ editData.miaoshaInfo.msTrailer}}</span>
        </el-col>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="editForm.Visible = false">
        Cancel
      </el-button>
      <el-button type="primary" @click="updateData()">
        Confirm
      </el-button>
    </div>
  </div>
</template>

<script>
import { Message } from 'element-ui'
import Command from '@/api/jdseckill/miaoshaList'

export default {
  name: 'Setting',
  props: {
    itemId: String,
    editForm: Object,
    instanceNameSource: Array,
    serverBigSource: Array,
    serverSource: Array,
    nodeSource: Array
  },
  data () {
    return {
      editData: Object
    }
  },
  computed: {
  },
  created () {
  },
  methods: {
    loadData () {
      const data = {}
      Command.send_get_item_info(data, this.itemId).then(response => {
        Message.info('success' + response.data.length)
       this.editData = response.data
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    },
    addItemCart () {
      const data = {}
      Command.send_add_item_cart(data, this.itemId).then(response => {
        Message.info('success ' + response.data)
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
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
