import Vue from "vue";
import Vuex from "vuex";
import { countries } from "./countries";
Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    countries
  },
  mutations: {},
  actions: {}
});
