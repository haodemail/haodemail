<template>
  <Form ref="loginForm" :model="loginForm" :rules="rules" @keydown.enter.native="handleSubmit">
    <Row class="login vm-panel">
      <Col span="8" offset="8" class="login-form">
        <div class="login-header">
          <img src="../assets/img/logo.png" height="80" alt="">
          <h1>Welcome To Partner of Postfix</h1>
        </div>
        <div class="login-form">
          <form-item prop="username">
            <Input v-model="loginForm.username" placeholder="please enter username"></Input>
          </form-item>
          <form-item prop="password">
            <Input v-model="loginForm.password" type="password" placeholder="Please enter password"></Input>
          </form-item>
          <Button type="primary" @click="handleSubmit">Login</Button>
        </div>
        <div class="login-footer">
          <span class="forget"><a href="#">Forget Password</a></span>
        </div>
      </Col>
    </Row>
  </Form>
</template>
<script>
  import APIManger from "../api/"

  export default {
    data: function () {
      return {
        loginForm: {
          username: 'admin@haodemail.com',
          password: 'postfix123'
        },
        rules: {
          username: [
            {required: true, type: "string", message: "Please input username", trigger: "blur"}
          ],
          password: [
            {required: true, type: "string", message: "Please input password", trigger: "blur"}
          ],
        }
      }
    },
    methods: {
      handleSubmit() {
        let that = this
        that.$refs["loginForm"].validate(valid => {
          if (valid) {
            APIManger.login(that.loginForm).then(resp => {
              let info = resp.data.info;
              let ok = resp.data.ok;
              let data = resp.data.data;
              let token = data.token;
              if (!ok) {
                that.$Message.error({
                  content: info,
                  duration: 5
                });
              } else {
                sessionStorage.setItem("JWT_TOKEN", token);
                sessionStorage.setItem("username", that.loginForm.username);
                that.$router.push({path: "/"});
              }
            }).catch(function (err) {
              that.$Message.error({
                content: err.toString(),
                duration: 5
              });
            });
          } else {
            that.$Message.error("invalid form!");
          }
        });
      }
    }
  }
</script>

<style>
  .login-form {
    margin-top: 50px;
  }

  .login-footer {
    margin-top: 20px;
  }
</style>
