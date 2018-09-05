<template>
  <div>
    {{hello}}
    {{number}}
    <button @click="emitMyEvent">emit</button>
    <slot name="header"> no header</slot>
    <span>sdfsdfsdfsd</span>
<slot name="footer"> no footer</slot>
</div>


</template>

<script>
  export default {
    props:{
      'number':[String,Number]
    },
    data(){
      return {
        hello:"i am component A",
      }
    },

    methods:{
      emitMyEvent(){
        console.log("comp-a click")
        this.$emit("my-event",this)
      }
    }
  }
</script>

<style scoped>

</style>
