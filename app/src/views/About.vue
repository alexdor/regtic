<template>
  <div class="full-height">
    <Topbar title="About" />
    <el-main>
      <h1>About us</h1>
      <p>Contents of this page yet to be agreed upon.</p>
    </el-main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src

@Component({
  components: {
    Topbar
  }
})
export default class Home extends Vue { }
</script>
