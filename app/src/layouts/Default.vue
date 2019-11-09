<template>
  <el-container>
    <el-aside>
      <el-menu
        class="el-menu-vertical-demo"
        background-color="#345bcc"
        text-color="rgba(255,255,255,.4)"
        active-text-color="#fff"
        default-active="/"
        :router="true"
      >
        <el-menu-item class="logo-wrapper" index="/">
          <img class="logo" src="../assets/regtic-logo.png" alt="Regtic logo" />
        </el-menu-item>
        <div class="menu-divider" />
        <el-menu-item index="/search">
          <i class="el-icon-search" aria-hidden="true" />Search
        </el-menu-item>
        <el-menu-item index="/watchlist" disabled title="coming soon...">
          <i class="el-icon-tickets" aria-hidden="true"></i>Watchlist
        </el-menu-item>
        <el-menu-item index="not-added-yet" disabled title="coming soon...">
          <i class="el-icon-s-finance" aria-hidden="true"></i>Subscription
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-container direction="vertical">
      <Topbar :show-go-back="showGoBack" />
      <el-main>
        <router-view />
        <Footer />
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts">
import Topbar from "@/components/Topbar.vue";
import Footer from "@/components/Footer.vue";
import Vue from "vue";

export default Vue.extend({
  components: {
    Topbar,
    Footer
  },
  props: {
    showGoBack: { type: Boolean, default: true }
  }
});
</script>

<style lang="scss">
$footer-height: 60px; /* Height of the footer element */

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

.fit-to-width {
  width: calc(
    100% - 4rem
  ); /* Adding spacing to the left/right of the given element, in case it'll otherwise hit the left/right edges of its outer container. Makes it so the total width of the element is still 100%. */
  padding: 0 2rem;
}

.el-aside {
  width: 250px !important;
}

.el-main {
  background-color: #fafafa;
  color: #858796;
}

.el-header .el-row {
  width: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
}

.el-menu {
  border-right: 0;
}

.el-menu-item:hover,
.el-menu-item:focus {
  color: #fff !important;
  background-color: #345bcc !important;
}

.el-menu-item i {
  color: inherit;
}

.align-center {
  text-align: center;
}

.align-right {
  text-align: right;
}

.vertical-center {
  display: inline-flex;
  flex-direction: row;
  align-items: center;
}

.el-autocomplete-suggestion li {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}

.el-footer {
  line-height: $footer-height;
}

.el-container {
  overflow-y: auto;
}

.full-height {
  background-color: #fafafa;
  display: inline-table;
  min-height: calc(100% - #{$footer-height});
  /* FIXME: this is a dirty fix and should be fixed properly in the feature */
  width: 100%;
}

.el-menu-item .logo-wrapper {
  height: none;
}

.logo {
  padding-top: 1.25rem;
  width: 60%;
  margin-left: auto;
  margin-right: auto;
  display: block;
}

.menu-divider {
  display: block;
  height: 1.5px;
  width: 85%;
  margin: 1.25rem auto;
  background-color: rgba(255, 255, 255, 0.3);
  position: relative;
}
</style>
