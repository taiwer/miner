<template>
  <div style="margin-top: 3px;">
    <el-input v-model="inputData" :readonly="readonly" placeholder="Please input" style="width:90%;max-width:100%;">
      <template v-show="title" slot="prepend">
        {{ title }}
      </template>
    </el-input>
    <el-button type="primary" icon="el-icon-document" @click="handleCopy(inputData,$event)">
      copy
    </el-button>
  </div>
</template>

<script>
import clip from '@/utils/clipboard' // use clipboard directly
import clipboard from '@/directive/clipboard/index.js' // use clipboard by v-directive

export default {
  name: 'Clipboard',
  props: {
    inputData: String,
    readonly: Boolean,
    title: String
  },
  data () {
    return {}
  },
  methods: {
    handleCopy (text, event) {
      clip(text, event)
    },
    clipboardSuccess () {
      this.$message({
        message: 'Copy successfully',
        type: 'success',
        duration: 1500
      })
    }
  }
}
</script>

<style scoped>

</style>
