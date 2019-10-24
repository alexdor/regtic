<template>
  <div>
    <el-row>
      <el-col :span="18"><h1>Overview</h1></el-col>
      <el-col :span="6" align="right">
        <el-button type="info" icon="el-icon-collection-tag">Add to watchlist</el-button>
        <el-button type="primary" icon="el-icon-download">Generate report</el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="6">
        <el-card>
          <span class="text-large">Daughter companies</span>
          <p>{{result.daughter_companies_total}}</p>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card>
          <span class="text-large">People checked</span>
          <p>{{result.people_total}}</p>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card>
          <span class="text-large">Companies found on your watchlist</span>
          <p>{{result.companies_on_watchlist_percent * 100}}%</p>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card>
          <span class="text-large">Politically exposed people detected</span>
          <p>{{result.pep_detect}}</p>
        </el-card>
      </el-col>
    </el-row>
    <el-card>
      <div slot="header" class="clearfix">
        <span>Detected companies</span>
      </div>
      <div v-for="company in result.detected_companies" class="text item">
        <el-tooltip class="item" :content="company.source" placement="top" effect="dark"><el-button circle size="mini" icon="el-icon-link"></el-button></el-tooltip> {{company.vat}} - {{company.name}}
      </div>
    </el-card>
    <el-card>
      <div slot="header" class="clearfix">
        <span>Detected people</span>
      </div>
      <div v-for="person in result.detected_people" class="text item">
        <el-badge v-if="person.pep" value="PEP" class="item mini" type="warning"><el-tooltip class="item" :content="person.source" placement="top" effect="dark">
          <el-button circle size="mini" icon="el-icon-link"></el-button></el-tooltip> {{person.full_name}}
        </el-badge>
        <span v-else>
          <el-tooltip class="item" :content="person.source" placement="top" effect="dark"><el-button circle size="mini" icon="el-icon-link"></el-button></el-tooltip> {{person.full_name}}
        </span>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: "Regtic Demo",
      id: this.$route.params.id,
      result: {
        daughter_companies_total: 5,
        people_total: 2500,
        companies_on_watchlist_percent: 0.2,
        pep_detect: 1,
        detected_companies: [
          { vat: "DK-12345678", name: "Bad company example", source: "EU Sanctions List 2019 (www.example.org/path-to-file) - 23/10/2019" },
          { vat: "DK-23456789", name: "Other company demo", source: "EU Sanctions List 2019 (www.example.org/path-to-file) - 23/10/2019" },
          { vat: "DK-34567890", name: "Example company", source: "FATF List 2019 (www.example.org/path-to-file2) - 18/10/2019" }
        ],
        detected_people: [
          { full_name: "John Doe", pep: false, source: "EU Sanctions List 2019 (www.example.org/path-to-file) - 23/10/2019" },
          { full_name: "Jane Doe", pep: false, source: "FN Sanctions List 2019 (www.other.org/path-to-file) - 21/10/2019" },
          { full_name: "James Doe", pep: false, source: "EU Sanctions List 2019 (www.example.org/path-to-file) - 23/10/2019" },
          { full_name: "Donald Trump", pep: true, source: "EU PEP List 2019 (www.example.org/path-to-file) - 22/10/2019" },
        ]
      }
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
