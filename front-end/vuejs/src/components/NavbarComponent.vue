<template>
  <nav class="main-navbar">
    <div class="main-logo">
      <img alt="logo" :src="mainBlackLogo">
    </div>

    <div>
      <router-link class="test" to="/">Home</router-link>
      <router-link to="/about">About</router-link>
      <router-link to="/posts">Posts</router-link>
      <router-link to="/authors">Authors</router-link>
      <router-link to="/contact">Contact</router-link>
    </div>

    <div>
      <input type="text" class="form-control" placeholder="Search...">
      <a href="" class="shopping-card">a</a>
    </div>
  </nav>
</template>

<script>
import blackLogo from '../assets/black_logo.png';

export default {
  data() {
    return {
      mainBlackLogo: blackLogo
    }
  }
}
</script>

<style>
.shopping-card {
  border-radius: 50%;
  background-color: #f0f2f5;
  padding: 10px;
}

.form-control {
  border: none;
  background-color: #f0f2f5;
  min-width: 143px;
  height: 35px;
  border-radius: 44px;

}

nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 23px 63px;
  box-shadow: rgba(100, 100, 111, 0.2) 0 7px 29px 0;
}

nav a {
  color: #878787;
  text-decoration: none;
  margin: 0 18px;
  font-weight: 500;
}

nav a.router-link-exact-active {
  color: #3858f6;
}

.main-logo {
  display: flex;
  align-items: center;
}

.main-logo > img {
  max-height: 37px;
}
</style>