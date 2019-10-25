<template>
  <el-header>
    <el-row>
      <el-col :span="12">
        <el-page-header @back="goBack"
                        :content="title"></el-page-header>
      </el-col>
      <el-col :span="12" class="align-right">
        <div class="vertical-center">
          <el-badge v-if="notifications.length > 0" :value="notifications.length" :max="99" class="item">
            <el-dropdown trigger="click">
              <i class="el-icon-bell el-dropdown-link"></i>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item v-for="msg in notifications" class="notification-item">
                  <el-row>
                    <el-col :span="4">
                      <i :class="'el-notification-icon el-icon-' + notification_types[msg.type].icon + ' el-color-' + notification_types[msg.type].color"></i>
                    </el-col>
                    <el-col :span="20">
                      <p v-html="msg.text"></p>
                    </el-col>
                  </el-row>
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </el-badge>
          <i class="el-icon-bell" v-else></i>
          <el-divider direction="vertical"></el-divider>
          <el-dropdown trigger="click">
            <el-avatar shape="circle"
                       :size="32"
                       class="el-dropdown-link"
                       src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"></el-avatar>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item icon="el-icon-s-tools">Settings</el-dropdown-item>
              <el-dropdown-item icon="el-icon-house">Log out</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </el-col>
    </el-row>
  </el-header>
</template>
<script>
  export default {
    data() {
      return {
        notification_types: {
          status: { color: "primary", icon: "document" },
          add: { color: "success", icon: "plus" },
          warn: { color: "warning", icon: "alarm-clock" }
        },
        notifications: [
          { type: "add", text: "<b>Company</b> was added to your watchlist", date: new Date() },
          { type: "status", text: "<b>A new monthly report is ready to download</b>", date: new Date() },
          { type: "warn", text: "<b>Donald Trump</b> has been added to the PEP list", date: new Date(2017, 1, 1) }
        ]
      }
    },
    props: ["title"],
    created() {
      document.title = this.title;
    },
    methods: {
      goBack() {
        history.back();
      }
    }
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

  .el-header {
    background-color: #fff;
  }

  .el-color-primary {
    color: #409EFF;
  }

  .el-color-success {
    color: #67C23A;
  }

  .el-color-warning {
    color: #E6A23C;
  }

  .notification-item {
    line-height: 1.5;
  }

  .el-notification-icon {
    font-size: 24px;
  }
</style>
