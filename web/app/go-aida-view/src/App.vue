<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
  mounted: function () {
    let deviceWidth = document.documentElement.clientWidth
    if (deviceWidth > 750) deviceWidth = 750
    let dpr = window.devicePixelRatio || 1
    document.documentElement.setAttribute('data-dpr', dpr)
    document.documentElement.style.fontSize = deviceWidth / 7.5 + 'px'
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
