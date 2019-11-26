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
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}

.el-notification__group {
  width: 230px;
}

.el-notification__title {
  overflow: hidden;
  text-overflow: ellipsis;
}

.searchInput .el-input__inner {
  border-width: 2px;
}

#app,
#app .el-container,
#app .el-container .el-aside,
#app .el-container .el-aside .el-menu {
  height: 100%;
}
</style>
