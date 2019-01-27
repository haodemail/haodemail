<template>
  <div class="home">
    <div class="header" type="flex" theme="dark" align="middle">
      <div class="logo">
        <img src="../assets/img/logo.png" height="30" alt="">
        <span>Make Postfix More Powerful</span>
      </div>
      <Dropdown class="logout">
        <a href="javascript:void(0)">
          <Avatar style="background-color: #87d068;color:#ffffff">
            <Icon type="md-person" style="font-size:32px"/>
          </Avatar>
        </a>
        <DropdownMenu slot="list">
          <DropdownItem>
            <a @click="logout">logout {{loginUser}}</a>
          </DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </div>
    <div class="sidebar">
      <Menu width="100%" class="menu" :active-name="activeName" :accordion="true">
        <Menu-item name="Dashboard">
          <router-link to="/">
            <i class="fa fa-dashboard"></i>
            Dashboard
          </router-link>
        </Menu-item>
        <Menu-item name="Domains">
          <router-link to="/domain">
            <i class="fa fa-user"></i>
            Domains
          </router-link>
        </Menu-item>
        <Menu-item name="Configure">
          <router-link to="/config">
            <i class="fa fa-cog"></i>
            Configure
          </router-link>
        </Menu-item>
        <Menu-item name="Queue">
          <router-link to="/queue">
            <i class="fa fa-list-alt"></i>
            Queue
          </router-link>
        </Menu-item>
        <Menu-item name="MailLog">
          <router-link to="/maillog">
            <i class="fa fa-file-text"></i>
            MailLog
          </router-link>
        </Menu-item>
      </Menu>
    </div>
    <div class="main-content">
      <router-view></router-view>
    </div>
  </div>
</template>
<script>
  export default {
    name: 'home',
    components: {},
    mounted: function () {
      this.activeName = this.$route.name
      this.loginUser = sessionStorage.getItem('username')
    },
    data() {
      return {
        activeName: 'Dashboard',
        loginUser: "",
      }
    },
    methods:{
      logout: function () {
        let that = this
        sessionStorage.removeItem('JWT_TOKEN')
        sessionStorage.removeItem('username')
        that.$router.push('/login')
      },
    }
  }
</script>

<style lang="less">

</style>
