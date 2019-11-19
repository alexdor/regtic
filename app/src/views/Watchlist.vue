<template>
  <div class="full-height">
    <el-row>
      <el-col :span="18"><h1>Watchlist</h1></el-col>
      <el-col :span="6" align="right">
        <el-button type="primary" icon="el-icon-download"
          >Generate report</el-button
        >
      </el-col>
    </el-row>
    <el-row>
      <el-col v-for="company in items" :key="company.id" :span="12">
        <el-card>
          <span class="text-large">{{ company.name }}</span>
          <el-popover
            v-model="company.deleteVisible"
            placement="top"
            width="200"
            trigger="click"
          >
            <p>Are you sure to delete this?</p>
            <div style="text-align: right; margin: 0">
              <el-button
                size="mini"
                type="text"
                @click="company.deleteVisible = false"
                >cancel</el-button
              >
              <el-button
                type="danger"
                size="mini"
                @click="company.deleteVisible = false"
                >confirm</el-button
              >
            </div>
            <el-button
              slot="reference"
              class="right-button btn-margin-left"
              type="danger"
              icon="el-icon-delete"
              circle
            ></el-button>
          </el-popover>
          <router-link :to="'/check/' + company.id">
            <el-button
              class="right-button"
              type="primary"
              icon="el-icon-check"
              round
              >Check</el-button
            >
          </router-link>
          <div class="body">
            <p>{{ company.address }}</p>
            <p>{{ company.vat }}</p>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Topbar from "@/components/Topbar.vue"; // @ is an alias to /src

@Component({
  components: {
    Topbar
  }
})
export default class Home extends Vue {
  data() {
    return {
      items: [
        {
          vat: "DK-12345678",
          address: "Demo Street 1",
          name: "Bad company example",
          id: "1234-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-23456789",
          address: "Demo Street 2",
          name: "Other company demo",
          id: "2345-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-34567890",
          address: "Demo Street 5",
          name: "Example company",
          id: "3456-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-12345678",
          address: "AML Street 1",
          name: "Bad company example 2",
          id: "4567-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-23456789",
          address: "AML Street 2",
          name: "Other company demo 2",
          id: "5678-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-34567890",
          address: "AML Street 3",
          name: "Example company 2",
          id: "6789-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-12345678",
          address: "AML Street 5",
          name: "Bad company example 3",
          id: "7890-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-23456789",
          address: "AML Street 111",
          name: "Other company demo 3",
          id: "8901-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-34567890",
          address: "AML Street 101",
          name: "Example company 3",
          id: "9012-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-12345678",
          address: "AML Street 101",
          name: "Bad company example 4",
          id: "0123-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-23456789",
          address: "AML Street 100",
          name: "Other company demo 4",
          id: "1223-5678-9012-3456",
          deleteVisible: false
        },
        {
          vat: "DK-34567890",
          address: "AML Street 4",
          name: "Example company 4",
          id: "1224-5678-9012-3456",
          deleteVisible: false
        }
      ]
    };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.text-large {
  display: inline-block;
  font-size: 18px;
  margin-bottom: 6px;
}

.el-card p {
  margin: 0;
  color: #777779;
}

.btn-margin-left {
  margin-left: 15px;
}

.el-popover p {
  margin-top: 0;
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

.el-card {
  margin: 10px;
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
</style>
