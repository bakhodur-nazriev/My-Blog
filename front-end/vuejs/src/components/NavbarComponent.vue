<template>
  <nav>
    <router-link class="test" to="/">Home</router-link>
    <router-link to="/about">About</router-link>
    <router-link to="/posts">Posts</router-link>
    <router-link to="/authors">Authors</router-link>
    <router-link to="/contact">Contact</router-link>
  </nav>
</template>

<script>
export default {
  data() {
    return {}
  }
}
</script>

<style>
nav {
  padding: 30px;
}

nav a {
  padding: 0 20px;
  font-weight: bold;
  color: #878787;
  text-decoration: none;
}

nav a.router-link-exact-active {
  color: #3858f6;
}
</style>