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
        <router-link to="/search" title="Search" class="menu-item">
          <i class="el-icon-search" aria-hidden="true" /><span>Search</span>
        </router-link>
        <router-link to="/watchlist" title="Watchlist" class="menu-item">
          <i class="el-icon-document-copy" aria-hidden="true"></i><span>Watchlist</span>
        </router-link>
        <router-link to="not-added-yet" disabled title="Subscription - coming soon..." class="menu-item">
          <i class="el-icon-chat-square" aria-hidden="true"></i><span>Subscription</span>
        </router-link>
        <router-link to="not-added-yet" disabled title="Settings - coming soon..." class="menu-item">
          <i class="el-icon-setting" aria-hidden="true"></i><span>Settings</span>
        </router-link>
      </div>
    </el-menu>
  </el-aside>
</template>

<style lang="scss" scoped>
  $sidebar-bg: #F5F6F8;
  $black: #303133;
  $gray: #90A0B7;

  $link-normal: #C2CFE0;
  $link-hover: #109CF1;

  .el-aside {
    width: 16rem !important; /* elem.io using inline styling! */
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
    width: 5.5rem;
    margin-top: 8rem;
  }

  .user-info {
    margin-left: 4rem;
    margin-top: 1.875rem;
    margin-bottom: 1.875rem;
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

  .menu-item {
    color: $link-normal;
    font-size: 0.8125rem;
    line-height: 1.1875rem;
    padding: 0.5rem 0;
    margin: 0.5rem 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    text-decoration: none;
    font-weight: 500;
    padding-left: 4rem;
    &[disabled] {
      cursor: not-allowed;
      opacity: 0.75;
      color: $link-normal !important;
    }
    &:hover,
    &:focus {
      color: $link-hover;
    }
    i {
      font-size: 1.25rem;
    }
    span {
      margin-left: 0.875rem;
    }
  }

  .menu-item.router-link-active {
    color: $link-hover;
  }

  /* Default content styling, for small screens (phones) */
  @media screen and (max-width: 800px) {
    .el-aside {
      width: 80px !important; /* Override other CSS width of 250px, as well as elem.io's in-line styling width. */
    }

    .logo {
      margin-top: 2rem;
      margin-bottom: 2rem;
      width: 60%;
    }

    .user-info {
      display: none;
    }

    .menu-item {
      padding: 1rem;
      margin: 1rem auto;
      display: flex;
      flex-direction: column;
      align-items: center;
    }

    .menu-item span {
      display: none;
    }
  }
</style>
