<template>
  <div id="app">
    <component :is="layout" :show-go-back="showGoBack">
      <router-view />
    </component>
  </div>
</template>

<script>
const defaultLayout = "default";
export default {
  name: "App",
  computed: {
    showGoBack: function() {
      return this.$route.name === "company";
    },
    layout() {
      return (this.$route.meta.layout || defaultLayout) + "-layout";
    }
  }
};
</script>

<style>
#app,
.el-popper,
.el-tooltip__popper {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}

html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
}

#app,
#app .el-container,
#app .el-container .el-aside,
#app .el-container .el-aside .el-menu {
  height: 100%;
}
</style>
