import Vue from "vue";
import App from "./App.vue";
import BlankLayout from "./layouts/Blank.vue";
import DefaultLayout from "./layouts/Default.vue";
import "./plugins/element.js";
import "./registerServiceWorker";
import router from "./router";

Vue.config.productionTip = false;

Vue.component("default-layout", DefaultLayout);
Vue.component("blank-layout", BlankLayout);

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#root");
