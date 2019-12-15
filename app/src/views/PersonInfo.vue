<template>
  <div class="container-center container-width-medium">
    <div class="flex-row row-spacing-bottom-large row-spacing-top-small">
      <div class="person-icon"></div>
      <div class="header">{{result.name}}</div>
      <StatusIcon :status="result.checkStatus" :statusType="result.statusType" :source="result.source" :statusNotes="result.statusNotes" class="row-spacing-left-small"></StatusIcon>
    </div>
    <div class="section-row row-spacing-top-small">
      <div class="header-small">Info</div>
      <table class="padding-all-small">
        <tbody>
          <tr>
            <td class="key">VAT</td><td class="value">{{result.vat}}</td>
          </tr>
          <tr>
            <td class="key">Country</td><td class="value">{{getCountry(result.countryCode)}}</td>
          </tr>
          <tr>
            <td class="key">Address</td><td class="value">{{result.address}}</td>
          </tr>
          <tr>
            <td class="key">Type</td><td class="value">{{result.type}}</td>
          </tr>
          <tr>
            <td class="key">Status of owning entities</td><td class="value">{{status.ok}} OK, {{status.warning}} Warning, {{status.issue}} Issue
            <StatusBar :ok="status.ok" :warning="status.warning" :issue="status.issue"></StatusBar></td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="section-row row-spacing-top-medium">
      <div class="header-small">
        Related companies
        <a class="expand-collapse expand" v-show="!companiesOpen" aria-hidden="true" title="Expand" v-on:click.stop="companiesOpen = !companiesOpen"><i class="el-icon-arrow-down"></i></a>
        <a class="expand-collapse collapse" v-show="companiesOpen" aria-hidden="true" title="Collapse" v-on:click.stop="companiesOpen = !companiesOpen"><i class="el-icon-arrow-up"></i></a>
      </div>
      <table class="padding-all-small full-width" v-show="companiesOpen">
        <tbody>
          <CompanyRelationListItem :company="item" v-for="item in result.owns" v-bind:key="item.id"></CompanyRelationListItem>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
  //import Vue from "vue";
  import api from "../utils/mockapi";
  import CompanyRelationListItem from "../components/CompanyRelationListItem.vue";
  import StatusIcon from "../components/StatusIcon.vue";
  import StatusBar from "../components/StatusBar.vue";
  import store from "../store";

  export default {
    components: {
      CompanyRelationListItem,
      StatusIcon,
      StatusBar
    },
    data() {
      return {
        companiesOpen: true,
        detectedListsOpen: true,
        result: {
          info: {},
          companies: [],
          people: []
        },
        entities: [],
        status: {
          ok: 0,
          warning: 0,
          issue: 0
        },
        loading: true
      };
    },
    async mounted() {
      const personId = this.$route.params.id;
      this.result = await api.getPerson(personId);
      console.log(this.result);
      
      this.status.ok = this.result.owns.filter(entity => entity.status == 'ok').length;
      this.status.warning = this.result.owns.filter(entity => entity.status == 'warning').length;
      this.status.issue = this.result.owns.filter(entity => entity.status == 'issue').length;

      this.loading = false;
    },
    methods: {
      entityById(id) {
        return this.entities.find(entity => entity.id == id);
      },
      getCountry(code) {
        const found = store.state.countries.filter(entry => entry.alpha2Code == code);
        if (found.length > 0) return found[0].name + " / " + code;
        else return "Unknown / ZZ";
      }
    }
  }
</script>

<style scoped lang="scss">
  $black: #303133;
  $blue: #1989FA;

  .flex-vertical {
    display: flex;
    flex-direction: column;
  }

  .canvas-scroll {
    display: block;
    position: relative;
    width: 100%;
    flex-grow: 1;
    overflow: auto;
  }

  .flex-row {
    display: flex;
    flex-direction: row;
    align-items: center;
    font-family: 'Poppins', sans-serif;
  }

  .section-row {
    font-family: 'Poppins', sans-serif;
  }

  .flex-start {
    align-items: flex-start;
  }

  .pull-right {
    position: absolute;
    right: 0;
  }

  .row-spacing-left-small {
    margin-left: 1rem;
  }

  .row-spacing-bottom-large {
    margin-bottom: 3rem;
  }

  .row-spacing-top-small {
    margin-top: 1rem;
  }

  .row-spacing-top-medium {
    margin-top: 3rem;
  }

  .expand-collapse {
    color: $blue;
    float: right;
    margin-right: 0.5rem;
  }

  .expand-collapse i {
    display: inline-block;
    position: relative;
  }

  .expand-collapse.expand i {
    top: -0.125rem;
  }

  .expand-collapse.collapse i {
    top: -0.25rem;
  }

  .person-icon {
    background-image: url('../assets/icon-person.svg');
    background-size: contain;
    background-position: center;
    background-repeat: no-repeat;
    width: 2.25rem;
    height: 2.25rem;
  }

  .header {
    font-size: 1.5rem;
    line-height: 2.5rem;
    font-weight: bold;
    color: $black;
    margin-left: 1.5rem;
  }

  .header-small {
    font-size: 1.375rem;
    font-weight: bold;
    color: $black;
    font-family: 'Poppins', sans-serif;
  }

  .button-small {
    padding: 0 0.75rem;
    height: 2rem;
    font-size: 0.875rem;
    font-family: 'Poppins', sans-serif;
    margin-left: 0.75rem;
  }

  .col-width-30 {
    width: 30%;
    margin-right: 2rem;
  }

  .col-width-70 {
    width: 70%;
  }

  .padding-all-small {
    padding: 0.5rem;
  }

  tr td {
    padding-top: 0.5rem;
  }

  .key {
    padding-right: 2rem;
    white-space: nowrap;
  }

  .value {
    color: $black;
  }

.full-width {
  width: 100%;
}

.title-loading {
  height: 30px;
}
</style>
