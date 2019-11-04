<template>
  <div class="wrapper">
    <div class="banner">
      <img class="logo" src="../assets/regtic-logo.png" alt="Regtic" />
    </div>
    <div class="content">
      <div>
        <h2><i class="el-icon-check" /> Check your company relations with a click of a button. Simple, easy.</h2>
        <h2><i class="el-icon-phone-outline" /> Get results of your check in seconds. Keep updated with companies.</h2>
        <ul>
          <li>Simple interface for finding companies by name.</li>
          <li>Cheap solution for your small or medium-sized company.</li>
          <li>Combines checks of multiple registers in one search.</li>
        </ul>
      </div>
      <el-container>
        <el-row>
          <el-col :span="8" v-for="info in buttons" :key="info.text">
            <router-link :to="info.destination">
              <el-card shadow="hover" :class="info.classes">
                <i :class="info.icon" />
                <h4>{{info.text}}</h4>
              </el-card>
            </router-link>
          </el-col>
        </el-row>
      </el-container>
      <el-container>
        <el-row class="extra-padding">
          <h2><i class="el-icon-message" /> Stay notified</h2>
          <el-input placeholder="Your e-mail" v-model="newsletter_email">
            <el-button slot="append" icon="el-icon-message" title="Receive e-mails from us"></el-button>
          </el-input>
        </el-row>
      </el-container>
    </div>
    <Footer />
  </div>
</template>

<script lang="ts">
  import { Component, Vue, Prop } from "vue-property-decorator";
  import Footer from "@/components/Footer.vue"; // @ is an alias to /src

  @Component({
    components: {
      Footer
    }
  })
  export default class Frontpage extends Vue {
    buttons: { icon: string, text: string, destination: string, classes: string }[] = [];
    newsletter_email: string = "";

    created() {
      this.buttons = [
        { icon: "el-icon-user", text: "Sign in", destination: "/search", classes: "" },
        { icon: "el-icon-info", text: "About us", destination: "/", classes: "disabled" },
        { icon: "el-icon-s-finance", text: "Pricing", destination: "/", classes: "disabled" }
      ];
    }
  }
</script>

<style scoped lang="scss">
  .wrapper {
    background-color: #345bcc;
    min-height: 100%;
    color: white;
  }

  .banner {
    display: block;
    background-image: url('../assets/splash_bg.jpg');
    background-size: cover;
    background-position-x: center;
    width: 100%;
    height: 600px;
    -webkit-mask-image: -webkit-gradient(linear, left 20%, left 100%, from(rgba(0,0,0,1)), to(rgba(0,0,0,0)));
  }

  .banner .logo {
    zoom: 4;
    max-width: 100%;
  }

  .content {
    max-width: 940px;
    margin: auto;
    width: calc(100% - 60px);
  }

  .el-row {
    width: 100%;
    margin-top: 16px;
  }

  .extra-padding {
    padding: 0 8px;
  }

  .el-row .el-col {
    padding: 8px;
    text-align: center;
  }

  .el-card i {
    font-size: 48px;
    margin-top: 16px;
  }

  .el-col a {
    text-decoration: none;
  }

  .el-card.disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  li {
    margin-bottom: 5px;
  }
</style>
