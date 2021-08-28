<template>
  <div>
    <el-row :gutter="10" style="margin-top: 20px">
      <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
        <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
          Add
        </el-button>
      </el-col>
      <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
        <div />
      </el-col>
    </el-row>
    <div style="padding: 15px">
      <ComplexTable
        ref="edit_table"
        :can-edit="editModeEnabled"
        :can-search="editModeEnabled"
        :columns="columns"
        :get-data="getList"
        :fiter="fiter"
        :header-cell-style="{background:'#F4F5F6',color:'#131D34',padding: '8px 0'}"
        @envEditableSave="onEditableSave"
        @envActionHandle="onActionHandle"
      >
        <!--        工具插槽-->
        <template v-slot:toolbar style="padding: 0 0">
          <el-button class="filter-item" @click="reLoadList">查询</el-button>
        </template>
        <!--        行扩展插槽-->
        <template v-slot:expand style="padding: 0 0">
          <el-table-column type="expand" cell-class-name="table_expand">
            <template slot-scope="{row}">
              <div class="tab-container">
                <el-collapse>
                  <Clipboard
                    title="安装命令"
                    :input-data="row.install_node_shell"
                    :readonly="true"
                  />
                  <div style="margin-top: 3px;">
                    <el-input v-model="shellCommand" :readonly="readonly" placeholder="Please input" style="width:90%;max-width:100%;">
                      <template v-show="title" slot="prepend">
                        运行脚本
                      </template>
                    </el-input>
                    <el-button type="primary" icon="el-icon-document" @click="handleRunShell(row.name, shellCommand,$event)">
                      run shell
                    </el-button>
                  </div>
                </el-collapse>
                <el-tabs v-model="activeName" style="margin-top:15px;" type="border-card">
                  <el-tab-pane key="tb-data" label="任务列表" name="tb-data">
                    <JobsTableData
                      :row="row"
                    />
                  </el-tab-pane>
                  <el-tab-pane key="tb-data2" label="磁盘列表" name="tb-data2">
                    <PlotDiskTable
                      :node_id="row.name"
                    />
                  </el-tab-pane>
                </el-tabs>
              </div>
            </template>
          </el-table-column>
        </template>
      </ComplexTable>
    </div>

    <Setting
      :edit-data="editData"
      :edit-form="editForm"
      @reLoadList="reLoadList"
    />
  </div></template>

<script>
import Command from '@/api/plot/plotPc'
import { Message } from 'element-ui'
import ComplexTable from '@/views/goplot/components/complex-table'
import Setting from './setting'
import JobsTableData from './modelJobsTableData'
import PlotDiskTable from '@/views/goplot/plotDisk/components/tableData'
import Clipboard from '@/views/goplot/components/clipboard'

export default {
  name: 'TableData',
  components: {
    JobsTableData,
    PlotDiskTable,
    Setting,
    ComplexTable,
    Clipboard
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
          update: '编辑',
          create: '新建'
        }
      },
      editData: {}, // 编辑数据
      activeName: 'tb-data',
      editModeEnabled: true,
      getList: Command.get_list,
      columns: [
        {
          field: 'id',
          title: 'id',
          width: 50,
          align: 'center',
          valign: 'middle',
          sort: true
        },
        {
          field: 'name',
          title: '名称',
          width: 200,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          },
          editable: {
            type: 'text',
            title: '部门',
            mode: 'inline'
          }
        },
        {
          field: 'online',
          title: '在线',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'descs',
          title: '描述',
          width: 140,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'user_name',
          title: '用户名',
          width: 150,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'ip',
          title: 'IP地址',
          width: 120,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'version',
          title: '版本号',
          width: 100,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'ruing_job_count',
          title: '运行',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_count_today',
          title: '今日',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_count_yesterday',
          title: '昨日',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_size',
          title: 'K',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_thread',
          title: '线程',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_bucket',
          title: '块',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_memory',
          title: '内存',
          width: 60,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_interval',
          title: '任务间隔',
          min_width: 20,
          align: 'center',
          valign: 'middle'
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
      shellCommand: ''
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
          data.name = row.name
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
    handleRunShell (name, shellCommand) {
      const data = {}
      data.ploter_name = name
      Command.send_run_shell(data, shellCommand).then(response => {
        Message.info('success')
        console.log('delete response:', response)
        this.$refs.edit_table.getList() /* 重新刷新 */
      }).catch(error => {
        console.log('onEditableSave catch:', error)
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
