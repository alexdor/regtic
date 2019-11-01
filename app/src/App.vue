<template>
  <div id="app">
    <el-container>
      <el-aside>
        <el-menu
          class="el-menu-vertical-demo"
          background-color="#345bcc"
          text-color="rgba(255,255,255,.4)"
          active-text-color="#fff"
          :default-active="$route.path"
          :router="true"
        >
          <img class="logo" src="./assets/regtic-logo.png" alt="Regtic" />
          <div class="menu-divider" />
          <el-menu-item divided index="/">
            <i class="el-icon-search" aria-hidden="true" />Search
          </el-menu-item>
          <el-menu-item index="/watchlist">
            <i class="el-icon-tickets" aria-hidden="true"></i>Watchlist
          </el-menu-item>
          <el-menu-item index="not-added-yet">
            <i class="el-icon-s-finance" aria-hidden="true"></i>Subscription
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container direction="vertical">
        <Topbar title="Watchlist" />
        <el-main>
          <router-view />
          <Footer />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src
import Footer from "@/components/Footer.vue"; // @ is an alias to /src

export default {
  name: "app",
  components: {
    Footer,
    Topbar
  }
};
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
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
.fit-to-width {
  width: calc(100% - 60px);
  padding: 0 30px;
}
.el-aside {
  width: 250px !important;
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
.el-menu-item:hover,
.el-menu-item:focus {
  color: #fff !important;
  background-color: #345bcc !important;
}
.el-menu-item i {
  color: inherit !important;
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
.el-container {
  overflow-y: auto;
}
.full-height {
  background-color: #fafafa;
  display: inline-table;
  min-height: calc(100% - 60px);
  /* FIXME: this is a dirty fix and should be fixed properly in the feature */
  width: 100%;
}
.logo {
  padding-top: 20px;
  width: 150px;
  margin-left: auto;
  margin-right: auto;
  display: block;
}
.menu-divider {
  display: block;
  height: 1.5px;
  width: 85%;
  margin-left: auto;
  margin-right: auto;
  margin-top: 20px;
  margin-bottom: 20px;
  background-color: rgba(255, 255, 255, 0.3);
  position: relative;
}
</style>
