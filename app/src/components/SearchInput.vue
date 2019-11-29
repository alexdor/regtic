<template>
  <div class="searchInput">
    <el-input
      v-if="includeButton"
      v-model="searchStr"
      class="inline-input full-with"
      placeholder="Enter company name or VAT"
      prefix-icon="el-icon-search"
      clearable
      @keyup.native.enter="search()"
    />
    <el-input
      v-if="!includeButton"
      v-model="searchStr"
      class="inline-input full-with"
      placeholder="Enter company name or VAT"
      suffix-icon="el-icon-search"
      clearable
      @keyup.native.enter="search()"
    />
    <div v-if="includeButton" class="vertical-spacing">
      <el-button
        type="primary"
        class="full-width"
        :disabled="!isSearchStrValid"
        @click.native="search()"
        >Search Company</el-button
      >
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class SearchInput extends Vue {
  @Prop({ default: false })
  includeButton!: boolean;

  @Prop({ default: "" })
  initialSearchStr!: string;

  searchStr = "";

  search(): void {
    if (!this.isSearchStrValid) return;

    this.$emit("search", encodeURIComponent(this.searchStr));
    this.$router.push(`/search/${encodeURIComponent(this.searchStr)}`);
  }

  get isSearchStrValid(): boolean {
    return this.searchStr.length > 0;
  }

  mounted(): void {
    const isInitialSearchStrValid = this.initialSearchStr.length > 0;
    if (!isInitialSearchStrValid) return;

    this.searchStr = this.initialSearchStr;
  }
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
  margin: 20px 0;
}

.el-button {
  font-weight: 700;
  font-family: 'Poppins', sans-serif;
}
</style>
