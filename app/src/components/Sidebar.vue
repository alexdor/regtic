<template>
  <el-aside>
    <el-menu
      class="el-menu-vertical-demo"
      text-color="rgba(255,255,255,.4)"
      active-text-color="#fff"
      default-active="/"
      :router="true"
    >
      <img src="../assets/regtic.png" class="logo" alt="Regtic logo" />
      <div class="user-info">
        <span class="user-name text-bold">Jeff Hendriks</span>
        <span class="user-email text-gray">example@example.com</span>
      </div>
      <div class="menu-items">
        <el-menu-item index="/search" title="Search">
          <i class="el-icon-search" aria-hidden="true" /><span>Search</span>
        </el-menu-item>
        <el-menu-item index="/watchlist" disabled title="Watchlist - coming soon...">
          <i class="el-icon-tickets" aria-hidden="true"></i><span>Watchlist</span>
        </el-menu-item>
        <el-menu-item index="not-added-yet" disabled title="Subscription - coming soon...">
          <i class="el-icon-s-finance" aria-hidden="true"></i><span>Subscription</span>
        </el-menu-item>
      </div>
    </el-menu>
  </el-aside>
</template>

<style lang="scss" scoped>
  $sidebar-bg: #F5F6F8;
  $black: #303133;
  $gray: #90A0B7;

  @import url('https://fonts.googleapis.com/css?family=Poppins&display=swap');

  .el-aside {
    width: 256px !important;
  }

  .el-menu {
    background-color: $sidebar-bg;
    box-shadow: inset -2px 0 15px rgba(0, 0, 0, 0.12);
    font-family: 'Poppins', sans-serif;
  }

  .logo {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 88px;
    margin-top: 96px;
  }

  .user-info {
    margin-left: 48px;
    margin-top: 30px;
  }

  .text-bold {
    display: block;
    color: $black;
    font-size: 0.875rem;
    font-weight: 500;
    line-height: 1.3125rem;
  }

  .text-gray {
    display: block;
    color: $gray;
    font-size: 0.6875rem;
    font-weight: 500;
    line-height: 1rem;
    letter-spacing: 0.01em;
  }

  /* Default content styling, for small screens (phones) */
  @media screen and (max-width: 800px) {
    .el-aside {
      width: 80px !important; /* Override other CSS width of 250px, as well as elem.io's in-line styling width. */
    }

    .el-aside .el-menu-item {
      text-align: center;
    }

      .el-aside .el-menu-item i:before {
        font-size: 1.5rem;
      }

    .el-aside .el-menu-item span {
      display: none;
    }

    .logo {
      padding-left: 5px !important; /* elem.io has in-line styling for padding-left specifically. */
      padding-right: 5px;
    }
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

.menu-divider {
  display: block;
  height: 1.5px;
  width: 85%;
  margin: 1.25rem auto;
  background-color: rgba(255, 255, 255, 0.3);
  position: relative;
}
</style>
