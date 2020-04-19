<template>
  <div id="app">
    <component :is="layout"></component>
  </div>
</template>

<script>
import EmptyLayout from "@/layouts/EmptyLayout";
import MainLayout from "@/layouts/MainLayout";

import axios from "axios";
import router from "./router";

export default {
  computed: {
    layout() {
      return (this.$route.meta.layout || "empty") + "-layout";
    }
  },
  created() {
    axios.interceptors.response.use(
      response => {
        return response;
      },
      err => {
        if (err.response) {
          if (err.response.status === 401) {
            console.log(router.currentRoute.fullPath);
            router.replace({
              path: "/login",
              query: { redirect: router.currentRoute.fullPath }
            });

            this.$store.dispatch("logout");
          } else if (err.response.status === 403) {
            console.log("403 status");
            router.replace({
              path: "/login",
              query: { redirect: router.currentRoute.fullPath }
            });
          }
        }

        return Promise.reject(err.response.data);
      }
    );
  },
  components: {
    EmptyLayout,
    MainLayout
  }
};
</script>

<style lang="scss">
@import "~materialize-css/dist/css/materialize.min.css";
@import "assets/index.css";
</style>
