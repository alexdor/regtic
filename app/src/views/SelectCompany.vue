<template>
  <div class="full-height">
    <Topbar title="Select company" />
    <el-main>
      <SelectCompany :name="name" />
    </el-main>
  </div>
</template>

<script lang="ts">
  import { Component, Vue } from "vue-property-decorator";
  import SelectCompany from "@/components/SelectCompany.vue"; // @ is an alias to /src
  import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src

  @Component({
    components: {
      SelectCompany,
      Topbar
    }
  })
  export default class Home extends Vue {
    data() {
      return {
        name: this.$route.params.name
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
