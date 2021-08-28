<template>
  <div>
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
            <template slot-scope="{row}" />
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
import Command from '@/api/plot/plotDisk'
import Conversion from '@/utils/conversion'
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
    nodeId: String
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
          field: 'node_id',
          title: '节点名称',
          width: 180,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'mount',
          title: 'Mount',
          width: 150,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'priority',
          title: 'priority',
          width: 70,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'disk_type',
          title: 'disk_type',
          width: 70,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'size',
          title: 'Size',
          width: 50,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return Conversion.conver(value)
          }
        },
        {
          field: 'free_size',
          title: 'Free',
          width: 60,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return Conversion.conver(value)
          }
        },
        {
          field: 'used_percent',
          title: 'Used',
          width: 60,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          }
        },
        {
          field: 'plot_count',
          title: 'plot_count',
          width: 50,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_max_count',
          title: 'plot_max_count',
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
          width: 80,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_farmer_public_key',
          title: 'farmer_key',
          width: 80,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'plot_pool_public_key',
          title: 'pool_key',
          width: 80,
          align: 'center',
          valign: 'middle'
        },
        {
          field: 'descs',
          title: '描述',
          min_width: 20,
          align: 'center',
          valign: 'middle',
          editable: {
            type: 'text',
            title: '部门',
            mode: 'inline'
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
      if (this.node_id !== undefined) {
        result.push('node_id=\'' + this.node_id + '\'')
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
    }
  }
}
</script>

<style >
.el-table__expanded-cell[class*="cell"]{
  padding: 0 0px;
}
</style>
