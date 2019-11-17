<template>
  <div class="full-height">
    <el-card>
      <h1>{{ result.name }}</h1>
      <el-button class="button-right" type="primary" icon="el-icon-plus">Add to watchlist</el-button>
      <div class="wrapper">
        <table cellspacing="20">
          <tbody>
            <tr>
              <td class="info-name">Location</td>
              <td class="info-value">{{ result.location }}<br>{{ result.country }}</td>
            </tr>
            <tr>
              <td class="info-name">VAT</td>
              <td class="info-value">{{ result.vat }}</td>
            </tr>
            <tr>
              <td class="info-name">Business</td>
              <td class="info-value">{{ result.business }}</td>
            </tr>
            <tr>
              <td class="info-name">Entities</td>
              <td class="info-value">{{ result.beneficialOwners.length }} Beneficial owners, {{ result.relatedCompaniesCount }} Related companies</td>
            </tr>
          </tbody>
        </table>
        <div class="bottom-right">
          <div class="statusPercent">
            <span class="text">28 Good, 4 Warnings, 0 Critical</span>
            <div class="line">
              <div class="bg-good" style="width:80%"></div>
              <div class="bg-warning" style="width:20%"></div>
              <div class="bg-critical" style="width:0%"></div>
            </div>
          </div>
          <i class="status-icon el-icon-warning"></i>
        </div>
      </div>
    </el-card>
    <el-card>
      <div slot="header" class="v-center">
        <div class="v-center space-content">
          <h2>Beneficial owners</h2>
          <div class="statusPercent">
            <span class="text">28 Good, 4 Warnings, 0 Critical</span>
            <div class="line">
              <div class="bg-good" style="width:80%"></div>
              <div class="bg-warning" style="width:20%"></div>
              <div class="bg-critical" style="width:0%"></div>
            </div>
          </div>
          <el-checkbox v-model="filterCritical">Sanctioned people</el-checkbox>
          <el-checkbox v-model="filterWarning">Political people</el-checkbox>
        </div>
        <el-button class="right-button btn-large" type="text">
          <i class="el-icon-plus" v-if="!beneficialOwnersOpen" v-on:click="beneficialOwnersOpen = true"></i>
          <i class="el-icon-minus" v-if="beneficialOwnersOpen" v-on:click="beneficialOwnersOpen = false"></i>
        </el-button>
      </div>
      <div class="body" v-if="beneficialOwnersOpen">
        <table class="simple-table">
          <tbody>
            <tr v-for="person in applyFilter(result.beneficialOwners, filterWarning, filterCritical)" v-bind:key="person.full_name">
              <td class="table-text-large">{{ person.fullName }}</td>
              <td>
                <div class="flex-row">
                  <span class="table-text-bold">Functions</span>
                  <div>
                    <div v-for="func in person.functions" v-bind:key="func">{{ func }}</div>
                  </div>
                </div>
              </td>
              <td>
                <div class="flex-row">
                  <div>
                    <div class="table-text-bold">Ownership</div>
                    <div class="table-text-bold">Voting rights</div>
                  </div>
                  <div>
                    <div>{{ person.ownership * 100 }}%</div>
                    <div>{{ person.votingRights * 100 }}%</div>
                  </div>
                </div>
              </td>
              <td align="right">
                <i class="status-icon" v-bind:class="status[person.status.type]" v-bind:title="person.status.source"></i>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </el-card>
    <!--<el-row>
      <el-col v-for="company in items" :key="company.id" :span="12">
        <el-card>
          <span class="text-large">{{ company.name }}</span>
          <el-popover v-model="company.deleteVisible"
                      placement="top"
                      width="200"
                      trigger="click">
            <p>Are you sure to delete this?</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini"
                         type="text"
                         @click="company.deleteVisible = false">cancel</el-button>
              <el-button type="danger"
                         size="mini"
                         @click="company.deleteVisible = false">confirm</el-button>
            </div>
            <el-button slot="reference"
                       class="right-button btn-margin-left"
                       type="danger"
                       icon="el-icon-delete"
                       circle></el-button>
          </el-popover>
          <router-link :to="'/check/' + company.id">
            <el-button class="right-button"
                       type="primary"
                       icon="el-icon-check"
                       round>Check</el-button>
          </router-link>
          <div class="body">
            <p>{{ company.address }}</p>
            <p>{{ company.vat }}</p>
          </div>
        </el-card>
      </el-col>
    </el-row>-->
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
    beneficialOwnersOpen = false;
    filterWarning = false;
    filterCritical = false;

    data() {
      return {
        status: {
          good: "el-icon-success",
          warning: "el-icon-warning",
          critical: "el-icon-error"
        },
        result: {
          name: "Company name goes here",
          location: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
          country: "Denmark",
          vat: "DK-12345678",
          business: "University",
          relatedCompaniesCount: 3,
          beneficialOwners: [
            {
              fullName: "John Doe",
              functions: ["UK Prime Minister", "Secondary job"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "warning",
                source: "EU PEP List 2019, 16-11-2019"
              }
            },
            {
              fullName: "John Doe",
              functions: ["UK Prime Minister", "Secondary job"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "warning",
                source: "EU PEP List 2019, 16-11-2019"
              }
            },
            {
              fullName: "John Doe",
              functions: ["UK Prime Minister", "Secondary job"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "warning",
                source: "EU PEP List 2019, 16-11-2019"
              }
            },
            {
              fullName: "John Doe",
              functions: ["UK Prime Minister", "Secondary job"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "warning",
                source: "EU PEP List 2019, 16-11-2019"
              }
            },
            {
              fullName: "John Doe",
              functions: ["UK Prime Minister", "Secondary job"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "warning",
                source: "EU PEP List 2019, 16-11-2019"
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            },
            {
              fullName: "John Doe",
              functions: ["CTO"],
              ownership: 0.03,
              votingRights: 0.12,
              status: {
                type: "good",
                source: null
              }
            }
          ]
        }
      };
    }
    toggleBeneficialOwners() {
      this.beneficialOwnersOpen = !this.beneficialOwnersOpen;
    }
    applyFilter(beneficialOwners: { fullName: string, functions: string[], ownership: number, votingRights: number, status: {type: string, source: string}}[], filterWarning : boolean, filterCritical: boolean) {
      return beneficialOwners.filter(function (person) {
        if (!filterWarning && !filterCritical) return true;
        if (filterWarning && person.status.type === 'warning') return true;
        if (filterCritical && person.status.type === 'critical') return true;
        return false;
      });
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  h1, h2 {
    display: inline-block;
    margin-top: 0;
    color: #717171;
  }

  h1 {
    font-size: 1.75rem;
    margin-bottom: 0.375rem;
  }

  h2 {
    font-size: 1.375rem;
    margin-bottom: 0;
  }

  .wrapper {
    position: relative;
    width: 100%;
  }

  .bottom-right {
    position: absolute;
    right: 0;
    bottom: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  .info-name {
    font-weight: bold;
    font-size: 1rem;
    line-height: 1.125rem;
    text-align: right;
    padding-right: 0.5rem;
  }

  .el-card__header .el-checkbox {
    zoom: 1.375;
  }

  .button-right {
    float: right;
  }

  .btn-large {
    font-size: 1.5rem;
    padding: 0.25rem;
  }

  .v-center {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }

  .simple-table {
    width: 100%;
    border-collapse: collapse;
  }

  .table-text-large {
    font-size: 1.125rem;
    font-weight: 400;
  }

  .simple-table tr:not(:last-child) td {
    border-bottom: 1px solid #EEE;
  }

  .simple-table td {
    padding-top: 0.75rem;
    padding-bottom: 0.75rem;
  }

  .flex-row {
    display: flex;
    flex-direction: row;
    align-items: center;
    font-size: 0.925rem;
  }

  .flex-row div {
    margin: 0.25rem;
  }

  .flex-row .table-text-bold {
    font-weight: bold;
    margin-right: 0.75rem;
  }

  .status-icon {
    font-size: 2.5rem;
  }

  .status-icon.el-icon-success {
    color: #67C23A;
  }

  .status-icon.el-icon-warning {
    color: #E6A23C;
  }

  .status-icon.el-icon-error {
    color: #F56C6C;
  }

  .bg-good {
    background-color: #67C23A;
  }

  .bg-warning {
    background-color: #E6A23C;
  }

  .bg-critical {
    background-color: #F56C6C;
  }

  .space-content > * {
    margin-right: 3rem;
  }

  .statusPercent {
    display: flex;
    flex-direction: column;
    margin-right: 2rem;
  }

  .statusPercent .line {
    width: 10vw;
    height: 4px;
    display: flex;
    flex-direction: row;
    border-radius: 2px;
    overflow: hidden;
    margin-top: 0.25rem;
  }

  .statusPercent .line div {
    height: 100%;
  }

  .el-card p {
    margin: 0;
    color: #777779;
  }

  .btn-margin-left {
    margin-left: 15px;
  }

  .el-popover p {
    margin-top: 0;
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
