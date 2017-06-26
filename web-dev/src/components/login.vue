<template>
  <div id="">
    <h1>{{msg}}</h1>
    <img :src="QRSrc" alt="" class="ximg-demo">
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
      msg: '扫码登陆',
      QRSrc: ''
    }
  },
  methods: {
    getCode () {
      this.$ajax({
        method: 'get',
        url: 'http://localhost:8080/create'
      }).then(res => {
        let data = res.data
        console.log(res)
        localStorage['uuid'] = data
        let timeLoop = setInterval(() => {
          this.getQR(data, timeLoop)
        }, 500)
      }, res => {
        console.log('error', res)
      })
    },
    getQR (code, timeLoop) {
      this.$ajax({
        method: 'get',
        url: 'http://localhost:8080/status',
        params: {
          uuid: code
        }
      }).then(res => {
        console.log(res)
        if (res.data.status === 'SERVING') {
          clearInterval(timeLoop)
          this.$router.push('/manage')
          return
        }
        this.QRSrc = res.data.qrcode.replace('/web', '')
      }, res => {
        this.getCode()
      })
    }
  },
  created () {
    console.log('created')
    this.getCode()
  },
  updated () {
    console.log('updated')
  }
}
</script>

<style>

</style>
