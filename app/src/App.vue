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
      return ["company", "search-name"].indexOf(this.$route.name) > -1;
    },
    layout() {
      return (this.$route.meta.layout || defaultLayout) + "-layout";
    }
  }
};
</script>

<style lang="scss">
$footer-height: 76px;

@import url("https://fonts.googleapis.com/css?family=Poppins&display=swap");

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

/*Override elem.io color on buttons*/
.el-button--primary {
  background-color: #1989fa !important;
  border-color: #1989fa !important;
}

.el-button--primary:hover,
.el-button--primary:focus {
  background-color: #167be0 !important;
  border-color: #167be0 !important;
}

.el-button--primary[disabled] {
  background-color: #4fa4f9 !important;
  border-color: #4fa4f9 !important;
}

#app,
#app .el-container,
#app .el-container .el-aside,
#app .el-container .el-aside .el-menu {
  height: 100%;
}

.container-center {
  position: relative;
  left: 50%;
  transform: translate(-50%, 0);
  min-height: calc(100% - #{$footer-height});
}

.container-width-small {
  width: 524px;
}

.container-width-medium {
  width: 924px;
}
</style>
