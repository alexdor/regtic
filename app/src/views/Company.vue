<template>
  <div class="full-height">
    <div class="canvas">
      <div class="view-mode">
        <el-radio v-model="expanded" label="collapsed">Collapsed</el-radio>
        <el-radio v-model="expanded" label="expanded">Expanded</el-radio>
      </div>
      <LineCurve :x1="720" :y1="191" :x2="920" :y2="231"></LineCurve>
      <EntityCard :data="testCompany" :x="400" :y="160" :open="testCompany.open"></EntityCard>
      <EntityCard :data="testPerson" :x="920" :y="200" :open="testCompany.open"></EntityCard>
    </div>
  </div>
</template>

<script>
  import api from "../utils/mockapi";
  import EntityCard from "../components/EntityCard.vue";
  import LineCurve from "../components/LineCurve.vue";

  export default {
    components: {
      EntityCard,
      LineCurve
    },
    data() {
      return {
        expanded: "",
        result: {},
        testCompany: {},
        testPerson: {},
        loading: false
      };
    },
    async created() {
      const companyId = this.$route.params.id;
      this.result = await api.validateCompany(companyId);
      this.testCompany = this.parseCompany(this.result.companies[0]);
      this.testPerson = this.parsePerson(this.result.people.bad[0], "bad");
    },
    methods: {
      parseCompany(company) {
        company.entityType = "company";
        company.open = false;
        return company;
      },
      parsePerson(person, status) {
        person.entityType = "person";
        person.open = false;
        person.status = status;
        person.relation = person.relation || "Ultimate Beneficial Owner";
        return person;
      },
      setAllExpanded(expanded) {
        this.testCompany.open = expanded;
        this.testPerson.open = expanded;
      }
    },
    watch: {
      expanded: {
        handler(val) {
          if (val.length > 0) this.setAllExpanded(val == "expanded");
        }
      }
    }
  }
</script>

<style scoped lang="scss">
  .canvas {
    display: flex;
    position: relative;
    width: 100%;
    height: -webkit-fill-available;
    overflow: auto;
  }

  .view-mode {
    margin-left: 3rem;
  }

  .view-mode .el-radio {
    display: block;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    margin-bottom: 0.75rem;
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
