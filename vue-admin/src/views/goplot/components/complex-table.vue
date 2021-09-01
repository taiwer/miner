<template>
  <div class="app-container" style="padding:0 0">
    <div class="filter-container">
      <el-row
        v-if="canEdit || canSearch"
        :gutter="10"
        style="text-align: left;line-height: 40px"
      >
        <el-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4">
          <div style="margin-bottom: 10px">
            <el-switch
              v-if="canEdit"
              v-model="editModeEnabled"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="开启"
              inactive-text="禁用"
            />
          </div>
        </el-col>
        <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
          <div style="margin-bottom: 10px">
            <slot name="toolbar" />
          </div>
        </el-col>
        <el-col
          v-if="canSearch"
          :xs="8"
          :sm="8"
          :md="8"
          :lg="8"
          :xl="8"
        >
          <div style="float: right">
            <el-input v-model="listQuery.search" placeholder="Search" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
              Search
            </el-button>
          </div>
        </el-col>
      </el-row>

    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;line-height: 34px"
      @sort-change="sortChange"
      @selection-change="handleSelectionChange"
    >
      >
      <slot name="expand" />
      <template v-for="item in columns">
        <template v-if="item.type==='selection'">
          <el-table-column
            :type="item.type"
            :width="item.width"
          />
        </template>
        <template v-else slot-scope="{row,$index}">
          <el-table-column
            :key="item.title"
            :type="item.type"
            :label="item.title"
            :prop="item.field"
            :width="item.width"
            :min-width="item.min_width"
            :align="item.align"
            :valign="item.valign"
            :class-name="item.sort?getSortClass(item.field):undefined"
            :sortable="item.sort?'custom':undefined"
          >
            <template v-if="item.type===undefined" slot-scope="{row,$index}">
              <template v-if="item.buttons === undefined">
                <span v-if="item.editable === undefined">
                  {{ item.formatter===undefined?row[item.field]:item.formatter(row[item.field], row, $index) }}
                </span>
                <editable-cell
                  v-else
                  slot-scope="{row}"
                  :row="row"
                  :field="item.field"
                  :can-edit="editModeEnabled"
                  :editable-component="getEditableType(item.editable.type)"
                  :source="item.editable.source"
                  @onEditableSave="onEditableSave"
                />
              </template>
              <template v-else slot-scope="{row,$index}">
                <el-button
                  v-for="item_btn in item.buttons"
                  :key="item_btn.label"
                  :type="item_btn.type"
                  :size="item_btn.size"
                  @click="onActionHandle(row, $index, item_btn.action)"
                >
                  {{ item_btn.title }}

                </el-button>
              </template>
            </template>
          </el-table-column>
        </template>

      </template>
    </el-table>
    <pagination v-show="total>listQuery.limit" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination
import EditableCell from './editEableCell'
import { Message } from 'element-ui'
export default {
  name: 'ComplexTable',
  components: { Pagination, EditableCell },
  directives: { waves },
  props: {
    canEdit: Boolean,
    canSearch: Boolean,

    envEditableSave: {
      type: Function,
      default: () => {}
    },
    envActionHandle: {
      type: Function,
      default: () => {}
    },
    getData: {
      type: Function,
      default: undefined
    },
    fiter: String,
    columns: {
      type: Array,
      default: () => {
        return [{
          field: 'id',
          title: 'id',
          width: 20,
          min_width: 20,
          align: 'center',
          valign: 'middle',
          formatter: function (value, row, index) {
            return value
          },
          buttons: [
            {
              type: 'primary',
              size: 'mini',
              action: 'edit',
              title: '编辑'
            }
          ],
          editable: {
            type: 'select',
            title: '部门',
            mode: 'inline',
            source: Function
          }
        }]
      }
    },
    tableData: {
      type: Array,
      default: undefined
    }
  },
  data () {
    return {
      editModeEnabled: false,
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        offset: 0,
        search: '',
        fiter: '',
        sort: '+id'
      },
      importanceOptions: [1, 2, 3],
      showReviewer: false,
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogPvVisible: false,
      pvData: [],
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      downloadLoading: false,
      multipleSelection: []
    }
  },
  created () {
    console.log('getList')
    this.getList()
  },
  methods: {
    getList () {
      console.log('this.tableData', this.tableData)
      if (this.getData !== undefined) {
        console.log('do getData')
        this.listQuery.fiter = this.fiter
        this.listLoading = true
        this.listQuery.offset = this.listQuery.limit * this.listQuery.page - this.listQuery.limit
        this.getData(this.listQuery).then(response => {
          console.log(response)
          this.list = response.data.Rows
          this.total = response.data.Total
          this.list = this.list.map(v => {
            this.$set(v, 'edit', false) // https://vuejs.org/v2/guide/reactivity.html
            return v
          })
          this.listLoading = false
        }).catch(error => {
          this.list = []
          this.total = 0
          Message.error({ message: error })
          this.listLoading = false
        })
        return
      }
      if (this.tableData !== undefined) {
        console.log('do tableData ' + this.tableData.length)
        this.listLoading = true
        this.list = this.tableData
        this.total = this.tableData.length
        this.list = this.list.map(v => {
          this.$set(v, 'edit', false)
          return v
        })
        Message.info('success ' + this.tableData.length)
        this.listLoading = false
        return
      }
    },
    onExpandRow: function (index, row, $detail) {
    },
    onClickCell: function (field, value, row, element) {
    },
    onEditableSave: function (field, row, oldValue, callback) {
      console.log('onEditableSave', field, row, oldValue)
      this.$emit('envEditableSave', field, row, oldValue, success => {
        console.log(success)
        if (!success) {
          callback(false)
        } else {
          callback(true)
        }
      })
    },
    onActionHandle: function (row, index, action) {
      console.log('onActionHandle', row, index, action)
      this.$emit('envActionHandle', row, index, action)
    },
    handleFilter () {
      this.listQuery.page = 1
      this.getList()
    },
    handleModifyStatus (row, status) {
      this.$message({
        message: '操作Success',
        type: 'success'
      })
      row.status = status
    },
    handleSelectionChange (val) {
      console.log('handleSelectionChange', val)
      this.multipleSelection = val
    },
    sortChange (data) {
      const { prop, order } = data
      console.log('sortChange', data)
      if (order === 'ascending') {
        this.listQuery.sort = `+${prop}`
      } else {
        this.listQuery.sort = `-${prop}`
      }
      console.log(this.listQuery.sort)
      this.handleFilter()
    },
    getSortClass: function (key) {
      const sort = this.listQuery.sort
      return sort === `+${key}` ? 'ascending' : 'descending'
    },
    getEditableType: function (type) {
      if (type === 'text') {
        return 'el-input'
      }
      if (type === 'select') {
        return 'el-select'
      }
      return 'el-input'
    }
  }
}
</script>

<style scoped>
.el-table--medium th, .el-table--medium td {
  padding: 0px 0;
}
</style>
