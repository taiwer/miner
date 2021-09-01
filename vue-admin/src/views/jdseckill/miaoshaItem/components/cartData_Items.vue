<template>
  <div>
    <div v-for="(val, key) in tableData " >
      <div v-if="val.itemType == 1">
        <CartDataItem
            :item-data = "val.item"
            :item = "val"
        />
      </div>
      <div v-else-if="tableData.itemType !== 1">   <!--分类 -->
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <el-row :gutter="20">
              <el-col :span="12">
                <span>{{val.itemType}}</span>
                <span>{{val.item.suitLabel}}</span>
                <span>{{val.item.STip}}</span>
                <span>{{val.item.PriceShow}}</span>
              </el-col>
              <el-col :span="6">
                <el-tag>{{val.skuTypeNum}}</el-tag>
                <el-tag>{{val.checkedNum}}</el-tag>
                <el-tag>{{val.totalNum}}</el-tag>
              </el-col>
              <el-col :span="6">
                <el-button style="float: right; padding: 3px 0" type="text">{{val.item.PriceShow}}</el-button>
              </el-col>
            </el-row>
          </div>
            <div v-for="(item, key) in val.item.items " >
              <CartDataItem
                  :item-data = "item.item"
                  :item = "item"
              >
              </CartDataItem>
            </div>
        </el-card>


      </div>
    </div>
  </div>
</template>

<script>
 import CartDataItem from './cartData_Item'

export default {
  name: 'CartDataItems',
  components: {
    CartDataItem

  },
  props: {
    tableData: Array
  },
  data () {
    return {
    }
  },
  computed: {
  },
  watch: {
  },
  created () {
  },
  methods: {
    }
}
</script>

<style >
.el-table__expanded-cell[class*="cell"]{
  padding: 0 0px;
}
</style>
