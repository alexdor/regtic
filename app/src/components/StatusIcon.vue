<template>
  <div>
    <el-tooltip v-if="data.type == 'good'" class="item" effect="dark" content="Good" placement="top">
      <div :class="'status-icon ' + data.type" :aria-label="data.type"></div>
    </el-tooltip>
    <el-tooltip v-if="data.type != 'good'" class="item" effect="dark" placement="top">
      <div slot="content"><li v-for="list in data.lists" v-bind:key="list.url">{{list.name}} | <a :href="list.url">{{list.url}}</a></li>
      <br v-if="data.lists.length > 0 && data.notes.length > 0" />
      <br v-if="data.lists.length > 0 && data.notes.length > 0" />
      {{data.notes}}</div>
      <div :class="'status-icon ' + data.type"></div>
    </el-tooltip>
  </div>
</template>

<script>
  export default {
    props: {
      data: {
        type: Object
      }
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

    .status-icon.good {
      background-color: $green;
      background-image: url('../assets/icon-good.svg');
    }

    .status-icon.warning {
      background-color: $orange;
      background-image: url('../assets/icon-warning.svg');
    }

    .status-icon.bad {
      background-color: $red;
      background-image: url('../assets/icon-bad.svg');
    }

    li {
      list-style: none;
    }

    li a {
      color: $link-color;
    }
</style>
