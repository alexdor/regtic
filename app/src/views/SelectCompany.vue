<template>
  <div class="full-height">
    <h1>{{ name }}</h1>
    <el-card v-for="company in results" v-bind:key="company.id">
      <span class="text-large">{{ company.name }}</span>
      <router-link :to="'/check/' + company.id">
        <el-button class="right-button" type="primary" icon="el-icon-check"
          >Check</el-button
        >
      </router-link>
      <div class="body">
        <p>{{ company.address }}</p>
        <p>{{ company.vat }}</p>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({})
export default class Home extends Vue {
  data() {
    return {
      name: this.$route.params.name,
      results: [
        {
          name: "Demo Regtic",
          id: "1234-56789-01234-5678",
          address: "Demo Street 1",
          vat: "DK-12345678"
        },
        {
          name: "Regtic Example",
          id: "2345-56789-01234-5678",
          address: "Demo Street 1",
          vat: "DK-23456789"
        },
        {
          name: "Regtic A/S",
          id: "3456-56789-01234-5678",
          address: "Demo Street 2",
          vat: "DK-34567890"
        },
        {
          name: "Regtic Finance",
          id: "4567-56789-01234-5678",
          address: "Test Street 100",
          vat: "DK-45678901"
        },
        {
          name: "Regtic Law",
          id: "5678-56789-01234-5678",
          address: "Test Street 100",
          vat: "DK-56789012"
        },
        {
          name: "Listed Regtic",
          id: "6789-56789-01234-5678",
          address: "Demo Blv. 5",
          vat: "DK-67890123"
        },
        {
          name: "Regtic Test Name",
          id: "7890-56789-01234-5678",
          address: "DTU Lyngby 101 A",
          vat: "DK-78901234"
        }
      ]
    };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
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
</style>
