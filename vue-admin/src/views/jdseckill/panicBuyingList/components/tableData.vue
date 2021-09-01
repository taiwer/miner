<template>
  <div>
    <div>
      <el-row :gutter="10" style="margin-top: 20px">
        <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
            Add
          </el-button>
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleGetMiaoshaList">
            getMiaoshaList
          </el-button>
        </el-col>
        <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
          <div />
        </el-col>
      </el-row>
      <ComplexTable
        ref="edit_table"
        :can-edit="editModeEnabled"
        :can-search="editModeEnabled"
        :columns="columns"
        :get-data="getList"
        :tableData = "tableData"
        :fiter="fiter"
        :header-cell-style="{background:'#F4F5F6',color:'#131D34',padding: '8px 0'}"
        @envEditableSave="onEditableSave"
        @envActionHandle="onActionHandle"
      >
        <template v-slot:expand style="padding: 0 0">
          <el-table-column type="expand" cell-class-name="table_expand">
            <template slot-scope="{row}">
              <el-collapse>
                <el-collapse-item title="账号控制功能" name="1">
                  <settings
                    :row="row"
                  />
                </el-collapse-item>
              </el-collapse>
              <div class="tab-container">
                <el-tabs v-model="activeName" style="margin-top:15px;" type="border-card">
                  <el-tab-pane key="tb-data" label="角色列表" name="tb-data">
                    <charactersTableData
                      :account="row.account"
                      :server-name="row.server_name"
                      :task-name="row.task"
                    />
                  </el-tab-pane>
                  <el-tab-pane key="tb-data2" label="登录" name="tb-data2" />
                </el-tabs>
              </div>
              {{ row["task"] }}
            </template>
          </el-table-column>
        </template>
      </ComplexTable>
    </div>
    <div>
      <setting
        :edit-data="editData"
        :edit-form="editForm"
        @reLoadList="reLoadList"
      />
    </div>
  </div>
</template>

<script>
import Command from '@/api/jdseckill/panicBuyingList'
import { Message } from 'element-ui'
import ComplexTable from '@/views/goplot/components/complex-table'
import Setting from './setting'

export default {
  name: 'TableData',
  components: {
    Setting,
    ComplexTable
  },
  props: {
    account: String
  },
  data () {
    return {
      editForm: {
        Visible: false,
        dialogStatus: '',
        textMap: {
          update: 'Edit',
          create: 'Create'
        }
      },
      editData: {}, // 编辑数据
      activeName: 'tb-data',
      editModeEnabled: true,
      getList: Command.get_list,
      tableData: [],
      columns: [
        {
          field: 'id',
          title: 'Id',
          width: 180,
          align: 'center',
          valign: 'middle',
          sort: true,
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'user_name',
          title: 'user_name',
          width: 180,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'item_id',
          title: 'item_id',
          width: 180,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'item_name',
          title: 'item_name',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'num',
          title: 'num',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'limit_price',
          title: 'limit_price',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'enable',
          title: 'enable',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'start_at',
          title: 'start_at',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'stop_at',
          title: 'stop_at',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'action',
          title: '操作',
          width: 180,
          align: 'center',
          valign: 'middle',
          buttons: [
            {
              type: 'primary',
              size: 'mini',
              action: 'edit',
              title: '编辑'
            },
            {
              type: 'danger',
              size: 'mini',
              action: 'delete',
              title: '删除'
            }
          ]
        }
      ],
      serverBigSource: [],
      serverSource: [],
      taskSource: []
    }
  },
  computed: {
    fiter () {
      console.log('fiter')
      const result = []
      if (this.account !== '') {
        result.push('account=\'' + this.account + '\'')
      }
      if (result.length > 0) {
        const res = result.join(' and ')
        console.log(res)
        return res
      } else {
        return ''
      }
    }
  },
  watch: {
    activeName (val) {
      this.$router.push(`${this.$route.path}?tab=${val}`)
    }
  },
  created () {
    console.log('created:')
    if (this.account !== '') {
      this.fiter = 'account=\'' + this.account + '\''
    }
  },
  methods: {
    onExpandRow: function (index, row, $detail) {
    },
    onClickCell: function (field, value, row, element) {
    },
    onEditableSave: function (field, row, oldValue, callback) {
      console.log('index onEditableSave', field, row, oldValue)
      const data = {}
      data.id = row.id
      data.name = row.name
      data[field] = row[field]
      Command.update(data).then(response => {
        console.log('onEditableSave response:', response)
        callback(true)
      }).catch(error => {
        console.log('onEditableSave catch:', error)
        callback(false)
      })
    },
    onActionHandle: function (row, index, action) {
      console.log('index onActionHandle', row, index, action)
      switch (action) {
        case 'edit': {
          this.editData = row
          this.editForm.dialogStatus = 'update'
          this.editForm.Visible = true
          this.$nextTick(() => {
            // this.$refs['dataForm'].clearValidate()
          })
          break
        }
        case 'delete': {
          if (confirm('确认删除？') !== true) {
            return
          }
          const data = {}
          data.id = row.id
          Command.del(data).then(response => {
            Message.info('success')
            console.log('delete response:', response)
            this.$refs.edit_table.getList() /* 重新刷新 */
          }).catch(error => {
            console.log('delete error:', error)
            Message.error({ message: error })
          })
          break
        }
        default:
      }
    },
    reLoadList () {
      this.$refs.edit_table.getList()
    },
    resetTemp () {
      this.editData = {
      }
    },
    handleCreate () {
      this.resetTemp()
      this.editForm.dialogStatus = 'create'
      this.editForm.Visible = true
    },
    handleGetMiaoshaList () {
      const data = {}
      Command.send_get_miaosha_list(data).then(response => {
        Message.info('success' + response.data.length)
        this.$refs.edit_table.tableData = response.data
        this.$refs.edit_table.getList() /* 重新刷新 */
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
