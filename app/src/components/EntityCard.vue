<template>
  <div :class="'entity-card ' + data.entityType + (active ? ' active' : '')" :style="'left: ' + x + 'px; top: ' + y + 'px'" v-on:click="setActive()">
    <div class="flex-row flex-row-space">
      <div class="flex-row-inline">
        <div class="entity-type" :aria-label="data.entityType" :title="data.entityType"></div>
        <span class="title">{{data.full_name || data.name}}</span>
      </div>
      <StatusIcon :status="data.status" :statusType="data.statusType" :source="data.source" :statusNotes="data.statusNotes"></StatusIcon>
    </div>
    <table class="info" v-show="openState">
      <tr v-for="info in info[data.entityType]" v-bind:key="info.key">
        <td class="key">{{info.name}}</td>
        <td class="value" v-if="info.type == 'literal'">{{data[info.key]}}</td>
        <td class="value" v-if="info.type == 'country'">{{getCountry(data[info.key])}}</td>
        <td class="value" v-if="info.type == 'percent'">{{data[info.key] * 100}}%</td>
      </tr>
    </table>
    <a class="expand-collapse expand" v-show="!openState" aria-hidden="true" title="Expand" v-on:click.stop="openState = !openState"><i class="el-icon-arrow-down"></i></a>
    <a class="expand-collapse collapse" v-show="openState" aria-hidden="true" title="Collapse" v-on:click.stop="openState = !openState"><i class="el-icon-arrow-up"></i></a>
  </div>
</template>

<script>
  import StatusIcon from "../components/StatusIcon.vue";
  import store from "../store";

  export default {
    components: {
      StatusIcon
    },
    props: {
      data: {
        type: Object
      },
      x: Number,
      y: Number,
      open: Boolean,
    },
    data() {
      return {
        active: false,
        openState: this.$props.open,
        info: {
          company: [
            { name: "VAT", key: "vat", type: "literal" },
            { name: "Country", key: "countryCode", type: "country" },
            { name: "Address", key: "address", type: "literal" },
            { name: "Type", key: "type", type: "literal" }
          ],
          person: [
            { name: "Address", key: "address", type: "literal" },
            { name: "Country", key: "countryCode", type: "country" }
          ]
        }
      };
    },
    methods: {
      setActive() {
        this.$parent.entityCardSelected(this);
      },
      getCountry(code) {
        const found = store.state.countries.filter(entry => entry.alpha2Code == code);
        if (found.length > 0) return found[0].name + " / " + code;
        else return "Unknown / ZZ";
      }
    },
    watch: {
      open(val) {
        this.openState = val;
      }
    }
  }
</script>

<style lang="scss" scoped>
  $black: #303133;
  $border: #DCDFE5;
  $border-active: #1989FA;

  $green: #67C23A;
  $orange: #E6A23C;
  $red: #F56C6C;

  $card-bg: white;

  $card-width-company: 320px;
  $card-width-person: 390px;
  $card-shadow: 0 0 0.5rem rgba(0, 0, 0, 0.15);

  .entity-card {
    position: absolute;
    padding: 1rem;
    border-radius: 1.125rem;
    font-family: 'Poppins', sans-serif;
    box-shadow: $card-shadow;
    width: $card-width-company;
    background-color: $card-bg;
    border: 1px solid $border;
    margin: 1px;
    margin-bottom: 34px;
  }

  .entity-card.person {
    width: $card-width-person;
  }

  .entity-card.active {
    border-color: $border-active;
    border-width: 2px;
    margin: 0;
    margin-bottom: 33px;
  }

  .title {
    font-size: 1.125rem;
    font-weight: bold;
    display: inline;
    color: $black;
  }

  .flex-row {
    display: flex;
    flex-direction: row;
    align-content: center;
    align-items: center;
  }

  .flex-row-space {
    justify-content: space-between;
  }

  .flex-row-inline {
    display: inline-flex;
  }

  .expand-collapse {
    display: block;
    position: absolute;
    width: 1.25rem;
    height: 1.25rem;
    background-color: $card-bg;
    border: 1px solid $border;
    border-radius: 50%;
    margin-left: calc(#{$card-width-company} / 2);
    transform: translate(-50%, 25%);
    text-align: center;
    cursor: pointer;
  }

  .entity-card.person .expand-collapse {
    margin-left: calc(#{$card-width-person} / 2);
  }

  .entity-card.active .expand-collapse {
    border-color: $border-active;
  }

  .expand-collapse i {
    display: inline-block;
    position: relative;
  }

    .expand-collapse.expand i {
      top: 0.125rem;
    }

  table {
    border-spacing: 0;
    font-size: 0.875rem;
    margin-left: 2.5rem;
    margin-top: 0.5rem;
    margin-bottom: 0.5rem;
  }

  tr td {
    padding-top: 0.5rem;
  }

  .key {
    padding-right: 0.75rem;
    white-space: nowrap;
  }

  .value {
    color: $black;
  }

  .entity-type {
    display: inline-block;
    width: 1.5rem;
    height: 1.5rem;
    background-size: contain;
    margin-right: 1rem;
  }

   .entity-card.company .entity-type {
      background-image: url('../assets/icon-company.svg');
    }

   .entity-card.person .entity-type {
      background-image: url('../assets/icon-person.svg');
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
