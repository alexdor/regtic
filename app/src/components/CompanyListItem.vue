<template>
  <div class="company-list-item">
    <span class="title">{{ data.name }}</span>
    <table class="info">
      <tr>
        <td class="key">
          VAT
        </td>
        <td class="value">
          {{ data.vat }}
        </td>
        <td class="key">
          Address
        </td>
        <td class="value">
          {{ data.address }}
        </td>
      </tr>
      <tr>
        <td class="key">
          Country
        </td>
        <td class="value">
          {{ data.country }}
        </td>
        <td class="key">
          Type
        </td>
        <td class="type">
          {{ data.type }}
        </td>
      </tr>
    </table>

    <router-link class="check-button" :to="'/company/' + data.id">
      <el-button
        class="check-button"
        type="primary"
        icon="el-icon-search"
        aria-label="Check"
        title="Check"
      />
    </router-link>
  </div>
</template>

<script>
export default {
  props: {
    data: {
      type: Object,
      default: () => {}
    }
  }
};
</script>

<style lang="scss" scoped>
$black: #303133;

.company-list-item {
  padding: 1rem 0;
  margin: 1rem 0;
  font-family: "Poppins", sans-serif;
}

.title {
  font-size: 1.125rem;
  font-weight: bold;
  display: block;
  color: $black;
  margin-bottom: 0.5rem;
}

table {
  border-spacing: 0;
  font-size: 0.875rem;
}

.key {
  padding-right: 0.75rem;
}

.value {
  color: $black;
  padding-right: 1.5rem;
}

.check-button {
  width: 2rem;
  height: 2rem;
  font-size: 1rem;
  padding: 0;
  position: absolute;
  right: 0;
  margin-top: -2rem;
}
</style>
