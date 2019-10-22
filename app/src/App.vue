<template>
  <div id="app">
    <el-container>
      <el-aside>
        <el-menu
          class="el-menu-vertical-demo"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#409eff"
          default-active="1"
        >
          <img src="./assets/regtic-logo.svg" class="fit-to-width" />
          <el-menu-item index="1"
            ><i class="el-icon-menu"></i>Dashboard</el-menu-item
          >
          <el-menu-item index="2"
            ><i class="el-icon-star-on"></i>Watchlist</el-menu-item
          >
          <el-menu-item index="3"
            ><i class="el-icon-s-finance"></i>Subscription</el-menu-item
          >
        </el-menu>
      </el-aside>
      <el-container>
        <el-header>
          <el-row>
            <el-col :span="12">
              <el-page-header
                @back="goBack"
                content="Check company"
              ></el-page-header>
            </el-col>
            <el-col :span="12" class="align-right">
              <div class="vertical-center">
                <i class="el-icon-bell"></i>
                <el-divider direction="vertical"></el-divider>
                <el-avatar
                  shape="circle"
                  :size="32"
                  src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
                ></el-avatar>
              </div>
            </el-col>
          </el-row>
        </el-header>
        <el-main>
          <router-view></router-view>
        </el-main>
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
  components: {},
  methods: {
    goBack() {
      history.back();
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
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
