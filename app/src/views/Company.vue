<template>
  <div class="full-height">
    <div class="canvas">
      <EntityCard :data="testCompany" x="10" y="60"></EntityCard>
      <EntityCard :data="testPerson" x="20" y="300"></EntityCard>
    </div>
  </div>
</template>

<script>
  import api from "../utils/mockapi";
  import EntityCard from "../components/EntityCard.vue";

  export default {
    components: {
      EntityCard
    },
    data() {
      return {
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
        return company;
      },
      parsePerson(person, status) {
        person.entityType = "person";
        person.status = status;
        person.relation = person.relation || "Ultimate Beneficial Owner";
        return person;
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
