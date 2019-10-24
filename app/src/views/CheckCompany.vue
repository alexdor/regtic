<template>
  <div class="full-height">
    <Topbar title="Check company" />
    <el-main>
      <CheckCompany />
    </el-main>
  </div>
</template>

<script lang="ts">
  import { Component, Vue } from "vue-property-decorator";
  import CheckCompany from "@/components/CheckCompany.vue"; // @ is an alias to /src
  import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src

  @Component({
    components: {
      CheckCompany,
      Topbar
    }
  })
  export default class Home extends Vue { }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.full-height {
  height: 100%;
  background-color: #fafafa;
}
</style>
