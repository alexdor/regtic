<template>
  <tr class="company-relation-list-item">
    <td class="value"><div class="entity-type" title="company" aria-label="company"></div>
                      <el-tooltip class="item" effect="dark" placement="top">
                        <div slot="content">
                          Relations
                          <p v-for="relation in company.relation" v-bind:key="relation" class="relation-text">{{relation}}</p>
                        </div>
                        <router-link :to="'/company/' + company.id + '/info'">{{company.name}}</router-link></el-tooltip></td>
    <td><span class="key">VAT:</span> <span class="value">{{company.vat}}</span></td>
    <td><span class="key">Business:</span> <span class="value">{{company.type}}</span></td>
    <td><StatusIcon :status="company.checkStatus" :statusType="company.statusType" :source="company.source" :statusNotes="company.statusNotes" class="pull-right"></StatusIcon></td>
  </tr>
</template>

<script>
  import StatusIcon from "../components/StatusIcon.vue";

  export default {
    components: {
      StatusIcon
    },
    props: {
      company: {
        type: Object
      }
    }
  }
</script>

<style lang="scss" scoped>
  $black: #303133;

  $green: #67C23A;
  $orange: #E6A23C;
  $red: #F56C6C;

  $link-hover: #1871F8;

  .company-relation-list-item {
    font-family: 'Poppins', sans-serif;
    width: 100%;
  }

  td {
    padding: 0.5rem 0;
  }

  a {
    color: $black;
    text-decoration: none;
  }

  a:hover {
    color: $link-hover;
  }

  .pull-right {
    float: right;
  }

  .entity-type {
    display: inline-block;
    width: 1.25rem;
    height: 1.25rem;
    background-size: contain;
    margin-right: 1rem;
  }

  .company-relation-list-item .entity-type {
    background-image: url('../assets/icon-company.svg');
  }

  .value {
    color: $black;
  }

  .small-text {
    font-size: 0.75rem;
  }

  .relation-text:last-child {
    margin-bottom: 0;
    display: list-item;
  }

  .relation-text:before {
    content: '• ';
  }
</style>
