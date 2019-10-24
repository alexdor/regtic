<template>
  <div id="app">
    <el-container>
      <el-aside>
        <el-menu
          class="el-menu-vertical-demo"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#409eff"
          :default-active="$route.path"
          :router="true"
        >
          <img src="./assets/regtic-logo.svg" class="fit-to-width" />
          <el-menu-item index="/"
            ><i class="el-icon-menu"></i>Dashboard</el-menu-item
          >
          <el-menu-item index="/watchlist"
            ><i class="el-icon-star-on"></i>Watchlist</el-menu-item
          >
          <el-menu-item index="not-added-yet"
            ><i class="el-icon-s-finance"></i>Subscription</el-menu-item
          >
        </el-menu>
      </el-aside>
      <el-container>
        <router-view></router-view>
        <el-footer class="align-center">
          Copyright © Regtic 2019
        </el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "app",
  components: {}
};
</script>

<style>
  #app {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
  }
  #app, .el-tooltip__popper {
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
.fit-to-width {
  width: calc(100% - 60px);
  padding: 0 30px;
}
.el-main {
  background-color: #fafafa;
  box-shadow: inset 0 2px 6px -5px;
  color: #858796;
}
.el-header {
  display: flex;
  flex-direction: row;
  align-items: center;
}
.el-header .el-row {
  width: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
}
.align-center {
  text-align: center;
}
.align-right {
  text-align: right;
}
.el-header i {
  color: #606266;
  font-size: 1.5em;
  vertical-align: middle;
}
.vertical-center {
  display: inline-flex;
  flex-direction: row;
  align-items: center;
}
.el-header .el-divider--vertical {
  margin: 0 16px;
}
.el-autocomplete-suggestion li {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}
.el-footer {
  line-height: 60px;
}
</style>
