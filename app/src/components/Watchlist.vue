<template>
  <div>
    <el-row>
      <el-col :span="18"><h1>Watchlist</h1></el-col>
      <el-col :span="6" align="right">
        <el-button type="primary" icon="el-icon-download">Generate report</el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12" v-for="company in items">
        <el-card>
          <span class="text-large">{{company.name}}</span>
          <el-button class="right-button btn-margin-left" type="danger" icon="el-icon-delete" circle></el-button>
          <router-link :to="'/check/' + company.id">
            <el-button class="right-button" type="primary" icon="el-icon-check" round>Check</el-button>
          </router-link>
          <div class="body">
            <p>{{company.address}}</p>
            <p>{{company.vat}}</p>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [
        { vat: "DK-12345678", name: "Bad company example", id: "1234-5678-9012-3456" },
        { vat: "DK-23456789", name: "Other company demo", id: "2345-5678-9012-3456" },
        { vat: "DK-34567890", name: "Example company", id: "3456-5678-9012-3456" },
        { vat: "DK-12345678", name: "Bad company example 2", id: "1234-5678-9012-3456" },
        { vat: "DK-23456789", name: "Other company demo 2", id: "2345-5678-9012-3456" },
        { vat: "DK-34567890", name: "Example company 2", id: "3456-5678-9012-3456" },
        { vat: "DK-12345678", name: "Bad company example 3", id: "1234-5678-9012-3456" },
        { vat: "DK-23456789", name: "Other company demo 3", id: "2345-5678-9012-3456" },
        { vat: "DK-34567890", name: "Example company 3", id: "3456-5678-9012-3456" },
        { vat: "DK-12345678", name: "Bad company example 4", id: "1234-5678-9012-3456" },
        { vat: "DK-23456789", name: "Other company demo 4", id: "2345-5678-9012-3456" },
        { vat: "DK-34567890", name: "Example company 4", id: "3456-5678-9012-3456" }
      ]
    };
  }
};
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
    color: #777779;
  }

  .btn-margin-left {
    margin-left: 15px;
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
  margin: 10px;
}
.right-button {
  float: right;
}

.item {
  margin: 10px 0;
}
.item:first-child {
  margin-top: 0;
}
.item:last-child {
  margin-bottom: 0;
}
</style>
