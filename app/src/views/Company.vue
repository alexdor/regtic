<template>
  <div class="full-height flex-vertical">
    <div class="canvas-scroll">
      <div class="canvas">
        <div class="view-mode">
          <el-radio v-model="expanded" label="collapsed">Collapsed</el-radio>
          <el-radio v-model="expanded" label="expanded">Expanded</el-radio>
        </div>
      </div>
    </div>
    <div class="company-info" v-if="selectedCompany != null">
      <div class="flex-row row-spacing-bottom-large row-spacing-top-small">
        <div class="company-icon"></div>
        <div class="header">{{selectedCompany.data.name}}</div>
        <el-button type="primary" round class="button-small"><i class="el-icon-plus"></i> Watchlist</el-button>
      </div>
      <div class="flex-row responsive-flex flex-start">
        <div class="col-width-30">
          <div class="header-small">Info</div>
          <table class="padding-all-small">
            <tbody>
              <tr>
                <td class="key">VAT</td><td class="value">{{selectedCompany.data.vat}}</td>
              </tr>
              <tr>
                <td class="key">Country</td><td class="value">{{getCountry(selectedCompany.data.countryCode)}}</td>
              </tr>
              <tr>
                <td class="key">Address</td><td class="value">{{selectedCompany.data.address}}</td>
              </tr>
              <tr>
                <td class="key">Type</td><td class="value">{{selectedCompany.data.type}}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="col-width-70">
          <div class="header-small">Beneficiaries</div>
          <table class="padding-all-small full-width">
            <tbody>
              <BeneficiaryListItem :entity="entityById(item.id)" :relation="item" v-for="item in selectedCompany.data.ownedBy" v-bind:key="item.id"></BeneficiaryListItem>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Vue from "vue";
  import api from "../utils/mockapi";
  import EntityCard from "../components/EntityCard.vue";
  import LineCurve from "../components/LineCurve.vue";
  import BeneficiaryListItem from "../components/BeneficiaryListItem.vue";
  import store from "../store";

  export default {
    components: {
      BeneficiaryListItem
    },
    data() {
      return {
        expanded: "collapsed",
        result: {},
        companyCards: [],
        cards: [],
        entities: [],
        selectedCompany: null,
        loading: false
      };
    },
    async mounted() {
      console.log(this, this.$el);
      const canvas = this.$el.querySelector(".canvas");

      const COMPANY = "COMPANY";
      const PERSON = "PERSON";

      const companyId = this.$route.params.id;
      this.result = await api.validateCompany(companyId);
      
      const companyGrid = [];
      const visitedCompanies = {};

      const companyCardHeightBase = 260;
      const personCardHeightBase = 180;

      const companyCardWidthBase = 354;
      const personCardWidthBase = 424;

      const companyCardWidth = 352;
      const companyCardClosedHeightCenter = 31;
      
      const people = this.result.people;
      this.entities.push(this.result.info);
      this.entities.push(...this.result.companies);
      this.entities.push(...people);

      this.entities.forEach(entity => entity.checkStatus = entity.checkStatus || "OK");
      this.entities.forEach(entity => { if (entity.checkStatus != "OK" && entity.source == undefined) { entity.statusNotes = entity.statusNotes || "Inherited from beneficiaries"; } });

      console.log(this.entities, people);
      
      this.mapCompaniesToGrid(this.result.companies, companyGrid, visitedCompanies, this.result.info, 0);
      let maxHeight = 0;
      for (let depth = 0; depth < companyGrid.length; depth++)
        maxHeight = Math.max(maxHeight, companyGrid[depth].length);

      for (let depth = 0; depth < companyGrid.length; depth++) {
        for (let i = 0; i < companyGrid[depth].length; i++)
          companyGrid[depth][i].y = maxHeight * 0.5 * companyCardHeightBase + (i - (companyGrid[depth].length - 1) * 0.5) * companyCardHeightBase;
      }
      
      for (let i = 0; i < people.length; i++) {
        people[i].x = 220 + companyGrid.length * 450;
        people[i].y = 10 + i * personCardHeightBase;
      }
      
      const EntityCardClass = Vue.extend(EntityCard);
      const LineCurveClass = Vue.extend(LineCurve);

      for (let depth = 0; depth < companyGrid.length; depth++) {
        for (let i = 0; i < companyGrid[depth].length; i++) {
          const company = companyGrid[depth][i];
          for (let ei = 0; ei < company.ownedBy.length; ei++) {
            const otherCompany = this.entityById(company.ownedBy[ei].id);
            if (otherCompany != undefined && otherCompany.entityType == COMPANY) {
              const line = new LineCurveClass({
                propsData: {
                  x1: company.x + companyCardWidth - 10,
                  y1: company.y + companyCardClosedHeightCenter,
                  x2: otherCompany.x + 10,
                  y2: otherCompany.y + companyCardClosedHeightCenter,
                  xWeight1: 0.5,
                  xWeight2: 0.5
                }
              });
              line.$mount();
              canvas.append(line.$el);
            }
            const person = this.entityById(company.ownedBy[ei].id);
            if (person != undefined && person.entityType == PERSON) {
              const line = new LineCurveClass({
                propsData: {
                  x1: company.x + companyCardWidth - 10,
                  y1: company.y + companyCardClosedHeightCenter,
                  x2: person.x + 10,
                  y2: person.y + companyCardClosedHeightCenter,
                  xWeight1: 0.95,
                  xWeight2: 0.75
                }
              });
              line.$mount();
              canvas.append(line.$el);
            }
          }
        }
      }
      
      // Makes sure the radio toggles matches the current view mode.
      const updateViewMode = () => {
        if (this.cards.every(card => card.openState))
          this.expanded = "expanded";
        else if (this.cards.every(card => !card.openState))
          this.expanded = "collapsed";
        else
          this.expanded = "";
      };

      for (let depth = 0; depth < companyGrid.length; depth++) {
        for (let i = 0; i < companyGrid[depth].length; i++) {
          const card = new EntityCardClass({
            propsData: {
              data: companyGrid[depth][i],
              open: companyGrid[depth][i].open,
              x: companyGrid[depth][i].x,
              y: companyGrid[depth][i].y
            }
          });
          card.$parent = this;
          card.$mount();
          canvas.append(card.$el);
          this.cards.push(card);
          this.companyCards.push(card);

          this.$watch(
            "cards." + (this.cards.length - 1) + ".openState",
            updateViewMode
          );
        }
      }

      for (let i = 0; i < people.length; i++) {
        const card = new EntityCardClass({
          propsData: {
            data: people[i],
            open: people[i].open,
            x: people[i].x,
            y: people[i].y
          }
        });
        card.$parent = this;
        card.$mount();
        canvas.append(card.$el);
        this.cards.push(card);

        this.$watch(
          "cards." + (this.cards.length - 1) + ".openState",
          updateViewMode
        );
      }

      let canvasWidth = 300;
      let canvasHeight = 200;
      for (let i = 0; i < this.cards.length; i++) {
        canvasWidth = Math.max(canvasWidth, this.cards[i].x + (this.cards[i].entityType == PERSON ? personCardWidthBase : companyCardWidthBase) + 32);
        canvasHeight = Math.max(canvasHeight, this.cards[i].y + (this.cards[i].entityType == PERSON ? personCardHeightBase : companyCardHeightBase) + 32);
      }
      
      canvas.style.width = canvasWidth + "px";
      //canvas.style.height = canvasHeight + "px";
    },
    methods: {
      entityById(id) {
        return this.entities.find(entity => entity.id == id);
      },
      // Recursively scans and maps companies into a 2D array, and avoids recursive loops.
      mapCompaniesToGrid(companies, grid, visited, company, depth) {
        const COMPANY = "COMPANY";

        if (company.id in visited) return;
        grid[depth] = grid[depth] || [];
        grid[depth].push(company);
        visited[company.id] = { depth: depth, index: grid[depth].length - 1 };
        company.x = 20 + depth * 450;
        for (let i = 0; i < company.ownedBy.length; i++) {
          const otherCompany = this.entityById(company.ownedBy[i].id);
          if (otherCompany != undefined && otherCompany.entityType == COMPANY)
            this.mapCompaniesToGrid(companies, grid, visited, otherCompany, depth + 1);
        }
      },
      getCountry(code) {
        const found = store.state.countries.filter(entry => entry.alpha2Code == code);
        if (found.length > 0) return found[0].name + " / " + code;
        else return "Unknown / ZZ";
      },
      // Callback for when 'expanded' is either expanded/collapsed, to toggle all cards.
      setAllExpanded(expanded) {
        for (let i = 0; i < this.cards.length; i++)
          this.cards[i].open = expanded;
      },
      entityCardSelected(card) {
        const COMPANY = "COMPANY";

        if (card.data.entityType == COMPANY) {
          if (this.selectedCompany != null)
            this.selectedCompany.active = false;
          if (this.selectedCompany != card)
            card.active = !card.active;
          this.selectedCompany = card.active ? card : null;
        }
      }
    },
    watch: {
      expanded: {
        handler(val) {
          if (val.length > 0) this.setAllExpanded(val == "expanded");
        }
      }
    }
  }
