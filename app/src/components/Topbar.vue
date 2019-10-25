<template>
  <el-header>
    <el-row>
      <el-col :span="12">
        <el-page-header @back="goBack" :content="title"></el-page-header>
      </el-col>
      <el-col :span="12" class="align-right">
        <div class="vertical-center">
          <el-badge
            v-if="notifications.length > 0"
            :value="notifications.length"
            :max="99"
            class="item"
          >
            <el-dropdown trigger="click">
              <i class="el-icon-bell el-dropdown-link"></i>
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
          <i class="el-icon-bell" v-else></i>
          <el-divider direction="vertical"></el-divider>
          <el-dropdown trigger="click">
            <el-avatar
              shape="circle"
              :size="32"
              class="el-dropdown-link"
              src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
            ></el-avatar>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item icon="el-icon-s-tools"
                >Settings</el-dropdown-item
              >
              <el-dropdown-item icon="el-icon-s-home">Log out</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </el-col>
    </el-row>
  </el-header>
</template>
<script>
import NotificationItem from "@/components/NotificationItem.vue";

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
  components: { NotificationItem },
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
