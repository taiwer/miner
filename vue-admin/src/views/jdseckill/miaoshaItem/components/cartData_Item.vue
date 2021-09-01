<template>
  <div>
    <el-row :gutter="20" style="line-height: 38px">
      <el-col :span="2">
          <el-checkbox
              v-model="item.checkedNum>0"
              @change="handleCartCheckSingle()"
          >
            {{item.checkedNum}}/{{item.totalNum}}
          </el-checkbox>
      </el-col>
      <el-col :span="3">
          <el-tag style="width: 100%">{{ itemData.Id }}</el-tag>
      </el-col>
      <el-col :span="6">
          <p class="moduletitle" style=" margin:0px">
            <el-tag>
            {{ itemData.Name }}
            </el-tag>
          </p>
      </el-col>
      <el-col :span="6">
        <el-input-number v-model="itemData.Num" @change="handleSetnum(itemData.Id, itemData.Num)" :min="1" :max="10" label="描述文字" style="height: 28px; width: 120px"></el-input-number>
        <el-tag>{{ itemData.Price }}</el-tag>
      </el-col>
      <el-col :span="6">
        <el-tag>{{ itemData.stockState }}</el-tag>
        <el-tag>{{ itemData.maxNum }}</el-tag>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Command from '@/api/jdseckill/miaoshaList'
import { Message } from 'element-ui'
import ComplexTable from '@/views/goplot/components/complex-table'
import Setting from '@/views/goplot/globalList/components/setting'
 import CartDataItem from './cartData_Items'

export default {
  name: 'CartDataItem',
  components: {
    Setting,
    CartDataItem,
    ComplexTable

  },
  props: {
    item: Object,
    itemData: Object
  },
  data () {
    return {
    }
    // Id  Name  Num  Price stockState maxNum
  },
  computed: {

  },
  watch: {
  },
  created () {
  },
  methods: {
    handleSetnum(id, num) {
      Message.info('success ' + id + " " + num)
      return
      const data = {}
      Command.send_add_item_cart(data, this.itemId).then(response => {
        Message.info('success ' + response.data)
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    },
    handleCartCheckSingle(id, num) {
      Message.info('handleCartCheckSingle ' + id + " " + num)
      if (this.item.checkedNum==0){
        const data = {}
        Command.send_cart_check_single(data, this.itemData.Id, this.item.totalNum).then(response => {
          Message.info('success ' + response.data)
        }).catch(error => {
          console.log('delete error:', error)
          Message.error({ message: error })
        })
      }else{
        const data = {}
        Command.send_cart_uncheck_single(data, this.itemData.Id, this.item.totalNum).then(response => {
          Message.info('success ' + response.data)
        }).catch(error => {
          console.log('delete error:', error)
          Message.error({ message: error })
        })
      }
      return

    }
  }
}
</script>

<style >
.el-table__expanded-cell[class*="cell"]{
  padding: 0 0px;
}
.moduletitle{
  white-space:nowrap;
  overflow:hidden;
  text-overflow:ellipsis
}
</style>
