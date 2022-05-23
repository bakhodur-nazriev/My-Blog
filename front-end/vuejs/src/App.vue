<template>
  <navbar-component></navbar-component>
  <router-view/>
  <footer-component></footer-component>
</template>

<script>
import NavbarComponent from "@/components/NavbarComponent";
import FooterComponent from "@/components/FooterComponent";

export default {
  components: {
    NavbarComponent,
    FooterComponent
  }
}
</script>

<style lang="scss">
body {
  margin: 0;
}

#app {
  font-family: "Red Hat Display", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
