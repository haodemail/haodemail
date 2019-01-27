<template>
  <div>
    <br>
    <Modal title="Are you sure of deleting user" v-model="showDelete" :closable="false" @on-ok="deleteUser">
      <H3 style="color: #ed4014;">All mails will be REMOVED!!</H3>
    </Modal>
    <Modal v-model="showAddUserForm" :loading="loading" :title="titleNewUser" @on-ok="submitAddUser">
      <Form :label-width="120" ref='addUserForm' :model='addUserForm' :rules="ruleValidateAddUserForm">
        <Row>
          <Col :span="24">
            <FormItem label="User Name" prop="userName">
              <Input v-model="addUserForm.userName" :disabled="modifyUser">
                <span slot="append">@{{addUserForm.domain}}</span>
              </Input>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="24">
            <FormItem label="Password" :label-width="120" prop="password">
              <Input v-model="addUserForm.password" style="width: 150px"></Input>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="12">
            <FormItem label="Max Quota(M)" prop="maxQuota">
              <InputNumber :max="1000" :min="-1" :step="100" v-model="addUserForm.maxQuota" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
          <Col :span="12">
            <FormItem label="Max Mails" prop="maxMail">
              <InputNumber :max="10000" :min="-1" :step="100" v-model="addUserForm.maxMail" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="12">
            <FormItem label="NickName" :label-width="120" prop="nickName">
              <Input v-model="addUserForm.nickName"></Input>
            </FormItem>
          </Col>
          <Col :span="12">
            <FormItem label="Expire Time" prop="expireTime">
              <DatePicker v-model="addUserForm.expireTime" type="date" :options="greatNow" placeholder=""></DatePicker>
            </FormItem>
          </Col>
        </Row>
      </Form>
    </Modal>
    <Form align="left" :label-width="80" ref='searchForm' :model='searchForm' inline>
      <Button shape="circle" type="default" @click="$router.push({'name':'Domain'})">
        <Icon type="md-arrow-round-back" style="font-size:20px"/>
      </Button>
      <FormItem :label-width="0">
        <Input v-model="searchForm.userName" search enter-button="Search" @on-search="getUserList" placeholder="input user name"
               style="width: 380px"/>
      </FormItem>
      <Button shape="circle" type="success" @click="prepareAdd">
        <Icon type="md-person-add" style="font-size:20px"/>
      </Button>
    </Form>
    <Table ref="refUsersTable" stripe :columns="tableColumns" :data="tableData" @on-sort-change="sortTableData"></Table>
    <Page :total="allData.length" :page-size="pageSize" :current="page" show-sizer show-total @on-change="changePage"
          style="margin-top: 10px; float:right"/>
  </div>
</template>

