<template>
  <div class="edit-cell" @click="onFieldClick">
    <div
      v-if="!editMode && !showInput"
      tabindex="0"
      class="cell-content"
      :class="{'edit-enabled-cell': canEdit}"
      @keyup.enter="onFieldClick"
    >
      <span>{{ model }}</span>
    </div>
    <component
      :is="editableComponent"
      v-if="editMode || showInput"
      ref="input"
      v-model="model"
      v-bind="$attrs"
      @focus="onFieldClick"
      @keyup.enter.native="onInputExit"
      v-on="listeners"
    >
      <template v-if="editableComponent=='el-select'">
        <template v-for="item in sourceData">
          <el-option :key="item.label" :value="item.value" :label="item.label" />
        </template>
      </template>
    </component>
  </div>
</template>
<script>

import { sleep } from '@/utils/auth'

export default {
  name: 'EditableCell',
  inheritAttrs: false,
  props: {
    value: {
      type: Object,
      default: () => {}
    },
    row: {
      type: Object,
      default: () => {}
    },
    field: {
      type: String,
      default: ''
    },
    onEditableSave: {
      type: Function,
      default: () => {}
    },
    source: {
      type: Function,
      // eslint-disable-next-line vue/require-valid-default-prop
      default: () => {}
    },
    showInput: {
      type: Boolean,
      default: false
    },
    editableComponent: {
      type: String,
      default: 'el-input'
    },
    closeEvent: {
      type: String,
      default: 'blur'
    },
    canEdit: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      editMode: false,
      oldValue: undefined,
      blur: false
    }
  },
  computed: {
    model: {
      get () {
        return this.row[this.field]
      },
      set (val) {
        this.row[this.field] = val
      }
    },
    listeners () {
      if (this.editableComponent === 'el-select') {
        // eslint-disable-next-line vue/no-side-effects-in-computed-properties
        // this.closeEvent = 'change'
        // eslint-disable-next-line vue/no-side-effects-in-computed-properties
        // this.closeEvent = 'deactivate'
        return {
          'blur': () => {
            // eslint-disable-next-line vue/no-side-effects-in-computed-properties
            this.blur = true
            sleep(200).then(() => {
              if (this.blur === true) {
                this.onInputExit()
              }
            })
          },
          'focus': () => {
            // eslint-disable-next-line vue/no-side-effects-in-computed-properties
            this.blur = false
          },
          'change': this.onInputExit,
          ...this.$listeners
        }
      }
      return {
        [this.closeEvent]: this.onInputExit,
        ...this.$listeners
      }
    },
    sourceData: {
      get () {
        return this.source(this.row)
      }
    }
  },
  created () {
    console.log('source', this.source)
  },
  methods: {
    onFieldClick () {
      if (this.canEdit) {
        this.editMode = true
        this.oldValue = this.row[this.field]
        this.onGetSelectSource()
        this.$nextTick(() => {
          const inputRef = this.$refs.input
          if (inputRef && inputRef.focus) {
            inputRef.focus()
          }
        })
      }
    },
    onInputExit () {
      if (this.row[this.field] === this.oldValue) {
        console.log('notChange', this.row[this.field], this.oldValue)
      } else {
        console.log('Changed', this.row[this.field], this.oldValue)
        this.$emit('onEditableSave', this.field, this.row, this.oldValue, success => {
          console.log(success)
          if (!success) {
            console.log('onEditableSave field')
            this.row[this.field] = this.oldValue
          } else {
            console.log('onEditableSave success')
          }
        })
      }
      this.editMode = false
    },
    onInputChange (val) {
      console.log('onInputChange:' + val)
      this.$emit('input', val)
    }
  }
}
</script>
<style>
.edit-cell{
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  word-break: break-all;
  line-height: 23px;
  padding-left: 0px;
  padding-right: 0px;
}
.cell-content {
  #min-height: 30px;
  #padding-left: 5px;
  #padding-top: 5px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  border: 1px solid transparent;
}
.edit-enabled-cell {
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  word-break: break-all;
  line-height: 23px;
  padding-left: 10px;
  padding-right: 10px;
  border: 1px dashed #409eff;
}
.edit-cell
.el-input__inner {
  height: 30px;
  line-height: 30px;
  padding: 0 0px;
}
</style>
