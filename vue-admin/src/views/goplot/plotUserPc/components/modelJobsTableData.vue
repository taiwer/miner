<template>
  <div>
    <el-row :gutter="10" style="margin-top: 20px">
      <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
        <!--        <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">-->
        <!--          Add1-->
        <!--        </el-button>-->
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
        :get-data="getJobList"
        :fiter="fiter"
        :table-data="undefined"
        :header-cell-style="{background:'#F4F5F6',color:'#131D34',padding: '8px 0'}"
        @envEditableSave="onEditableSave"
        @envActionHandle="onActionHandle"
      >
        <!--        工具插槽-->
        <template v-slot:toolbar style="padding: 0 0">
          <el-button class="filter-item" @click="reLoadList">刷新</el-button>
          <el-button class="filter-item" @click="createPlot">新建批图任务</el-button>
        </template>
        <!--        行扩展插槽-->
        <template v-slot:expand style="padding: 0 0">
          <el-table-column type="expand" cell-class-name="table_expand" />
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
import Command from '@/api/plot/plotUserPc'
import { Message } from 'element-ui'
import ComplexTable from '@/views/goplot/components/complex-table'
import Setting from './setting'

export default {
  name: 'JobsTableData',
  components: {
    Setting,
    ComplexTable
  },
  props: {
    row: Object
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
      tableData: undefined,
      columns: [
        {
          field: 'pid',
          title: 'pid',
          width: 80,
          align: 'center',
          valign: 'middle',
          sort: true
        },
        {
          field: 'plot_id',
          title: 'plot_id',
          width: 80,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'start_time_str',
          title: 'start_time_str',
          width: 80,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'size',
          title: 'size',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'n_threads',
          title: 'n_threads',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'n_buckets',
          title: 'n_buckets',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'job_buff',
          title: 'job_buff',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'tmpdir',
          title: 'tmpdir',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'dstdir',
          title: 'dstdir',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'phase',
          title: 'phase',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value + '/' + row.sub_phase
          }
        },
        {
          field: 'tmp_usage',
          title: 'tmp_usage',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'dst_usage',
          title: 'dst_usage',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'mem_usage',
          title: 'mem_usage',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'wall_time',
          title: 'wall_time',
          min_width: 20,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'action',
          title: '操作',
          width: 80,
          align: 'center',
          valign: 'middle',
          buttons: [
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
      if (this.row.name !== '') {
        result.push('node_id=\'' + this.row.name + '\'')
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
  mounted () {
  },
  created () {
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
          const data = this.getRequestParam()
          Command.send_kill(data, row.plot_id).then(response => {
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
    createPlot () {
      const data = this.getRequestParam()
      Command.send_plot(data).then(response => {
        Message.info('success')
        console.log('delete response:', response)
        this.$refs.edit_table.getList() /* 重新刷新 */
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    },
    getRequestParam () {
      const data = {}
      data.ploter_name = this.row.name
      data.user_name = this.row.user_name
      return data
    },
    getJobList () {
      const data = this.getRequestParam()
      return Command.get_job_list(data)
    }
  }
}
</script>

<style >
.el-table__expanded-cell[class*="cell"]{
  padding: 0 0px;
}
</style>
