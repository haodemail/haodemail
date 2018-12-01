<template>
  <div class="home">
    <Row class="header" type="flex" theme="dark" align="middle">
      <div class="logo">
        <img src="../assets/img/logo.png" height="30" alt="">
        <span>Setup Postfix Mail System All In One</span>
      </div>
    </Row>   
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
                <i class="fa fa-database"></i>
                Domains
              </router-link>
          </Menu-item>
          <Menu-item name="Users">
              <router-link to="/user">
                <i class="fa fa-user"></i>
                Users
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
  components: {
  },
  mounted: function () {
    this.activeName = this.$route.name
  },
  data () {
    return {
      activeName: 'Dashboard',
    }
  }
}
</script>

<style lang="less">

</style>
