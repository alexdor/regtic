<template>
  <el-row>
    <el-col :span="3" class="center-icon">
      <i
        :class="
          'el-notification-icon el-icon-' +
          notificationTypes[data.type].icon +
          ' el-color-' +
          notificationTypes[data.type].color
        "
      ></i>
    </el-col>
    <el-col :span="21">
      <p>
        <b>{{ data.boldText }}</b>
        {{ data.text }}
      </p>
    </el-col>
  </el-row>
</template>
<script>
export default {
  props: {
    data: {
      type: Object,
    },
  },
  data() {
    return {
      notificationTypes: {
        status: { color: "primary", icon: "document" },
        add: { color: "success", icon: "plus" },
        warn: { color: "warning", icon: "alarm-clock" },
      },
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.el-color-primary {
  color: #409eff;
}

.el-color-success {
  color: #67c23a;
}

.el-color-warning {
  color: #e6a23c;
}

.el-notification-icon {
  font-size: 24px;
}

.el-row {
  display: flex;
  flex-direction: row;
  align-items: center;
}
</style>
