<template>
  <div class="full-height">
    <el-row>
      <el-col :span="14" class="title-col">
        <div class="title vertical-center">
          <div v-if="loading" class="title-loading">
            <VclCode height="40" />
          </div>
          <div v-else>
            {{ name }}
          </div>
        </div>
      </el-col>
      <el-col :span="10" align="right">
        <el-button
          disabled
          title="Coming soon..."
          type="info"
          icon="el-icon-collection-tag"
        >
          Add to watchlist
        </el-button>
        <el-button
          disabled
          title="Coming soon..."
          type="primary"
          icon="el-icon-download"
        >
          Generate report
        </el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-card class="people-card">
          <div slot="header" class="clearfix">
            People
          </div>
          <div v-if="loading">
            <VclList v-for="index in 2" :key="index" />
          </div>
          <div v-else>
            <el-table class="full-width" :data="people">
              <el-table-column prop="name" label="Name" sortable />
              <el-table-column
                prop="type"
                label="Status"
                sortable
                width="100"
                align="center"
              >
                <template slot-scope="scope">
                  <el-tooltip
                    v-if="scope.row.type === 'warning'"
                    class="item"
                    :content="scope.row.source"
                    placement="top"
                    effect="dark"
                  >
                    <el-button
                      circle
                      size="mini"
                      icon="el-icon-view"
                      type="warning"
                    ></el-button>
                  </el-tooltip>
                  <el-tooltip
                    v-else-if="scope.row.type === 'bad'"
                    class="item"
                    :content="scope.row.source"
                    placement="top"
                    effect="dark"
                  >
                    <el-button
                      circle
                      size="mini"
                      icon="el-icon-close"
                      type="danger"
                    ></el-button>
                  </el-tooltip>
                  <el-button
                    v-else
                    circle
                    size="mini"
                    icon="el-icon-check"
                    type="success"
                  ></el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="companies-card">
          <div slot="header" class="clearfix">
            Companies
          </div>
          <div v-if="loading">
            <VclList v-for="index in 2" :key="index" />
          </div>
          <div v-else>
            <el-table class="full-width" :data="companies">
              <el-table-column prop="name" label="Name" sortable />
              <el-table-column prop="vat" label="VAT" sortable width="120" />
              <el-table-column prop="address" label="Address" sortable />
            </el-table>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import api from "@/utils/api";
//@ts-ignore
import { VclCode, VclList } from "vue-content-loading";
//@ts-ignore

@Component({
  components: {
    VclCode,
    VclList
  }
})
export default class Company extends Vue {
  name: string = "";
  vat: string = "";
  address: string = "";
  companies: {
    id: string;
    found: boolean;
    name: string;
    vat: string;
    address: string;
    country_code: string;
    starting_date: number;
  }[] = [];
  people: any[] = [];
  loading: boolean = true;

  async created() {
    const companyId = this.$route.params.id;
    const {
      info: { name, vat, address },
      companies,
      people
    } = await api.validateCompany(companyId);
    this.name = name;
    this.vat = vat;
    this.address = address;
    this.companies = companies;
    this.people = this.parsePeopleArray(people);
    this.loading = false;
  }

  parsePeopleArray(people: any) {
    const goodPeople = people.good
      ? people.good.map((person: any) => {
          return { name: person.full_name, ...person, type: "good" };
        })
      : [];
    const warningPeople = people.warning
      ? people.warning.map((person: any) => {
          return { name: person.full_name, ...person, type: "warning" };
        })
      : [];
    const badPeople = people.bad
      ? people.bad.map((person: any) => {
          return { name: person.full_name, ...person, type: "bad" };
        })
      : [];
    return [...badPeople, ...warningPeople, ...goodPeople];
  }
}
</script>

<style scoped lang="scss">
.title-col {
  height: 40px;
  display: flex;
}

.title {
  font-size: 18px;
  font-weight: 700;
}

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

.people-card {
  margin: 20px 10px 10px 0px;
}

.companies-card {
  margin: 20px 0px 10px 10px;
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

.title-loading {
  height: 30px;
}
</style>
