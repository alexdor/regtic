<template>
  <el-aside>
    <el-menu
      class="el-menu-vertical-demo"
      background-color="#345bcc"
      text-color="rgba(255,255,255,.4)"
      active-text-color="#fff"
      default-active="/"
      :router="true"
    >
      <el-menu-item index="/">
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
</template>

<style lang="scss" scoped>
.el-aside {
  width: 250px !important;
}

.el-menu {
  border-right: 0;
  &-item {
    &:hover,
    &:focus {
      color: #fff !important;
      background-color: #345bcc !important;
    }
    i {
      color: inherit;
    }
  }
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
