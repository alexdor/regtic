<template>
  <tr :class="'beneficiary-list-item ' + data.entityType">
    <td class="value"><div class="entity-type" :title="data.entityType" :aria-label="data.entityType"></div> {{data.name}}</td>
    <td><span class="key">Ownership:</span> <span class="value">{{data.ownership * 100}}%</span></td>
    <td><span class="key">Voting rights:</span> <span class="value">{{data.voting_rights * 100}}%</span></td>
    <td><StatusIcon :data="data.status"></StatusIcon></td>
  </tr>
</template>

<script>
  import StatusIcon from "../components/StatusIcon.vue";

  export default {
    components: {
      StatusIcon
    },
    props: {
      data: {
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

  .beneficiary-list-item {
    font-family: 'Poppins', sans-serif;
    width: 100%;
  }

  td {
    padding: 0.5rem 0;
  }

  .entity-type {
    display: inline-block;
    width: 1.25rem;
    height: 1.25rem;
    background-size: contain;
    margin-right: 1rem;
  }

  .beneficiary-list-item.company .entity-type {
    background-image: url('../assets/icon-company.svg');
  }

  .beneficiary-list-item.person .entity-type {
    background-image: url('../assets/icon-person.svg');
  }

  .value {
    color: $black;
  }

  .status-icon {
    border-radius: 50%;
    width: 1.375rem;
    height: 1.375rem;
    background-repeat: no-repeat;
    background-position: center;
  }

    .status-icon.good {
      background-color: $green;
      background-image: url('../assets/icon-good.svg');
    }

    .status-icon.warning {
      background-color: $orange;
      background-image: url('../assets/icon-warning.svg');
    }

    .status-icon.bad {
      background-color: $red;
      background-image: url('../assets/icon-bad.svg');
    }
</style>
