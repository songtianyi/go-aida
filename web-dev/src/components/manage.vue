<template>
  <div>
    <h1>{{msg}}</h1>
    <div>
      <group v-for="(plugin, index) in plugins" :key="index">
        <x-switch :title="plugin.name" v-model="plugin.enabled" @on-change="changeEnable(plugin.name, plugin.enabled, index)"></x-switch>
      </group>
    </div>
  </div>
</template>

<script>
import { XTable, XSwitch, Group } from 'vux'
export default {
  components: {
    XTable,
    XSwitch,
    Group
  },
  name: 'manage',
  data () {
    return {
      msg: '插件管理',
      plugins: [],
      currentIndex: 1,
      uuid: localStorage.uuid || null
    }
  },
  methods: {
    changeEnable (name, v, index) {
      if (v) {
        console.log(v)
        this.$ajax({
          method: 'PATCH',
          url: 'http://localhost:8080/enable',
          params: {
            uuid: this.uuid,
            name: name
          }
        }).then(res => {
          console.log(res)
        }, res => {
          console.log('enable error', res)
        })
      } else {
        this.$ajax({
          method: 'PATCH',
          url: 'http://localhost:8080/disable',
          params: {
            uuid: this.uuid,
            name: name
          }
        }).then(res => {
          console.log(res)
        }, res => {
          console.log('enable error', res)
        })
      }
    },
    getPlugins () {
      this.$ajax({
        method: 'get',
        url: 'http://localhost:8080/status',
        params: {
          uuid: this.uuid
        }
      }).then(res => {
        this.plugins = res.data.plugins
      }, res => {
        this.$router.push('/')
        console.log('get error', res)
      })
    }
  },
  mounted () {
    this.getPlugins()
  }
}
</script>

<style>

</style>
