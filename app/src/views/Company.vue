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
    <div class="company-info">
      Test!
    </div>
  </div>
</template>

<script>
  import Vue from "vue";
  import api from "../utils/mockapi";
  import EntityCard from "../components/EntityCard.vue";
  import LineCurve from "../components/LineCurve.vue";

  export default {
    data() {
      return {
        expanded: "",
        result: {},
        companyCards: [],
        cards: [],
        selectedCompany: null,
        loading: false
      };
    },
    async mounted() {
      console.log(this);
      const canvas = this.$el.querySelector(".canvas");

      const companyId = this.$route.params.id;
      this.result = await api.validateCompany(companyId);
      
      const companyGrid = [];
      const visitedCompanies = {};

      const companyCardHeightBase = 260;
      const personCardHeightBase = 240;

      const companyCardWidthBase = 354;
      const personCardWidthBase = 424;

      const companyCardWidth = 352;
      const companyCardClosedHeightCenter = 31;
      
      this.mapCompaniesToGrid(this.result.companies, companyGrid, visitedCompanies, this.result.info, 0);
      let maxHeight = 0;
      for (let depth = 0; depth < companyGrid.length; depth++)
        maxHeight = Math.max(maxHeight, companyGrid[depth].length);

      for (let depth = 0; depth < companyGrid.length; depth++) {
        for (let i = 0; i < companyGrid[depth].length; i++)
          companyGrid[depth][i].y = maxHeight * 0.5 * companyCardHeightBase + (i - (companyGrid[depth].length - 1) * 0.5) * companyCardHeightBase;
      }

      const people = this.mapPeople(this.result.people);
      for (let i = 0; i < people.length; i++) {
        people[i].x = 220 + companyGrid.length * 450;
        people[i].y = 10 + i * personCardHeightBase;
      }
      
      const EntityCardClass = Vue.extend(EntityCard);
      const LineCurveClass = Vue.extend(LineCurve);

      for (let depth = 0; depth < companyGrid.length; depth++) {
        for (let i = 0; i < companyGrid[depth].length; i++) {
          const company = companyGrid[depth][i];
          for (let ci = 0; ci < company.companies.length; ci++) {
            const otherCompany = this.companyById(company.companies[ci]);
            if (otherCompany != undefined) {
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
          }
          for (let pi = 0; pi < company.people.length; pi++) {
            const person = this.personById(people, company.people[pi]);
            if (person != undefined) {
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
        canvasWidth = Math.max(canvasWidth, this.cards[i].x + (this.cards[i].entityType == "person" ? personCardWidthBase : companyCardWidthBase) + 32);
        canvasHeight = Math.max(canvasHeight, this.cards[i].y + (this.cards[i].entityType == "person" ? personCardHeightBase : companyCardHeightBase) + 32);
      }

      console.log(canvasWidth, canvasHeight);
      canvas.style.width = canvasWidth + "px";
      //canvas.style.height = canvasHeight + "px";
    },
    methods: {
      // Finds company info by id, if we have any that matches.
      companyById(id) {
        if (this.result.info.id == id)
          return this.result.info;
        return this.result.companies.find(company => company.id == id);
      },
      personById(people, id) {
        return people.find(person => person.id == id);
      },
      // Recursively scans and maps companies into a 2D array, and avoids recursive loops.
      mapCompaniesToGrid(companies, grid, visited, company, depth) {
        if (company.id in visited) return;
        grid[depth] = grid[depth] || [];
        grid[depth].push(company);
        visited[company.id] = { depth: depth, index: grid[depth].length - 1 };
        this.parseCompany(company);
        company.x = 20 + depth * 450;
        for (let i = 0; i < company.companies.length; i++) {
          const otherCompany = this.companyById(company.companies[i]);
          if (otherCompany != undefined)
            this.mapCompaniesToGrid(companies, grid, visited, otherCompany, depth + 1);
        };
      },
      mapPeople(people) {
        const goodPeople = (people.good || []).map(person => ({
          name: person.full_name,
          ...person,
          status: "good"
        }));

        const warningPeople = (people.warning || []).map(person => ({
          name: person.full_name,
          ...person,
          status: "warning"
        }));

        const badPeople = (people.bad || []).map(person => ({
          name: person.full_name,
          ...person,
          status: "bad"
        }));

        return [...badPeople, ...warningPeople, ...goodPeople].map(this.parsePerson);
      },
      parseCompany(company) {
        company.entityType = "company";
        company.open = false;
        return company;
      },
      parsePerson(person) {
        person.entityType = "person";
        person.open = false;
        person.relation = person.relation || "Ultimate Beneficial Owner";
        return person;
      },
      // Callback for when 'expanded' is either expanded/collapsed, to toggle all cards.
      setAllExpanded(expanded) {
        for (let i = 0; i < this.cards.length; i++)
          this.cards[i].open = expanded;
      },
      entityCardSelected(card) {
        if (card.data.entityType == "company") {
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
