<template>
  <div class="full-height">
    <Topbar title="Check company" />
    <el-main>
      <div class="col-center">
        <img
          src="../assets/undraw_detailed_analysis.svg"
          alt="Analysis illustration"
          class="fit-to-width"
        />
        <p class="justify-center">
          Check your company relations with a click of a button. Simple, easy.
        </p>
        <div class="vertical-spacing">
          <el-input
            v-model="state"
            class="inline-input full-with"
            placeholder="Company name or CVR"
          ></el-input>
        </div>
        <div class="vertical-spacing">
          <el-button
            type="primary"
            class="full-width"
            :disabled="state.length == 0"
            @click.native="$router.push('/select-company/' + state)"
            >Search</el-button
          >
        </div>
      </div>
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
export default class Home extends Vue {
  data() {
    return {
      state: ""
    };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

.col-center {
  max-width: 420px;
  margin: auto;
  margin-top: 100px;
}

.full-width {
  width: 100%;
}

.justify-center {
  text-align: center;
}

.vertical-spacing {
  margin: 20px 0;
}
</style>
