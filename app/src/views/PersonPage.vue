<template>
  <div class="full-height">
    <el-card>
      <h1>{{ result.fullName }}</h1>
      <div class="wrapper">
        <table cellspacing="20">
          <tbody>
            <tr>
              <td class="info-name">Country</td>
              <td class="info-value">{{ result.country }}</td>
            </tr>
            <tr>
              <td class="info-name">Updated</td>
              <td class="info-value">{{ result.updated }}</td>
            </tr>
            <tr>
              <td class="info-name">Functions</td>
              <td class="info-value"><div v-for="func in result.functions" v-bind:key="func">{{ func }}</div></td>
            </tr>
          </tbody>
        </table>
        <div class="bottom-right">
          <div class="statusPercent">
            <span class="text">2 Good, 3 Warnings, 0 Critical</span>
            <div class="line">
              <div class="bg-good" style="width:40%"></div>
              <div class="bg-warning" style="width:60%"></div>
              <div class="bg-critical" style="width:0%"></div>
            </div>
          </div>
          <i class="status-icon el-icon-warning"></i>
        </div>
      </div>
    </el-card>
    <el-row>
      <el-col :span="12">
        <el-card>
          <div slot="header" class="v-center">
            <div class="v-center space-content">
              <h2>Related companies</h2>
              <div class="statusPercent">
                <span class="text">28 Good, 4 Warnings, 0 Critical</span>
                <div class="line">
                  <div class="bg-good" style="width:80%"></div>
                  <div class="bg-warning" style="width:20%"></div>
                  <div class="bg-critical" style="width:0%"></div>
                </div>
              </div>
            </div>
            <el-button class="right-button btn-large" type="text">
              <i class="el-icon-plus" v-if="!relatedCompaniesOpen" v-on:click="relatedCompaniesOpen = true"></i>
              <i class="el-icon-minus" v-if="relatedCompaniesOpen" v-on:click="relatedCompaniesOpen = false"></i>
            </el-button>
          </div>
          <div class="body" v-if="relatedCompaniesOpen">
            <table class="simple-table">
              <tbody>
                <tr v-for="company in result.relatedCompanies" v-bind:key="company.vat">
                  <td class="table-text-large">{{ company.name }}</td>
                  <td>
                    <div class="flex-row">
                      <div>
                        <div class="table-text-bold">VAT</div>
                        <div class="table-text-bold">Business</div>
                      </div>
                      <div>
                        <div>{{ company.vat }}</div>
                        <div>{{ company.business }}</div>
                      </div>
                    </div>
                  </td>
                  <td align="right">
                    <i class="status-icon" v-bind:class="status[company.status.type]" v-bind:title="company.status.source"></i>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <div slot="header" class="v-center">
            <div class="v-center space-content">
              <h2>Detected lists</h2>
            </div>
            <el-button class="right-button btn-large" type="text">
              <i class="el-icon-plus" v-if="!detectedListsOpen" v-on:click="detectedListsOpen = true"></i>
              <i class="el-icon-minus" v-if="detectedListsOpen" v-on:click="detectedListsOpen = false"></i>
            </el-button>
          </div>
          <div class="body" v-if="detectedListsOpen">
            <table class="simple-table">
              <tbody>
                <tr v-for="info in result.detectedLists" v-bind:key="info.name">
                  <td class="table-text-large"><a target="_blank" :href="info.source">{{ info.name }}</a></td>
                  <td>
                    <div class="flex-row">
                      <div>
                        <div class="table-text-bold">Type</div>
                        <div class="table-text-bold">Updated</div>
                      </div>
                      <div>
                        <div>{{ info.type }}</div>
                        <div>{{ info.updated }}</div>
                      </div>
                    </div>
                  </td>
                  <td align="right">
                    <i class="status-icon" v-bind:class="status[info.iconType]" v-bind:title="info.source"></i>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>
    </el-row>
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
    relatedCompaniesOpen = false;
    detectedListsOpen = false;

    data() {
      return {
        status: {
          good: "el-icon-success",
          warning: "el-icon-warning",
          critical: "el-icon-error"
        },
        result: {
          fullName: "Person full name goes here",
          country: "Denmark",
          updated: "15-11-2019 16:14",
          functions: ["UK Prime Minister", "Secondary job"],
          relatedCompanies: [
            {
              name: "Mother Company A/S",
              vat: "DK-12345678",
              business: "University",
              status: {
                type: "warning",
                source: "EU PEP List 2019, 16-11-2019"
              }
            },
            {
              name: "Comp Any",
              vat: "DK-23456789",
              business: "IT Firm",
              status: {
                type: "good",
                source: null
              }
            },
            {
              name: "Dansk Fiskeforening",
              vat: "DK-34567890",
              business: "Union",
              status: {
                type: "good",
                source: null
              }
            }
          ],
          detectedLists: [
            {
              name: "EU PEP List 2019",
              updated: "15-11-2019",
              type: "PEP",
              source: "https://example.org/link/to/pdf",
              iconType: "warning"
            },
            {
              name: "UN PEP List 2019, Q4",
              updated: "12-10-2019",
              type: "PEP",
              source: "https://example.org/link/to/pdf",
              iconType: "warning"
            }
          ]
        }
      };
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
