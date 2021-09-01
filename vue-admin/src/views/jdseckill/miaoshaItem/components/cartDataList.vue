<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">

      <el-row :gutter="20">
        <el-col :span="6">
          <span>购物车</span>
          <el-button style="float: right; padding: 3px 0" type="text" @click="handleUnCheckAll">取消全部</el-button>
        </el-col>
        <el-col :span="12" v-if="cartInfo !== undefined">
          <el-tag>{{cartInfo.checkedWareNum}}</el-tag>
          <el-tag>{{cartInfo.RePrice}}</el-tag>
          <el-tag>{{cartInfo.Price}}</el-tag>
          <el-tag>{{cartInfo.PriceShow}}</el-tag>
        </el-col>
        <el-col :span="6">
          <el-button style="float: right; padding: 3px 0" type="text"></el-button>
        </el-col>
      </el-row>
    </div>
    <div v-for="(val, key) in tableData" class="text item">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>{{val.shopName}}</span>
          <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
        </div>
        <div>
          <CartDataItems
              :tableData="val.sorted"
          >
          </CartDataItems>
        </div>
      </el-card>
    </div>
  </el-card>
</template>

<script>
import Command from '@/api/jdseckill/miaoshaList'
import CartDataItems from './cartData_Items'
import {Message} from "element-ui";

export default {
  name: "cartDataList",
  components: {
    CartDataItems
  },
  props: {
    tableData: Array,
    cartInfo: Object
  },
  methods:{
    handleUnCheckAll() {
      const data = {}
      Command.send_cart_uncheck_all(data).then(response => {
        Message.info('success ' + response.data)
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    }
  }
}
</script>

<style>
.el-card > .el-card__body{
  padding: 1px;
}
</style>
