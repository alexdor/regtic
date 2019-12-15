<template>
  <tr :class="'beneficiary-list-item ' + entity.entityType">
    <td class="value"><div class="entity-type" :title="entity.entityType" :aria-label="entity.entityType"></div> <router-link :to="'/' + entity.entityType.toLowerCase() + '/' + entity.id + '/info'">{{entity.name}}</router-link></td>
    <td><span class="key">Ownership:</span> <span class="value">{{percent(relation.ownership)}}</span></td>
    <td><span class="key">Voting rights:</span> <span class="value">{{percent(relation.votingRights)}}</span></td>
    <td><StatusIcon :status="entity.checkStatus" :statusType="entity.statusType" :source="entity.source" :statusNotes="entity.statusNotes" class="pull-right"></StatusIcon></td>
  </tr>
</template>

<script>
  import StatusIcon from "../components/StatusIcon.vue";

  export default {
    components: {
      StatusIcon
    },
    props: {
      entity: {
        type: Object
      },
      relation: {
        type: Object
      }
    },
    methods: {
      percent(value) {
        return Math.round(value * 100) + "%";
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

  .beneficiary-list-item {
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

  .beneficiary-list-item.COMPANY .entity-type {
    background-image: url('../assets/icon-company.svg');
  }

  .beneficiary-list-item.PERSON .entity-type {
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

    .status-icon.OK {
      background-color: $green;
      background-image: url('../assets/icon-good.svg');
    }

    .status-icon.WARNING {
      background-color: $orange;
      background-image: url('../assets/icon-warning.svg');
    }

    .status-icon.ISSUE {
      background-color: $red;
      background-image: url('../assets/icon-bad.svg');
    }
</style>
