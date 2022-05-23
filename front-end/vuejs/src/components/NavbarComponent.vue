<template>
  <nav class="main-navbar">
    <div>
      <img height="70" src="../assets/logo_blog.png" alt="logo">
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
    </div>
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
  display: flex;
  justify-content: space-around;
  padding: 30px;
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
</style>