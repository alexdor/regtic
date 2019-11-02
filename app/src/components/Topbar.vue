<template>
  <el-header>
    <el-row>
      <el-col :span="12">
        <i v-if="includeGoBack" class="el-icon-back" @click="goBack" />
      </el-col>
      <el-col :span="12" class="align-right">
        <div class="vertical-center">
          <!--
          <el-badge
            v-if="notifications.length > 0"
            :value="notifications.length"
            :max="99"
            class="item"
          >
            <el-dropdown trigger="click">
              <i class="el-icon-bell el-dropdown-link" aria-hidden="true"></i>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item
                  v-for="msg in notifications"
                  v-bind:key="msg.id"
                  class="notification-item"
                >
                  <router-link v-if="msg.link != null" :to="msg.link">
                    <NotificationItem :data="msg" />
                  </router-link>
                  <NotificationItem v-else :data="msg" />
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </el-badge>
          <i class="el-icon-bell" v-else aria-hidden="true"></i>
          <el-divider direction="vertical"></el-divider>
          -->
          <el-dropdown trigger="click">
            <el-avatar
              shape="circle"
              :size="28"
              class="el-dropdown-link"
              src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
            />
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item
                icon="el-icon-s-tools"
                aria-hidden="true"
                disabled
              >
                Settings
              </el-dropdown-item>
              <el-dropdown-item icon="el-icon-s-home" aria-hidden="true">
                Log out
              </el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </el-col>
    </el-row>
  </el-header>
</template>
<script>
// import NotificationItem from "@/components/NotificationItem.vue";

export default {
  data() {
    return {
      notification_types: {
        status: { color: "primary", icon: "document" },
        add: { color: "success", icon: "plus" },
        warn: { color: "warning", icon: "alarm-clock" }
      },
      notifications: [
        {
          id: 1,
          type: "add",
          boldText: "Company",
          text: " was added to your watchlist",
          date: new Date(),
          link: "/watchlist"
        },
        {
          id: 2,
          type: "status",
          boldText: "A new monthly report is ready to download",
          text: "",
          date: new Date(),
          link: null
        },
        {
          id: 3,
          type: "warn",
          boldText: "Donald Trump",
          text: " has been added to the PEP list",
          date: new Date(2017, 1, 1),
          link: null
        }
      ]
    };
  },
  // components: { NotificationItem },
  props: ["title", "includeGoBack"],
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

<style scoped lang="scss">
.el-header {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  top: 0;
  height: 60px;
  z-index: 1;
  display: flex;
  flex-direction: row;
  align-items: center;
}

.el-header .el-divider--vertical {
  margin: 0 16px;
}

.el-header i {
  color: #606266;
  font-size: 1.5em;
  vertical-align: middle;
  cursor: pointer;
}

.el-color-primary {
  color: #409eff;
}

.el-color-success {
  color: #67c23a;
}

.el-color-warning {
  color: #e6a23c;
}

.notification-item {
  line-height: 1.5;
}

.el-notification-icon {
  font-size: 24px;
}

.notification-item .el-row {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.el-dropdown-link {
  cursor: pointer;
}

.notification-item a {
  text-decoration: none;
  color: #606266;
}

.notification-item a:hover {
  background-color: #ecf5ff;
  color: #66b1ff;
}
</style>
