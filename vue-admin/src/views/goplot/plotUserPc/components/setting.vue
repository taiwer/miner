<template>
  <el-dialog :title="editForm.textMap[editForm.dialogStatus]" :visible.sync="editForm.Visible" style="min-width: 1000px">
    <el-form ref="dataForm" :model="editData" label-position="left" label-width="100px" style="margin-left:10px">
      <el-form-item label="Id" prop="title">
        <el-col :span="4">
          <el-sl-panel>{{ editData.id }}</el-sl-panel>
        </el-col>
        <el-col class="line" :span="2">plotId</el-col>
        <el-col :span="4">
          <el-sl-panel>{{ editData.name }}</el-sl-panel>
        </el-col>
        <el-col class="line" :span="2">Version</el-col>
        <el-col :span="4">
          <el-sl-panel>{{ editData.version }}</el-sl-panel>
        </el-col>
        <el-col class="line" :span="2">StartAt</el-col>
        <el-col :span="4">
          <el-sl-panel>{{ editData.start_at }}</el-sl-panel>
        </el-col>
      </el-form-item>
      <el-form-item label="备注名称" prop="title">
        <el-col :span="4">
          <el-input v-model="editData.descs" />
        </el-col>
      </el-form-item>
      <el-form-item label="KSize" prop="title">
        <el-col :span="4">
          <el-select v-model="editData.plot_size" class="filter-item" placeholder="please select">
            <el-option key="plot_size_32" label="32" :value="32" />
            <el-option key="plot_size_33" label="33" :value="33" />
          </el-select>
        </el-col>
        <el-col class="line" :span="2">thread</el-col>
        <el-col :span="2">
          <el-select v-model="editData.plot_thread" class="filter-item" placeholder="please select">
            <el-option key="plot_thread_1" label="1" :value="1" />
            <el-option key="plot_thread_2" label="2" :value="2" />
            <el-option key="plot_thread_3" label="3" :value="3" />
            <el-option key="plot_thread_4" label="4" :value="4" />
            <el-option key="plot_thread_6" label="6" :value="6" />
            <el-option key="plot_thread_8" label="8" :value="8" />
            <el-option key="plot_thread_8" label="16" :value="16" />
            <el-option key="plot_thread_8" label="32" :value="32" />
            <el-option key="plot_thread_8" label="64" :value="64" />
          </el-select>
        </el-col>
        <el-col :span="2">
          <el-input v-model="editData.plot_thread" placeholder="4" />
        </el-col>
        <el-col class="line" :span="2">bucket</el-col>
        <el-col :span="4">
          <el-select v-model="editData.plot_bucket" class="filter-item" placeholder="please select">
            <el-option key="plot_bucket_32" label="32" :value="32" />
            <el-option key="plot_bucket_64" label="64" :value="64" />
            <el-option key="plot_bucket_128" label="128" :value="128" />
            <el-option key="plot_bucket_256" label="256" :value="256" />
            <el-option key="plot_bucket_512" label="512" :value="512" />
          </el-select>
        </el-col>
        <el-col class="line" :span="2">job_memory </el-col>
        <el-col :span="4">
          <el-select v-model="editData.plot_memory" class="filter-item" placeholder="please select">
            <el-option key="plot_bucket_4g" label="4G" :value="4096" />
            <el-option key="plot_bucket_6g" label="6G" :value="6144" />
            <el-option key="plot_bucket_8g" label="8G" :value="8192" />
            <el-option key="plot_bucket_16g" label="16G" :value="16384" />
          </el-select>
        </el-col>
      </el-form-item>
      <el-form-item label="JobTmp(GB)">
        <el-col :span="4">
          <el-input v-model="editData.plot_job_tmp_size" placeholder="330" />
        </el-col>
        <el-col class="line" :span="2">P图间隔(分)</el-col>
        <el-col :span="4">
          <el-input v-model="editData.plot_interval" placeholder="5" />
        </el-col>
        <el-col class="line" :span="2">Dst并发数</el-col>
        <el-col :span="4">
          <el-input v-model="editData.plot_dst_work_num" placeholder="4" />
        </el-col>
        <el-col class="line" :span="2">最大任务</el-col>
        <el-col :span="4">
          <el-input v-model="editData.plot_max_job" placeholder="4" />
        </el-col>
      </el-form-item>
      <el-form-item label="farmer_key">
        <el-input v-model="editData.plot_farmer_public_key" placeholder="farmer_public_key" />
      </el-form-item>
      <el-form-item label="pool_key">
        <el-input v-model="editData.plot_pool_public_key" placeholder="pool_public_key" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="editForm.Visible = false">
        Cancel
      </el-button>
      <el-button type="primary" @click="editForm.dialogStatus==='create'?createData():updateData()">
        Confirm
      </el-button>
    </div>
  </el-dialog>
</template>

<script>
import Command from '@/api/plot/plotUserPc'
import { Message } from 'element-ui'

export default {
  name: 'Setting',
  props: {
    editData: Object,
    editForm: Object
  },
  data () {
    return {
      serverName: '',
      instanceNameSource: ['Saab', 'Volvo', 'BMW']
    }
  },
  created () {
  },
  methods: {
    createData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          Command.create(this.editData).then(() => {
            Message.info('add success')
            this.$emit('reLoadList')
            // this.editForm.Visible = false
          }).catch(error => {
            console.log('add error:', error)
            Message.error({ message: error })
            this.editForm.Visible = false
          })
        }
      })
    },
    updateData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          console.log('setting', this.editData)
          Command.update(this.editData).then(() => {
            Message.info('add success')
            this.$emit('reLoadList')
            this.editForm.Visible = false
          }).catch(error => {
            Message.error({ message: error })
            console.log('add error:', error)
            Message.error({ message: error })
            this.editForm.Visible = false
          })
        }
      })
    }
  }
}
</script>

<style>
div .el-dialog {
  width: 70%;
}
.el-form-item {
  margin-bottom: 5px;
}
</style>
