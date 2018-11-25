<template>
  <Row class="login vm-panel">
   <Col span="12" class="login-form">
     <div class="login-header">
       <img src="../assets/img/logo.png" height="80" alt="">
       <p>Welcome To Partner of Postfix</p>
     </div>
     <div class="login-form">
        <Input v-model="username" placeholder="please enter username"></Input>
        <Input v-model="password" type="password" placeholder="Please enter password"></Input>
        <Button type="primary">Login</Button>
     </div>
     <div class="login-footer">
       <Checkbox v-model="remenber">Remenber</Checkbox>
       <span class="forget"><a href="#">Forget Password</a></span>
     </div>
   </Col>
  </Row>
</template>
<script>
  export default {
    name: 'Login',
    data: function () {
      return {
        username: '',
        password: '',
        remenber: false
      }
    }
  }
</script>
