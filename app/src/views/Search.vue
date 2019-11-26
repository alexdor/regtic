<template>
  <div class="full-height">
    <div class="col-center">
      <img
        src="../assets/undraw_people_search.svg"
        alt
        class="fit-to-width"
      />
      <div class="vertical-spacing">
        <SearchInput :include-button="true" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import SearchInput from "@/components/SearchInput.vue";

@Component({
  components: {
    SearchInput
  }
})
export default class CheckCompany extends Vue {
  test = "";
}
</script>

<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
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

.justify-center {
  text-align: center;
}

.vertical-spacing {
  margin-top: 4rem;
  margin-bottom: 2rem;
}
</style>
