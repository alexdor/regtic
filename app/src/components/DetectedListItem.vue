<template>
  <tr class="detected-list-item">
    <td class="value">
      <a :href="source" target="_blank">{{ source }}</a>
    </td>
    <td>
      <span class="key">Type:</span> <span class="value">{{ type }}</span>
    </td>
    <td>
      <span class="key">Updated:</span> <span class="value">{{ updated }}</span>
    </td>
    <td>
      <StatusIcon
        :status="status"
        :status-type="type"
        :status-notes="status"
        class="pull-right"
      ></StatusIcon>
    </td>
  </tr>
</template>

<script>
import StatusIcon from "@/components/StatusIcon.vue";

export default {
  components: {
    StatusIcon
  },
  props: {
    source: String,
    type: String,
    updated: String,
    status: String
  }
};
</script>

<style lang="scss" scoped>
$black: #303133;

$green: #67c23a;
$orange: #e6a23c;
$red: #f56c6c;

$link-hover: #1871f8;

.detected-list-item {
  font-family: "Poppins", sans-serif;
  width: 100%;
}

td {
  padding: 0.5rem 0;
}

a {
  color: $black;
  text-decoration: none;
}

a:hover {
  color: $link-hover;
}

.pull-right {
  float: right;
}

.value {
  color: $black;
}
</style>
