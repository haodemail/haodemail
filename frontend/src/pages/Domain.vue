<template>
  <div>
    <br>
    <Modal v-model="showAddForm" :loading="loading" :title="titleNew" @on-ok="submitAdd">
      <Form :label-width="120" ref='addForm' :model='addForm' :rules="ruleValidateAddForm">
        <Row>
          <Col :span="24">
            <FormItem label="Domain name" prop="domain">
              <Input v-model="addForm.domain" :disabled="modifyDomain"></Input>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="24">
            <FormItem label="Postmaster password" :label-width="150" prop="password">
              <Input v-model="addForm.password" style="width: 150px"></Input>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="12">
            <FormItem label="Max users" prop="maxUserCount">
              <InputNumber :max="1000" :min="1" v-model="addForm.maxUserCount" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
          <Col :span="12">
            <FormItem label="Max quota(M)" prop="maxUserQuota">
              <InputNumber :max="10000" :min="-1" :step="100" v-model="addForm.maxUserQuota" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="12">
            <FormItem label="Max mails" prop="maxMailCount">
              <InputNumber :max="10000" :min="-1" :step="100" v-model="addForm.maxMailCount" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
          <Col :span="12">
            <FormItem label="Expire time" prop="expireTime">
              <DatePicker v-model="addForm.expireTime" type="date" :options="greatNow" placeholder=""></DatePicker>
            </FormItem>
          </Col>
        </Row>
      </Form>
    </Modal>
    <Modal title="Are you sure of deleting domain" v-model="showDelete" :closable="false" @on-ok="deleteDomain">
      <H3 style="color: #ed4014;">All user and mail in this domain will be REMOVED!!</H3>
    </Modal>
    <Modal v-model="showAddUserForm" :loading="loading" :title="titleNewUser" @on-ok="submitAddUser">
      <Form :label-width="120" ref='addUserForm' :model='addUserForm' :rules="ruleValidateAddUserForm">
        <Row>
          <Col :span="24">
            <FormItem label="User name" prop="userName">
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
            <FormItem label="Max quota(M)" prop="maxQuota">
              <InputNumber :max="1000" :min="-1" :step="100" v-model="addUserForm.maxQuota" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
          <Col :span="12">
            <FormItem label="Max mails" prop="maxMail">
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
            <FormItem label="Expire time" prop="expireTime">
              <DatePicker v-model="addUserForm.expireTime" type="date" :options="greatNow" placeholder=""></DatePicker>
            </FormItem>
          </Col>
        </Row>
      </Form>
    </Modal>
    <Form align="left" :label-width="80" ref='searchForm' :model='searchForm' inline>
      <FormItem :label-width="0">
        <Input v-model="searchForm.domain" search enter-button="Search" @on-search="getDomainList" placeholder="input domain name"
               style="width: 380px"/>
      </FormItem>
      <Button shape="circle" type="success" @click="prepareAdd"><Icon type="md-add" style="font-size:16px"/></Button>
    </Form>
    <div v-if="allData.length<=domainHiddenLimit" style="display: flex; flex-direction: row;flex-wrap: wrap;">
      <div class="domainFormat" v-for="domain, idx in allData.concat(emptyDomain)">
        <Card :bordered="false" style="height: 200px">
          <div v-if="!domain.creating">
            <p slot="title">{{domain.domain}}<span v-show="domain.expired" style="color:#ed4014">(expired)</span></p>
            <p>Users {{domain.userCount}}/<span style="color:#19be6b" v-if="domain.maxUserCount==-1">unlimited</span><span v-else>{{domain.maxUserCount}}</span>
            </p>
            <p>Quota {{domain.userQuota}}/<span style="color:#19be6b" v-if="domain.maxUserQuota==-1">unlimited</span><span v-else>{{domain.maxUserQuota}}</span>
            </p>
            <p>Mails {{domain.mailCount}}/<span style="color:#19be6b" v-if="domain.maxMailCount==-1">unlimited</span><span v-else>{{domain.maxMailCount}}</span>
            </p>
            <p>Create time {{domain.createTime | moment("YYYY-MM-DD HH:mm")}}</p>
            <p>Expire time {{domain.expireTime | moment("YYYY-MM-DD HH:mm")}}</p>
            <p>
              <Icon type="md-person-add" style="font-size:20px;color:#19be6b;cursor:pointer; margin-right: 10px"
                    @click="addUserForm.domain=domain.domain;addUserForm.domainID=domain.id;showAddUserForm=true"/>
              <Icon type="md-grid" style="font-size:20px;color:#5cadff;cursor:pointer; margin-right: 10px"
                    @click="$router.push({name: 'Account', params: {domain:domain.domain,ID: domain.id}})"/>
              <Icon type="md-create" style="font-size:20px;color:#559DF9;cursor:pointer; margin-right: 10px" @click="prepareModify(domain)"/>
              <Icon type="md-trash" style="font-size:20px;color:#ed4014;cursor:pointer; margin-right: 10px"
                    @click="deletingID=domain.id;showDelete=true"/>
            </p>
          </div>
          <div v-else style="line-height: 160px;">
            <Icon type="md-add" style="font-size:72px; color: brown; cursor: pointer" @click="prepareAdd"></Icon>
          </div>
        </Card>
      </div>
    </div>
    <Table v-else ref="refDomainsTable" stripe :columns="tableColumns" :data="tableData" @on-sort-change="sortTableData"></Table>
    <Page v-if="allData.length>domainHiddenLimit" :total="allData.length" :page-size="pageSize" :current="page" show-sizer show-total
          @on-change="changePage"
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
        showDelete: false,
        deletingID: "",
        domainHiddenLimit: 2,
        showAddForm: false,
        modifyDomain: false,
        showAddUserForm: false,
        modifyUser: false,
        loading: true,
        pageSize: 15,
        page: 1,
        allData: [],
        tableData: [],
        emptyDomain: {
          domain: "New domain",
          creating: true,
        },
        searchForm: {
          domain: "",
          orderBy: "createTime",
          orderSort: "desc",
        },
        titleNew: "Create a new domain",
        titleNewUser: "Create a new user",
        addForm: {
          id: "",
          domain: "",
          password: "",
          maxUserCount: 10,
          maxUserQuota: 1000,
          maxMailCount: 1000,
          expireTime: new Date(2019, 12, 30),
        },
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
        ruleValidateAddForm: {
          domain: [
            {
              required: true,
              type: "string",
              message: "Please input domain",
              trigger: "blur"
            }
          ],
          password: [
            {
              validator: validPassword,
              trigger: "blur"
            }
          ],
          maxUserCount: [
            {
              required: true,
              type: "number",
              message: "Please input max users",
              trigger: "blur"
            }
          ],
          maxUserQuota: [
            {
              required: true,
              type: "number",
              message: "Please input max user quota",
              trigger: "blur"
            }
          ],
          maxMailCount: [
            {
              required: true,
              type: "number",
              message: "Please input max user quota",
              trigger: "blur"
            }
          ],
        },
        ruleValidateAddUserForm: {
          userName: [
            {
              required: true,
              type: "string",
              minLength: 3,
              maxlength: 16,
              message: "Please input username, 3-16 characters",
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
            title: 'Domain name',
            key: 'domain',
            sortable: "custom",
            render: (h, params) => {
              if (params.row.expired) {
                return h("span", {
                  style: {
                    color: '#ed4014',
                  }
                }, params.row.domain + "(expired)")
              } else {
                return h("span", params.row.domain)
              }
            }
          }, {
            title: 'Users',
            key: 'userCount',
            width: 120,
            sortable: "custom",
            render: (h, params) => {
              let c = params.row.maxUserCount
              if (c < 0) {
                c = "unlimited"
              }
              return h("span", params.row.userCount + "/" + c)
            }
          }, {
            title: 'Quota',
            key: 'userQuota',
            width: 120,
            sortable: "custom",
            render: (h, params) => {
              let c = params.row.maxUserQuota
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
            sortable: "custom",
            render: (h, params) => {
              return h("span", moment(params.row.createTime).format("YYYY-MM-DD HH:mm"))
            }
          }, {
            title: 'Expire time',
            key: 'expireTime',
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
              return h('div', [
                h('Icon', {
                  props: {
                    type: 'md-person-add'
                  },
                  style: {
                    fontSize: '20px',
                    color: '#19be6b',
                    cursor: "pointer",
                  },
                  on: {
                    click: () => {
                      let that = this
                      that.addUserForm.domain=params.row.domain
                      that.addUserForm.domainID=params.row.id
                      that.showAddUserForm=true
                    }
                  }
                }),
                h('Icon', {
                  props: {
                    type: 'md-grid'
                  },
                  style: {
                    fontSize: '20px',
                    color: '#5cadff',
                    cursor: "pointer",
                  },
                  on: {
                    click: () => {
                      let that = this
                      console.log(params.row)
                      that.$router.push({
                        name: 'Account', params: {
                          domain:params.row.domain,
                          ID: params.row.id
                        }
                      })
                    }
                  }
                }),
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
      deleteDomain() {
        let that = this
        if (that.deletingID.length == 0) {
          return
        }
        let para = {
          id: that.deletingID
        }
        APIManger.deleteDomain(para).then((res) => {
          if (res.data.ok == true) {
            that.$Message.success(res.data.info)
            that.getDomainList()
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
        that.showAddForm = true
        that.modifyDomain = false
        that.titleNew = 'Create A New Domain'
        that.addForm.id = ""
        that.addForm.domain = "test126.com"
        that.addForm.password = "123abcD$"
        that.addForm.maxUserCount = -1
        that.addForm.maxUserQuota = -1
        that.addForm.maxMailCount = -1
        that.addForm.expireTime = ""
      },
      prepareModify(d) {
        let that = this
        that.showAddForm = true
        that.modifyDomain = true
        that.titleNew = 'Modify domain property'
        that.addForm.id = d.id
        that.addForm.domain = d.domain
        that.addForm.maxUserCount = d.maxUserCount
        that.addForm.maxUserQuota = d.maxUserQuota
        that.addForm.maxMailCount = d.maxMailCount
        if (d.expireTime == "0001-01-01T00:00:00Z") {
          that.addForm.expireTime = ""
        } else {
          that.addForm.expireTime = d.expireTime
        }
      },
      submitAdd() {
        let that = this
        that.loading = false;
        that.$refs["addForm"].validate(valid => {
          if (valid) {
            APIManger.createDomain(that.addForm).then((res) => {
              if (res.data.ok == true) {
                that.showAddForm = false;
                that.$Message.success(res.data.info)
                that.getDomainList()
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
      getDomainList() {
        let that = this
        let para = {
          domain: this.searchForm.domain,
          orderBy: this.searchForm.orderBy + " " + this.searchForm.orderSort
        }
        APIManger.listDomain(para).then((res) => {
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
      getUserList() {
        let that = this
        let para = {
          domain: this.searchForm.domain,
          orderBy: this.searchForm.orderBy + " " + this.searchForm.orderSort
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
      //   let h = window.innerHeight - that.$refs.refDomainsTable.$el.offsetTop - 80
      //   that.pageSize = Math.floor(h / 40)
      //   that.tableData = that.allData.slice(0, that.pageSize);
      // }
      that.getDomainList();
    }
    ,
  }
</script>
<style scoped>
  .domainFormat {
    /*border: 1px solid #c0ccda;*/
    /*border-radius: 10px;*/
    flex: auto;
    margin-bottom: 20px;
    margin-right: 20px;
    width: 50vh;
    word-break: break-all;
    word-wrap: break-word;
  }
</style>
