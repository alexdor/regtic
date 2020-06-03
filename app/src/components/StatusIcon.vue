<template>
  <div>
    <el-tooltip v-if="status == 'OK'" class="item" effect="dark" content="Not detected on any list" placement="top">
      <div :class="'status-icon ' + status" :aria-label="status"></div>
    </el-tooltip>
    <el-tooltip v-if="status != 'OK'" class="item" effect="dark" placement="top">
      <div slot="content">
        <a :href="source">{{source}}</a>
        <br v-if="source != undefined && statusNotes != undefined" />
        <br v-if="source != undefined && statusNotes != undefined" />
        {{statusNotes}}
      </div>
      <div :class="'status-icon ' + status"></div>
    </el-tooltip>
  </div>
</template>

<script>
  export default {
    props: {
      status: String,
      statusType: String,
      source: String,
      statusNotes: String
    }
  }
</script>

<style lang="scss" scoped>
  $green: #67C23A;
  $orange: #E6A23C;
  $red: #F56C6C;

  $link-color: white;

  .status-icon {
    border-radius: 50%;
    width: 1.375rem;
    height: 1.375rem;
    background-repeat: no-repeat;
    background-position: center;
  }

    .status-icon.OK {
      background-color: $green;
      background-image: url('../assets/icon-good.svg');
    }

    .status-icon.WARNING {
      background-color: $orange;
      background-image: url('../assets/icon-warning.svg');
    }

    .status-icon.ISSUE {
      background-color: $red;
      background-image: url('../assets/icon-bad.svg');
    }

    a {
      color: $link-color;
    }
</style>
