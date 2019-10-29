<template>
  <div class="full-height">
    <Topbar title="Check company" />
    <el-main>
      <div class="col-center">
        <img
          src="../assets/undraw_detailed_analysis.svg"
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
          <!--<el-autocomplete
      class="inline-input full-width"
      v-model="state"
      :fetch-suggestions="querySearchCompany"
      placeholder="Company name or CVR"
      :trigger-on-focus="false"
      @select="handleSelect"
    ></el-autocomplete>-->
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
//  import CheckCompany from "@/components/CheckCompany.vue"; // @ is an alias to /src
import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src

@Component({
  components: {
    //CheckCompany,
    Topbar
  }
})
export default class Home extends Vue {
  data() {
    return {
      //links: [],
      state: ""
    };
  } /*,
  methods: {
    querySearchCompany(queryString, cb) {
      var links = this.links;
      var results = queryString
        ? links.filter(this.createFilter(queryString))
        : links;
      // call callback function to return suggestions
      cb(results);
    },
    createFilter(queryString) {
      return link => {
        return (
          link.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
        );
      };
    },
    loadAll() {
      return [
        { value: "vue", link: "https://github.com/vuejs/vue" },
        { value: "element", link: "https://github.com/ElemeFE/element" },
        { value: "cooking", link: "https://github.com/ElemeFE/cooking" },
        { value: "mint-ui", link: "https://github.com/ElemeFE/mint-ui" },
        { value: "vuex", link: "https://github.com/vuejs/vuex" },
        { value: "vue-router", link: "https://github.com/vuejs/vue-router" },
        { value: "babel", link: "https://github.com/babel/babel" }
      ];
    },
    handleSelect(item) {
      console.log(item);
    }
  },
  mounted() {
    this.links = this.loadAll();
  }*/
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
