<template>
  <div>
    <up-load
      @reLoadList="reLoadList"
    />

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
import Command from '@/api/plot/upLoad'
import { Message } from 'element-ui'
import ComplexTable from '@/views/goplot/components/complex-table'
import Setting from '@/views/goplot/upLoadFile/components/setting'
import UpLoad from '@/views/goplot/components/upLoad'

export default {
  name: 'TableData',
  components: {
    UpLoad,
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
          width: 80,
          align: 'center',
          valign: 'middle',
          sort: true
        },
        {
          field: 'file_name',
          title: '文件名称',
          width: 240,
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
          field: 'file_class',
          title: '分类',
          width: 100,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          },
          editable: {
            type: 'select',
            title: '部门',
            mode: 'inline',
            source: () => { return [{ value: '.sh', label: '.sh' }, { value: '.tar.gz', label: '.tar.gz' }] }
          }
        },
        {
          field: 'version',
          title: '版本',
          width: 100,
          align: 'center',
          valign: 'middle',
          editable: {
            type: 'text',
            title: '部门',
            mode: 'inline'
          }
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
    getGameNameBigSource (row) {
      const source = []
      this.serverBigSource.forEach(element => {
        source.push({ value: element.name, label: element.name })
      })
      return source
    },
    getGameNameSource (row) {
      const source = []
      this.serverSource.forEach(element => {
        if (row.server_name_big === element.group_name) {
          source.push({ value: element.name, label: element.name })
        }
      })
      return source
    },
    getTaskSource (row) {
      const source = []
      this.taskSource.forEach(element => {
        source.push({ value: element.name, label: element.name })
      })
      return source
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
