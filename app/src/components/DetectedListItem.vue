<template>
  <tr class="detected-list-item">
    <td class="value"><a :href="source" target="_blank">{{source}}</a></td>
    <td><span class="key">Type:</span> <span class="value">{{type}}</span></td>
    <td><span class="key">Updated:</span> <span class="value">{{updated}}</span></td>
    <td><StatusIcon :status="status" :statusType="type" :statusNotes="status" class="pull-right"></StatusIcon></td>
  </tr>
</template>

<script>
  import StatusIcon from "../components/StatusIcon.vue";

  export default {
    components: {
      StatusIcon
    },
    props: {
      source: String,
      type: String,
      updated: String,
      status: String
    }
  }
</script>

<style lang="scss" scoped>
  $black: #303133;

  $green: #67C23A;
  $orange: #E6A23C;
  $red: #F56C6C;

  $link-hover: #1871F8;

  .detected-list-item {
    font-family: 'Poppins', sans-serif;
    width: 100%;
  }

  td {
    padding: 0.5rem 0;
  }

  a {
    color: $black;
    text-decoration: none;
  }

  a:hover {
    color: $link-hover;
  }

  .pull-right {
    float: right;
  }

  .value {
    color: $black;
  }
</style>
