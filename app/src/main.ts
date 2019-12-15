import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "./registerServiceWorker";
import "./plugins/element.js";

import api from "./utils/mockapi";

import DefaultLayout from "./layouts/Default.vue";
import BlankLayout from "./layouts/Blank.vue";

Vue.config.productionTip = false;

Vue.component("default-layout", DefaultLayout);
Vue.component("blank-layout", BlankLayout);

new Vue({
  router,
  store,
  render: h => h(App),
  created() {
    this.fetchCountries();
  },
  methods: {
    fetchCountries() {
      api.getCountries(function (result: any) { store.state.countries = result.data; });
    }
  }
}).$mount("#root");
