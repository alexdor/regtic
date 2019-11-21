import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      meta: { layout: "blank" },
      component: () => import("@/views/Frontpage.vue")
    },
    {
      path: "/search",
      name: "search",
      component: () => import("@/views/Search.vue")
    },
    {
      path: "/search/:name",
      name: "search-name",
      component: () => import("@/views/SearchResults.vue")
    },
    {
      path: "/company/:id",
      name: "company",
      component: () => import("@/views/Company.vue")
    },
    {
      path: "/watchlist",
      name: "watchlist",
      component: () => import("@/views/Watchlist.vue")
    }
  ]
});
