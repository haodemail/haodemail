<template>
  <div>
  <Form ref="loginForm" :model="loginForm" :rules="rules" @keydown.enter.native="handleSubmit">
    <Row class="login vm-panel">
      <Col span="6" offset="9" class="login-form">
        <div class="login-header">
          <img src="../assets/img/postfix.jpg" height="200" alt="">
          <h1><br>Welcome To Partner of Postfix</h1>
        </div>
        <div class="login-form">
          <Alert type="error" v-show="errorMsg.length>0">{{errorMsg}}</Alert>
          <form-item prop="username">
            <Input prefix="ios-contact" v-model="loginForm.username" placeholder="please enter username"></Input>
          </form-item>
          <form-item prop="password">
            <Input prefix="ios-lock" v-model="loginForm.password" type="password"
                   placeholder="Please enter password"></Input>
          </form-item>
          <Button type="primary" @click="handleSubmit">Login</Button>
        </div>
        <div class="login-footer">
          <span class="forget"><a href="#">Forget Password</a></span>
        </div>
      </Col>
    </Row>
  </Form>
  <Footer class="layout-footer-center footer">2013-2018 &copy; HaodeMail.Com</Footer>
  </div>
</template>
<script>
  import APIManger from "../api/"

  export default {
    data: function () {
      return {
        errorMsg: "",
        loginForm: {
          username: 'postmaster@haodemail.com',
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
        that.errorMsg = ""
        that.$refs["loginForm"].validate(valid => {
          if (valid) {
            APIManger.login(that.loginForm).then(resp => {
              let info = resp.data.info;
              let ok = resp.data.ok;
              let data = resp.data.data;
              if (!ok) {
                that.errorMsg = info
              } else {
                let token = data.token;
                sessionStorage.setItem("JWT_TOKEN", token);
                sessionStorage.setItem("username", that.loginForm.username);
                that.$router.push({path: "/"});
              }
            }).catch(function (err) {
              that.errorMsg=  err.toString()
            });
          } else {
            that.errorMsg = "invalid form"
          }
        });
      }
    }
  }
</script>

<style>
  .login-header {
    margin-top: 35px;
  }
  .login-form {
    margin-top: 30px;
  }

  .login-footer {
    margin-top: 20px;
  }

  .footer {
    position: absolute;
    bottom: 0px;
    width: 100%;
  }
</style>
