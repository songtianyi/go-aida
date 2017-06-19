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
      QRSrc: ''
    }
  },
  mounted () {
    let _self = this
    this.$http.get('http://localhost:8080/create').then(res => {
      let body = res.body
      console.log(body)
      _self.$http.get('http://localhost:8080/status?uuid=' + body).then(res => {
        _self.QRSrc = res.body.qrcode
        console.log(_self)
      }, res => {
        console.log('sub error', res)
      })
    }, res => {
      console.log('error', res)
    })
  }
}
</script>

<style>

</style>
