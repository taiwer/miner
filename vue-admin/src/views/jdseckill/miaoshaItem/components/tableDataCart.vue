<template>
  <div>
    <div>
      <el-row :gutter="10" style="margin-top: 20px">
        <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleUnCheckAll">
            取消全部
          </el-button>
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleGetCartList">
            获取购物车列表
          </el-button>
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleGetOrderInfo">
            获取订单信息
          </el-button>
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleSubmitOrder">
            提交订单
          </el-button>
        </el-col>
        <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
          <div />
        </el-col>
      </el-row>
      <cart-data-list
        ref="cart_data_list"
        :tableData = "tableData"
      >
      </cart-data-list>
    </div>
  </div>
</template>

<script>
import Command from '@/api/jdseckill/miaoshaList'
import { Message } from 'element-ui'
import CartDataList from "./cartDataList";

export default {
  name: 'TableDataCart',
  components: {
    CartDataList,
  },
  props: {
    account: String
  },
  data () {
    return {
      tableData: []
    }
  },
  computed: {
  },
  watch: {
  },
  created () {
  },
  methods: {
    handleUnCheckAll() {
      const data = {}
      Command.send_cart_uncheck_all(data).then(response => {
        Message.info('success ' + response.data.success)
        const data = response.data;
        const resultData = data.resultData
        const cartInfo = resultData.cartInfo
        if (data.success){
          this.$refs.cart_data_list.tableData = cartInfo.vendors
          this.$refs.cart_data_list.cartInfo = cartInfo
        }
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    },
    handleGetCartList () {
      const data = {}
      Command.send_get_cart_list(data).then(response => {
        Message.info('success ' + response.data.success)
        const data = response.data;
        const resultData = data.resultData
        const cartInfo = resultData.cartInfo
        if (data.success){
          this.$refs.cart_data_list.tableData = cartInfo.vendors
          this.$refs.cart_data_list.cartInfo = cartInfo
        }
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    },
    handleGetOrderInfo () {
      const data = {}
      Command.send_get_order_info(data).then(response => {
        Message.info('success ' + response.data)
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    },
    handleSubmitOrder () {
      const data = {}
      Command.send_submit_order(data).then(response => {
        Message.info('success ' + response.data.success)
        const data = response.data;
        const resultData = data.resultData
        const cartInfo = resultData.cartInfo
        if (data.success){
          this.$refs.cart_data_list.cartInfo = cartInfo
          this.$refs.cart_data_list.tableData = cartInfo.vendors
        }
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    }
  }
}
</script>

<style >
.el-table__expanded-cell[class*="cell"]{
  padding: 0 0px;
}
</style>
