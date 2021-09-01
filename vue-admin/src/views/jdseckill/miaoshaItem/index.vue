<template xmlns="http://www.w3.org/1999/html">
  <div class="dashboard-editor-container">
    <div class="chart-wrapper">
      <el-row :gutter="10" style="margin-top: 20px">
        <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
          <el-input v-model="itemId" />
        </el-col>
        <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleGetItemInfo">
            getItemInfo
          </el-button>
        </el-col>
      </el-row>
      <div style="padding: 15px">
        <ItemInfo key="data-info"
          :itemId="itemId"
          :editData="editData"
        />
      </div>
      <div style="padding: 15px">
        <TableDataCart key="data-info"

        />
      </div>
    </div>
  </div>
</template>

<script>
import Command from '@/api/plot/globalSetting'
import { Message } from 'element-ui'
import ItemInfo from './components/itemInfo'
import TableDataCart from './components/tableDataCart'

export default {
  name: 'MiaoshaItem',
  components: {
    ItemInfo,
    TableDataCart
  },
  data () {
    return {
      itemId: '10035163404173',
      editData: Object,
      dialogFormVisible: false,
      dialogStatus: '',
      statusOptions: ['published', 'draft', 'deleted'],
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      temp: {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        type: '',
        status: 'published'
      }
    }
  },
  created () {
  },
  methods: {
    resetTemp () {
      this.temp = {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        status: 'published',
        type: ''
      }
    },
    handleCreate () {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleGetItemInfo () {
      this.$refs.edit_table.loadData() /* 重新刷新 */
    },
    createData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.temp.id = parseInt(Math.random() * 100) + 1024 // mock a id
          this.temp.author = 'vue-element-admin'
          Command.create(this.temp).then(() => {
            Message.info('add success')
            this.$refs.edit_table.getList()
            this.dialogFormVisible = false
          }).catch(error => {
            Message.error({ message: error })
            console.log('add error:', error)
            Message.error({ message: error })
            this.dialogFormVisible = false
          })
        }
      })
    }
  }
}
</script>
