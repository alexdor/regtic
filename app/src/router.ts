import Vue from "vue";
import Router from "vue-router";
import Search from "@/views/Search.vue";
import Frontpage from "@/views/Frontpage.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      meta: { layout: "blank" },
      component: Frontpage
    },
    {
      path: "/search",
      name: "home",
      component: Search
    },
    {
      path: "/search/:name",
      name: "search",
      component: () => import("./views/SearchResults.vue")
    },
    {
      path: "/company/:id",
      name: "company",
      component: () => import("./views/Company.vue")
    },
    {
      path: "/watchlist",
      name: "watchlist",
      component: () => import("./views/Watchlist.vue")
    }
  ]
});
