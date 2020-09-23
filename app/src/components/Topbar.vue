<template>
  <el-header>
    <el-row>
      <el-col :span="24">
        <i v-if="showGoBack" class="el-icon-back" @click="goBack" />
      </el-col>
    </el-row>
  </el-header>
</template>
<script>
// import NotificationItem from "@/components/NotificationItem.vue";

export default {
  // components: { NotificationItem },
  name: "TopBar",
  props: ["showGoBack"],
  methods: {
    goBack() {
      history.back();
    },
  },
};
</script>

<style scoped lang="scss">
.el-header {
  top: 0;
  height: 3.75rem;
  z-index: 1;
  display: flex;
  flex-direction: row;
  align-items: center;
}

.el-header i {
  color: #606266;
  font-size: 1.5em;
  vertical-align: middle;
  cursor: pointer;
}
</style>
