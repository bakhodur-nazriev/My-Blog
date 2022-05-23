<template>
  <div class="home">
    <div class="home-banner">
      <img :src="bannerOne">
    </div>
    <h1>This is Home page</h1>
  </div>
</template>

<script>
import firstBanner from '../assets/gallery-post-03-1230x615.jpg';

export default {
  name: "HomeView",
  components: {},
  data() {
    return {
      bannerOne: firstBanner
    }
  }
};
</script>
