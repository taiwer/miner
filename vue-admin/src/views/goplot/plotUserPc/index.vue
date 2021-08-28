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
import TableData from './components/tableData'

export default {
  name: 'UserNodeList',
  components: {
    TableData
  },
  data () {
    return {
      account: ''
    }
  },
  created () {
  },
  methods: {
  }
}
</script>
