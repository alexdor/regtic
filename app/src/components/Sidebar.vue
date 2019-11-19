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
      <el-menu-item index="/" class="logo">
        <img class="full" src="../assets/regtic-logo.png" alt="Regtic logo" />
        <img class="small" src="../assets/regtic-logo-small.png" alt="Regtic logo" />
      </el-menu-item>
      <div class="menu-divider" />
      <el-menu-item index="/search" title="Search">
        <i class="el-icon-search" aria-hidden="true" /><span>Search</span>
      </el-menu-item>
      <el-menu-item index="/watchlist" disabled title="Watchlist - coming soon...">
        <i class="el-icon-tickets" aria-hidden="true"></i><span>Watchlist</span>
      </el-menu-item>
      <el-menu-item index="not-added-yet" disabled title="Subscription - coming soon...">
        <i class="el-icon-s-finance" aria-hidden="true"></i><span>Subscription</span>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<style lang="scss" scoped>
.el-aside {
  width: 250px !important;
}

  .logo img {
    margin-left: auto;
    margin-right: auto;
  }

  .logo .full {
    padding-top: 0.75rem;
    width: 60%;
    display: block;
  }

.logo .small {
  padding-top: 0.5rem;
  display: none;
  width: 90%;
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

    .logo .small {
      display: block;
    }

    .logo .full {
      display: none;
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
