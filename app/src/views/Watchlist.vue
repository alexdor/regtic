<template>
  <div class="full-height">
    <Topbar title="Watchlist" />
    <el-main>
      <Watchlist />
    </el-main>
  </div>
</template>

<script lang="ts">
  import { Component, Vue } from "vue-property-decorator";
  import Watchlist from "@/components/Watchlist.vue"; // @ is an alias to /src
  import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src

  @Component({
    components: {
      Watchlist,
      Topbar
    }
  })
  export default class Home extends Vue {
    data() {
      return {}
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.full-height {
  height: 100%;
  background-color: #fafafa;
}
</style>
