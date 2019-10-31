import Vue from "vue";
import Router from "vue-router";
import CheckCompany from "@/views/CheckCompany.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      component: CheckCompany
    },
    {
      path: "/select-company/:name",
      name: "select-company",
      component: () => import("./views/SelectCompany.vue")
    },
    {
      path: "/check/:id",
      name: "check",
      component: () => import("./views/CheckResult.vue")
    },
    {
      path: "/watchlist",
      name: "watchlist",
      component: () => import("./views/Watchlist.vue")
    }
  ]
});
