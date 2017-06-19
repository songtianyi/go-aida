<template>
  <div id="">
    <h1>{{msg}}</h1>
    <x-table>
      <tr v-for="(plugin, index) in plugins" class="tr" :class="{'current': currentIndex === index}">
        <td>{{'No.'+index}}</td>
        <td>{{plugin.name}}</td>
        <td>{{plugin.age}}</td>
      </tr>
    </x-table>
  </div>
</template>

<script>
import { XTable } from 'vux'
export default {
  components: {
    XTable
  },
  name: 'manage',
  data () {
    return {
      msg: 'Manage',
      plugins: [
        {name: '1', age: '2'},
        {name: '2', age: '4'}
      ],
      currentIndex: 1
    }
  },
  mounted () {
    console.log('管理页')
  }
}
</script>

<style>

</style>
