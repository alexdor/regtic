<template>
  <div class="full-height">
    <Topbar :title="name" />
    <el-main>
      <CheckResult :id="id" />
    </el-main>
  </div>
</template>

<script lang="ts">
  import { Component, Vue } from "vue-property-decorator";
  import CheckResult from "@/components/CheckResult.vue"; // @ is an alias to /src
  import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src

  @Component({
    components: {
      CheckResult,
      Topbar
    }
  })
  export default class Home extends Vue {
    data() {
      return {
        name: "Regtic Demo",
        id: this.$route.params.id
      }
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
