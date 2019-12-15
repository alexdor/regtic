<template>
  <div class="container-center container-width-medium">
    <div class="flex-row row-spacing-bottom-large row-spacing-top-small">
      <div class="company-icon"></div>
      <div class="header">{{result.info.name}}</div>
      <el-button type="primary" round class="button-small"><i class="el-icon-plus"></i> Watchlist</el-button>
    </div>
    <div class="section-row row-spacing-top-small">
      <div class="header-small">Info</div>
      <table class="padding-all-small">
        <tbody>
          <tr>
            <td class="key">VAT</td><td class="value">{{result.info.vat}}</td>
          </tr>
          <tr>
            <td class="key">Country</td><td class="value">{{getCountry(result.info.countryCode)}}</td>
          </tr>
          <tr>
            <td class="key">Address</td><td class="value">{{result.info.address}}</td>
          </tr>
          <tr>
            <td class="key">Type</td><td class="value">{{result.info.type}}</td>
          </tr>
          <tr>
            <td class="key">Beneficiaries</td><td class="value">{{result.people.length}} people, {{result.companies.length}} companies</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="section-row row-spacing-top-medium">
      <div class="header-small">Beneficiaries</div>
      <table class="padding-all-small full-width">
        <tbody>
          <BeneficiaryListItem :entity="entityById(item.id)" :relation="item" v-for="item in result.info.ownedBy" v-bind:key="item.id"></BeneficiaryListItem>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
  //import Vue from "vue";
  import api from "../utils/mockapi";
  import BeneficiaryListItem from "../components/BeneficiaryListItem.vue";
  import store from "../store";

  export default {
    components: {
      BeneficiaryListItem
    },
    data() {
      return {
        expanded: "collapsed",
        result: {
          info: {},
          companies: [],
          people: []
        },
        entities: [],
        loading: true
      };
    },
    async mounted() {
      const companyId = this.$route.params.id;
      this.result = await api.validateCompany(companyId);

      this.entities.push(...this.result.people);
      this.entities.push(...this.result.companies);

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

  .row-spacing-bottom-large {
    margin-bottom: 3rem;
  }

  .row-spacing-top-small {
    margin-top: 1rem;
  }

  .row-spacing-top-medium {
    margin-top: 3rem;
  }

  .company-icon {
    background-image: url('../assets/icon-company.svg');
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
    margin-bottom: 0.75rem;
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

  /* Default content styling, for semi-small screens (tablets) */
  @media screen and (max-width: 1400px) {
    .col-width-30 {
      width: 50%;
    }

    .col-width-70 {
      width: 50%;
    }
  }

  /* Default content styling, for small screens (phones) */
  @media screen and (max-width: 1000px) {
    .col-width-30, .col-width-70 {
      width: 100%;
      margin-right: 0;
      display: block;
    }

    .responsive-flex {
      display: inline;
    }
  }

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
