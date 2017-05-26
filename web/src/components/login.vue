<template>
  <div id="">
    <h1>{{msg}}</h1>
    <x-img :src="QRSrc" class="ximg-demo"></x-img>

  </div>
</template>

<script>
import { XImg } from 'vux'
export default {
  components: {
    XImg
  },
  name: '',
  data () {
    return {
      msg: 'Login',
      QRSrc: 'https://o5omsejde.qnssl.com/demo/test1.jpg'
    }
  },
  mounted () {
    console.log('登陆页')
//    this.$http.get().then(res => {
//
//    }, res => {
//
//    })
  }
}
</script>

<style>

</style>