<script>
  import APIManger from "../api/"
  import {compare} from "../api/compare"

  let moment = require('moment');

  export default {
    name: 'Domain',
    components: {},
    data() {
      let validPassword = (rule, value, callback) => {
        let that = this
        if (that.modifyDomain == false && that.modifyUser == false) {
          if (value.length < 8 || value.length > 16) {
            return callback(new Error("password must be 8-16 characters"))
          }
          let score = 0
          if (/\d/.test(value)) score++;
          if (/[a-z]/.test(value)) score++;
          if (/[A-Z]/.test(value)) score++;
          if (/\W/.test(value)) score++;
          if (score < 3) {
            return callback(new Error("Must contain uppercase, lowercase, and number"))
          }
        }
        callback()
      };
      return {
        domainName: this.$route.params.domain,
        domainID: this.$route.params.ID,
        showDelete: false,
        deletingID: "",
        showAddUserForm: false,
        modifyUser: false,
        loading: true,
        pageSize: 15,
        page: 1,
        allData: [],
        tableData: [],
        searchForm: {
          userName: "",
          orderBy: "createTime",
          orderSort: "desc",
        },
        titleNewUser: "Create a new user",
        addUserForm: {
          id: "",
          domainID: "",
          domain: "",
          userName: "",
          nickName: "",
          password: "",
          maxQuota: 1000,
          maxMail: 10000,
          expireTime: new Date(2019, 12, 30),
        },
        greatNow: {
          disabledDate(date) {
            return date && date.valueOf() < Date.now() + 86400000;
          }
        },
        ruleValidateAddUserForm: {
          userName: [
            {
              required: true,
              type: "string",
              minLength: 3,
              maxlength: 16,
              message: "Please input userName, 3-16 characters",
              trigger: "blur"
            }
          ],
          password: [
            {
              validator: validPassword,
              trigger: "blur"
            }
          ],
          maxQuota: [
            {
              required: true,
              type: "number",
              message: "Please input max quota",
              trigger: "blur"
            }
          ],
          maxMail: [
            {
              required: true,
              type: "number",
              message: "Please input max mail count",
              trigger: "blur"
            }
          ],
        },
        tableColumns: [
          {
            title: 'userName',
            key: 'userName',
            sortable: "custom",
            render: (h, params) => {
              if (params.row.expired) {
                return h("span", {
                  style: {
                    color: '#ed4014',
                  }
                }, params.row.userName + "@" + params.row.domain + "(expired)")
              } else {
                return h("span", params.row.userName + "@" + params.row.domain)
              }
            }
          }, {
            title: 'Quota',
            key: 'userQuota',
            width: 120,
            sortable: "custom",
            render: (h, params) => {
              let c = params.row.maxQuota
              if (c < 0) {
                c = "unlimited"
              }
              return h("span", params.row.userQuota + "/" + c)
            }
          }, {
            title: 'Mails',
            key: 'mailCount',
            width: 120,
            sortable: "custom",
            render: (h, params) => {
              let c = params.row.maxMailCount
              if (c < 0) {
                c = "unlimited"
              }
              return h("span", params.row.mailCount + "/" + c)
            }
          }, {
            title: 'Create time',
            key: 'createTime',
            width: 150,
            sortable: "custom",
            render: (h, params) => {
              return h("span", moment(params.row.createTime).format("YYYY-MM-DD HH:mm"))
            }
          }, {
            title: 'Expire time',
            key: 'expireTime',
            width: 150,
            sortable: "custom",
            render: (h, params) => {
              if (params.row.expireTime == "0001-01-01T00:00:00Z") {
                return h("span", "never expire")
              } else {
                return h("span", moment(params.row.expireTime).format("YYYY-MM-DD HH:mm"))
              }
            }
          }, {
            title: 'Operation',
            key: '',
            width: 120,
            align: 'center',
            render: (h, params) => {
              if (params.row.userName == "postmaster") {
                return
              }
              return h('div', [
                h('Icon', {
                  props: {
                    type: 'md-create'
                  },
                  style: {
                    fontSize: '20px',
                    color: '#559DF9',
                    cursor: "pointer",
                  },
                  on: {
                    click: () => {
                      let that = this
                      that.prepareModify(params.row)
                    }
                  }
                }),
                h('Icon', {
                  props: {
                    type: 'md-trash'
                  },
                  style: {
                    fontSize: '20px',
                    color: '#ed4014',
                    cursor: "pointer",
                  },
                  on: {
                    click: () => {
                      let that = this
                      that.showDelete = true
                      that.deletingID = params.row.id
                    }
                  }
                }),
              ])
            },
          }],
      };
    },
    methods: {
      sortData(k, desc) {
        this.allData.sort(compare(k, desc))
        this.changePage(this.page)
      },
      sortTableData(c) {
        this.sortData(c.key, c.order == "desc")
      },
      changePage(index) {
        let _start = (index - 1) * this.pageSize;
        let _end = index * this.pageSize;
        this.tableData = this.allData.slice(_start, _end);
      },
      deleteUser() {
        let that = this
        if (that.deletingID.length == 0) {
          return
        }
        let para = {
          id: that.deletingID
        }
        APIManger.deleteUser(para).then((res) => {
          if (res.data.ok == true) {
            that.$Message.success(res.data.info)
            that.getUserList()
          } else {
            that.$Message.info(res.data.info)
          }
        }).catch(function (err) {
          that.$Message.error({
            content: err.message,
            duration: 5
          });
        });
      },
      prepareAdd() {
        let that = this
        that.showAddUserForm = true
        that.modifyUser = false
        that.titleNewUser = 'Create a new user'
        that.addUserForm.id = ""
        that.addUserForm.domain = that.domainName
        that.addUserForm.domainID = that.domainID
        that.addUserForm.password = ""
        that.addUserForm.maxQuota = -1
        that.addUserForm.maxMailCount = -1
        that.addUserForm.expireTime = ""
      },
      prepareModify(u) {
        let that = this
        that.showAddUserForm = true
        that.modifyUser = true
        that.titleNewUser = 'Modify user property'
        that.addUserForm.id = u.id
        that.addUserForm.userName = u.userName
        that.addUserForm.nickName = u.nickName
        that.addUserForm.domain = u.domain
        that.addUserForm.domainID = u.domainID
        that.addUserForm.maxQuota = u.maxQuota
        that.addUserForm.maxMailCount = u.maxMailCount
        if (u.expireTime == "0001-01-01T00:00:00Z") {
          that.addUserForm.expireTime = ""
        } else {
          that.addUserForm.expireTime = u.expireTime
        }
      },
      submitAddUser() {
        let that = this
        that.loading = false;
        that.$refs["addUserForm"].validate(valid => {
          if (valid) {
            APIManger.createUser(that.addUserForm).then((res) => {
              if (res.data.ok == true) {
                that.showAddUserForm = false;
                that.$Message.success(res.data.info)
                that.getUserList()
              } else {
                that.$Message.error({
                  content: res.data.info,
                  duration: 5
                })
              }
            }).catch(function (err) {
              that.$Message.error({
                content: err.message,
                duration: 5
              });
            });
          } else {
            setTimeout(() => {
              that.loading = false;
            }, 0);
          }
        });
      },
      getUserList() {
        let that = this
        let para = {
          domainID: that.$route.params.ID,
          userName: that.searchForm.userName,
          orderBy: that.searchForm.orderBy + " " + that.searchForm.orderSort
        }
        APIManger.listUser(para).then((res) => {
          if (res.data.ok == true) {
            if (res.data.data) {
              that.allData = res.data.data
              that.tableData = that.allData.slice(0, that.pageSize);
              that.page = 1
            } else {
              that.allData = []
              that.tableData = []
              that.page = 1
            }
          }
        })
      },
    },
    mounted() {
      let that = this
      // if (that.allData.length > that.domainHiddenLimit) {
      //   let h = window.innerHeight - that.$refs.refUsersTable.$el.offsetTop - 80
      //   that.pageSize = Math.floor(h / 40)
      //   that.tableData = that.allData.slice(0, that.pageSize);
      // }
      that.getUserList();
    },
  }
</script>
<style scoped>
</style>

