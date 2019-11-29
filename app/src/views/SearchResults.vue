<template>
  <div class="container-center container-width-small">
    <SearchInput :initial-search-str="initialSearchStr" @search="search" class="bottom-spacing" />
    <div v-if="loading" class="align-center">
      <VclTable class="loading-screen" :rows="15" :columns="10" />
    </div>
    <div v-else>
      <CompanyListItem v-for="item in results" v-bind:key="item.id" :data="item" />
    </div>
  </div>
</template>

<script lang="ts">
import api from "@/utils/mockapi";
import { Component, Vue } from "vue-property-decorator";
import SearchInput from "@/components/SearchInput.vue";
import CompanyListItem from "@/components/CompanyListItem.vue";
import { VclTable } from "vue-content-loading";

@Component({
  components: {
    SearchInput,
    CompanyListItem,
    VclTable
  }
})
export default class SelectCompany extends Vue {
  loading = true;
  initialSearchStr = "";
  results: { name: string; id: string; address: string; vat: string }[] = [];

  created() {
    this.initialSearchStr = this.$route.params.name;
    this.search(this.initialSearchStr);
  }

  async search(searchStr: string) {
    this.loading = true;
    this.results = await api.findCompanies(searchStr);
    this.loading = false;
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  .bottom-spacing {
    margin-bottom: 3rem;
  }

.text-large {
  display: inline-block;
  font-size: 18px;
  margin-bottom: 6px;
}

.el-card p {
  margin: 0;
  font-size: 12px;
  color: #777779;
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

.el-card {
  margin-bottom: 20px;
}

.right-button {
  float: right;
  margin-top: 8px;
  margin-right: 4px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}

.vertical-spacing {
  margin: 20px 0;
}

.loading-screen {
  margin: 40px 0;
}
</style>
