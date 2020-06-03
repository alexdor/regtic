<template>
  <div class="status-bar">
    <div class="ok" :style="'flex:' + ok"></div>
    <div class="warning" :style="'flex:' + warning"></div>
    <div class="issue" :style="'flex:' + issue"></div>
  </div>
</template>

<script>
  export default {
    props: {
      ok: Number,
      warning: Number,
      issue: Number
    }
  }
</script>

<style lang="scss" scoped>
  $green: #67C23A;
  $orange: #E6A23C;
  $red: #F56C6C;

  .status-bar {
    width: 200px;
    height: 4px;
    display: flex;
    flex-direction: row;
  }

  .status-bar div {
    border-radius: 2px;
    margin: 0 1px;
  }

  .ok {
    background-color: $green;
  }

  .warning {
    background-color: $orange;
  }

  .issue {
    background-color: $red;
  }
</style>
