<template>
  <el-table :data="data">
    <el-table-column label="Name">
      <template slot-scope="scope">
        <span class="name">
          {{ scope.row.name }}
        </span>
      </template>
    </el-table-column>
    <el-table-column prop="address" label="Address" />
    <el-table-column prop="vat" label="VAT" />
    <el-table-column fixed="right" width="90">
      <template slot-scope="scope">
        <router-link :to="'/company/' + scope.row.id">
          <el-button type="text">
            Check
          </el-button>
        </router-link>
      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class SearchResultsTable extends Vue {
  @Prop({ default: false })
  data!: [];
}
</script>

<style lang="scss" scoped>
.el-table {
  width: 100%;
}

.name {
  font-weight: 700;
}
</style>
