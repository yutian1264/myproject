<template>
  <div>
    <div class="dialog-wrap">
      <div class="dialog-cover"  v-if="isShow" @click="closeMyself"></div>
      <transition name="drop">
        <div class="dialog-content"  :style="contentCss" v-if="isShow">
          <div class="header-model">
            <slot name="title"></slot>
            <p class="dialog-close" @click="closeMyself">x</p>
          </div>
          <div class="content-body">
            <slot></slot>
          </div>

        </div>
      </transition>
    </div>
  </div>
</template>


  <script>
  export default {
    props: {
      isShow: {
        type: Boolean,
        default: true
      },
      contentCss:{
        type:String,
        default:""
      }
    },
    data () {
      return {

      }
    },
    methods: {
      closeMyself () {
        this.$emit('on-close')
      }
    }
  }
</script>

<style scoped>
  .drop-enter-active {
    transition: all .5s ease;
  }
  .drop-leave-active {
    transition: all .3s ease;
  }
  .drop-enter {
    transform: translateY(-500px);
  }
  .drop-leave-active {
    transform: translateY(-500px);
  }

  .dialog-wrap {
    position: fixed;
    width: 100%;
    height: 100%;
  }
  .dialog-cover {
    background: #000;
    opacity: .3;
    position: fixed;
    z-index: 5;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
  .dialog-content {
    width: 50%;
    position: fixed;
    max-height: 100%;
    overflow: auto;
    background: #fff;
    top: 20%;
    left: 50%;
    z-index: 10;
    border: 2px solid #cdcdcd;
    padding: 0;
    line-height: 1.7;
    transform: translate(-50%,-18%); /*自己的50% */
  }
  .header-model{
    justify-content: space-between;
    padding: 15px;
    font-size: 20px;
    font-weight: bold;
    border-bottom: #cdcdcd 1px solid
  }
  .dialog-close {
    position: absolute;
    right: 5px;
    top: 5px;
    width: 20px;
    height: 20px;
    text-align: center;
    cursor: pointer;
    /*background-color: #5e5e5e;*/
  }
  .dialog-close:hover {
    color: #4fc08d;
  }
</style>
