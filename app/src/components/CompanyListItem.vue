<template>
  <router-link class="company-list-item" tag="div" :to="'/company/' + data.id">
    <span class="title">{{ data.name || "Name not found" }}</span>
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
        <td class="value">
          {{ data.type }}
        </td>
      </tr>
    </table>
  </router-link>
</template>

<script>
export default {
  props: {
    data: {
      type: Object,
      default: function () {
        return {};
      },
    },
  },
};
</script>

<style lang="scss" scoped>
$black: #303133;
$link-hover: #109cf1;

.company-list-item {
  cursor: pointer;
  padding: 1rem 0;
  margin: 1rem 0;
  font-family: "Poppins", sans-serif;
  &:hover *,
  &:focus * {
    color: $link-hover;
  }
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
</style>
