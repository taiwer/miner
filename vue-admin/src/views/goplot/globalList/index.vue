<template xmlns="http://www.w3.org/1999/html">
  <div class="dashboard-editor-container">
    <div class="chart-wrapper">
      <div style="padding: 15px">
        <table-data
          :account="account"
        />
      </div>
    </div>
  </div>
</template>

<script>
import Command from '@/api/plot/globalSetting'
import { Message } from 'element-ui'
import TableData from './components/tableData'

export default {
  name: 'AccountTask',
  components: {
    TableData
  },
  data () {
    return {
      account: '',
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
