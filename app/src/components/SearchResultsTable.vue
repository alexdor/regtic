<template>
  <el-table :data="data">
    <el-table-column label="Name" show-overflow-tooltip>
      <template slot-scope="scope">
        <span class="name">
          {{ scope.row.name }}
        </span>
      </template>
    </el-table-column>
    <el-table-column prop="address" label="Address" show-overflow-tooltip />
    <el-table-column prop="vat" label="VAT" width="150" show-overflow-tooltip />
    <el-table-column fixed="right" width="120">
      <template slot-scope="scope">
        <router-link :to="'/company/' + scope.row.id">
          <el-button type="primary" size="medium">
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