</script>

<style scoped lang="scss">
  $black: #303133;

  .flex-vertical {
    display: flex;
    flex-direction: column;
  }

  .canvas-scroll {
    display: block;
    position: relative;
    width: 100%;
    flex-grow: 1;
    overflow: auto;
    min-height: 40vh;
  }

  .canvas {
    padding-right: 4.25rem;
  }

  .view-mode {
    margin-left: 3rem;
  }

  .view-mode .el-radio {
    display: block;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    margin-bottom: 0.75rem;
  }

  .flex-row {
    display: flex;
    flex-direction: row;
    align-items: center;
    font-family: 'Poppins', sans-serif;
  }

  .flex-start {
    align-items: flex-start;
  }

  .row-spacing-bottom-large {
    margin-bottom: 3rem;
  }

  .row-spacing-top-small {
    margin-top: 1rem;
  }

  .company-icon {
    background-image: url('../assets/icon-company.svg');
    background-size: contain;
    background-position: center;
    background-repeat: no-repeat;
    width: 2.25rem;
    height: 2.25rem;
  }

  .header {
    font-size: 1.5rem;
    line-height: 2.5rem;
    font-weight: bold;
    color: $black;
    margin-left: 1.5rem;
  }

  .header-small {
    font-size: 1.375rem;
    margin-bottom: 0.75rem;
    color: $black;
  }

  .button-small {
    padding: 0 0.75rem;
    height: 2rem;
    font-size: 0.875rem;
    font-family: 'Poppins', sans-serif;
    margin-left: 0.75rem;
  }

  .col-width-30 {
    width: 30%;
    margin-right: 2rem;
  }

  .col-width-70 {
    width: 70%;
  }

  .padding-all-small {
    padding: 0.5rem;
  }

  tr td {
    padding-top: 0.5rem;
  }

  .key {
    padding-right: 2rem;
    white-space: nowrap;
  }

  .value {
    color: $black;
  }

  /* Default content styling, for semi-small screens (tablets) */
  @media screen and (max-width: 1400px) {
    .col-width-30 {
      width: 50%;
    }

    .col-width-70 {
      width: 50%;
    }
  }

  /* Default content styling, for small screens (phones) */
  @media screen and (max-width: 1000px) {
    .col-width-30, .col-width-70 {
      width: 100%;
      margin-right: 0;
      display: block;
    }

    .responsive-flex {
      display: inline;
    }
  }

.title-col {
  height: 40px;
  display: flex;
}

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

.people-card {
  margin: 20px 10px 10px 0px;
}

.companies-card {
  margin: 20px 0px 10px 10px;
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

.title-loading {
  height: 30px;
}
</style>
