<template>
  <div class="wrapper">
    <div class="banner">
      <img class="logo" src="../assets/regtic-logo.png" alt="Regtic" />
    </div>
    <div class="content">
      <div>
        <h2>
          <i class="el-icon-check" /> Check your company relations with a click
          of a button. Simple, easy.
        </h2>
        <h2>
          <i class="el-icon-phone-outline" /> Get results of your check in
          seconds. Keep updated with companies.
        </h2>
        <ul>
          <li>Simple interface for finding companies by name.</li>
          <li>Cheap solution for your small or medium-sized company.</li>
          <li>Combines checks of multiple registers in one search.</li>
        </ul>
      </div>
      <el-container>
        <el-row>
          <el-col v-for="info in buttons" :key="info.text" :span="8">
            <router-link :to="info.destination">
              <el-card shadow="hover" :class="info.classes">
                <i :class="info.icon" />
                <h4>{{ info.text }}</h4>
              </el-card>
            </router-link>
          </el-col>
        </el-row>
      </el-container>
      <el-container>
        <el-row class="extra-padding">
          <h2><i class="el-icon-message" /> Stay notified</h2>
          <el-input v-model="newsletter_email" placeholder="Your e-mail">
            <el-button
              slot="append"
              icon="el-icon-message"
              title="Receive e-mails from us"
              v-on:click="sendMailToSlack"
            ></el-button>
          </el-input>
        </el-row>
      </el-container>
    </div>
    <Footer />
  </div>
</template>

<script lang="ts">
  import Footer from "@/components/Footer.vue";
  import api from "@/utils/api";

 import Vue from "vue"

  export default Vue.extend({
    components: {
      Footer
    },
    
    data: function () {
      return {
        newsletter_email: "",
        buttons: [
          {
            icon: "el-icon-user",
            text: "View demo",
            destination: "/search",
            classes: ""
          },
          {
            icon: "el-icon-info",
            text: "About us",
            destination: "/",
            classes: "disabled"
          },
          {
            icon: "el-icon-s-finance",
            text: "Pricing",
            destination: "/",
            classes: "disabled"
          }
        ]
      };
    },

    methods: {
      async sendMailToSlack() {
        api.sendMailToSlack(this.newsletter_email);
      }
    }
  })
</script>

<style scoped lang="scss">
$side-spacing: 30px; /* The spacing left/right of the content section, in case the screen width is too low. */

.wrapper {
  background-color: #345bcc;
  min-height: 100%;
  color: white;
}

.banner {
  display: block;
  background-image: url("../assets/splash_bg.jpg");
  background-size: cover;
  background-position-x: center;
  width: 100%;
  height: 30vh;
  -webkit-mask-image: -webkit-gradient(
    linear,
    left 20%,
    left 100%,
    from(rgba(0, 0, 0, 1)),
    to(rgba(0, 0, 0, 0))
  );
}

.banner .logo {
  max-width: 60vh; /* We make use of the max-width constraint to make sure the image does not overflow outside of the banner area, as the banner height is now using a relative metric. */
  width: 50vw;
  padding-top: 0.5vh;
  position: absolute;
  left: 50%;
  transform: translate(-50%, 0);
}

/* Default content styling, for quite small screens (tablets) */
.content {
  max-width: 640px;
  margin: auto;
  width: calc(
    100% - #{$side-spacing} * 2
  ); /* Adding spacing left and right of the content section, in case the screen width is too low. */
}

/* Default content styling, for small screens (phones) */
@media screen and (max-width: 800px) {
  .content {
    min-width: 400px;
    zoom: 0.8; /* Zoom to make the buttons appear less squished. */
  }
}

/* Default content styling, for small-ish medium screens (tablets) */
@media screen and (min-width: 1200px) {
  .content {
    min-width: 940px;
  }
}

/* Default content styling, for normal screens (laptop, desktop) */
@media screen and (min-width: 1600px) {
  .content {
    min-width: 1200px;
    zoom: 1.25;
  }
}

/* Default content styling, for large screens (desktop) */
@media screen and (min-width: 3600px) {
  .content {
    min-width: 1800px;
    zoom: 2;
  }

  .el-card i {
    zoom: 2.5; /* Button icons are enlarged on large screens. */
  }
}

.el-row {
  width: 100%;
  margin-top: 1rem;
}

.extra-padding {
  padding: 0 0.5rem;
}

.el-row .el-col {
  padding: 0.5rem;
  text-align: center;
}

.el-card i {
  font-size: 3rem;
  margin-top: 1rem;
}

.el-col a {
  text-decoration: none;
}

.el-card.disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

li {
  margin-bottom: 0.25rem;
}
</style>
