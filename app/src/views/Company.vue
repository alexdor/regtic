<template>
  <div class="full-height">
    <el-row>
      <el-col :span="4">
        <div class="title">
          {{ name }}
        </div>
      </el-col>
      <el-col :span="20" align="right">
        <el-button
          disabled
          title="Coming soon..."
          type="info"
          icon="el-icon-collection-tag"
        >
          Add to watchlist
        </el-button>
        <el-button
          disabled
          title="Coming soon..."
          type="primary"
          icon="el-icon-download"
        >
          Generate report
        </el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-card>
          <div slot="header" class="clearfix">
            <span>People</span>
          </div>
          <el-table
            class="full-width"
            :data="result.people"
            :default-sort="{ prop: 'type', order: 'descending' }"
          >
            <el-table-column prop="full_name" label="Name" sortable>
            </el-table-column>
            <el-table-column
              prop="type"
              label="Status"
              sortable
              width="100"
              align="center"
            >
              <template slot-scope="scope">
                <el-tooltip
                  v-if="scope.row.type == 'pep'"
                  class="item"
                  :content="scope.row.source"
                  placement="top"
                  effect="dark"
                >
                  <el-button
                    circle
                    size="mini"
                    icon="el-icon-link"
                    type="warning"
                  ></el-button>
                </el-tooltip>
                <el-tooltip
                  v-else-if="scope.row.type == 'sanction'"
                  class="item"
                  :content="scope.row.source"
                  placement="top"
                  effect="dark"
                >
                  <el-button
                    circle
                    size="mini"
                    icon="el-icon-link"
                    type="danger"
                  ></el-button>
                </el-tooltip>
                <el-button
                  v-else
                  circle
                  size="mini"
                  icon="el-icon-check"
                  type="success"
                ></el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <div slot="header" class="clearfix">
            <span>Companies</span>
          </div>
          <el-table
            class="full-width"
            :data="result.companies"
            :default-sort="{ prop: 'type', order: 'descending' }"
          >
            <el-table-column prop="vat" label="VAT" sortable width="120">
            </el-table-column>
            <el-table-column prop="name" label="Name" sortable>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({})
export default class Home extends Vue {
  data() {
    return {
      name: "Regtic Demo",
      id: this.$route.params.id,
      typeSort: {
        ok: 0,
        pep: 1,
        sanction: 2
      },
      result: {
        companies: [
          { vat: "DK-12345678", name: "Bad company example" },
          { vat: "DK-23456789", name: "Other company demo" },
          { vat: "DK-34567890", name: "Example company" }
        ],
        people: [
          {
            full_name: "John Doe",
            source:
              "EU Sanctions List 2019 (www.example.org/path-to-file) - 23/10/2019",
            type: "sanction"
          },
          {
            full_name: "Jane Doe",
            source:
              "FN Sanctions List 2019 (www.other.org/path-to-file) - 21/10/2019",
            type: "sanction"
          },
          {
            full_name: "James Doe",
            source:
              "EU Sanctions List 2019 (www.example.org/path-to-file) - 23/10/2019",
            type: "sanction"
          },
          {
            full_name: "Donald Trump",
            source:
              "EU PEP List 2019 (www.example.org/path-to-file) - 22/10/2019",
            type: "pep"
          },
          { full_name: "John Doe", type: "ok" },
          { full_name: "Jane Doe", type: "ok" },
          { full_name: "James Doe", type: "ok" },
          { full_name: "Donald Trump", type: "ok" }
        ]
      }
    };
  }
}
</script>

<style scoped lang="scss">
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
