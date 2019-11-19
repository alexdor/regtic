<template>
  <el-container>
    <sidebar></sidebar>
    <el-container direction="vertical">
      <topbar :show-go-back="showGoBack" />
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
import Sidebar from "@/components/Sidebar.vue";
import Vue from "vue";

export default Vue.extend({
  components: {
    Topbar,
    Sidebar,
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
</style>
