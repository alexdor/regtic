<template>
  <CheckCompany />
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import CheckCompany from "@/components/CheckCompany.vue"; // @ is an alias to /src

@Component({
  components: {
    CheckCompany
  }
})
export default class Home extends Vue {}
</script>
