<template>
    <div class="modal-backdrop">
      <div class="modal">
        <header class="modal-header">
          <slot name="header">
            <botton class="btn-close">X</botton>
          </slot>
        </header>
        <section class="modal-body">
          <slot name="body">
            container
          </slot>
        </section>
       <footer calss="modal-footer">
         <button class="btn-green" @click="close">cancel</button>
         <button @click="save">submit</button>
       </footer>

      </div>
    </div>
</template>

<script>
    export default {
        name: "dialog_modal",
      methods:{
          close(){
            this.$emit('close')
          },
        save(){
            this.$emit('save')
           }


      }
    }
</script>

<style scoped>
.modal-backdrop{
  position: fixed;
  top: 0px;
  bottom: 0px;
  left: 0px;
  right: 0px;
  background-color: rgba(0,0,0,0.3);
  display: flex;
  justify-content: center;
  align-items: center;
  width: 500px;
  height: 500px;
}
  .modal{
    background:#ffffff;
    box-shadow: 2px 2px 20px 1px;
    overflow-x: auto;
    display: flex;
    flex-direction: column;
  }
  .modal-header,
  .modal-footer{
    padding: 15px;
    display: flex;
  }
.modal-header{
  border-bottom: 1px solid #eeeeee;
  color: #4aae9b;
  justify-content: space-between;
}
  .modal-footer{
    border-top: 1px solid #eeeeee;
    justify-content: flex-end;
  }
  .modal-body{
    position: relative;
    padding:20px 10px;
  }
  .btn-close{
    border:none;
    font-size: 20px;
    padding: 20px;
    cursor: pointer;
    font-weight: bold;
    color: #4aae9b;
    background: transparent;
  }
  .btn-green{
    color:white;
    background: #4aae9b;
    border: 1px solid #4aae9b;
    border-radius: 2px;
  }
</style>
