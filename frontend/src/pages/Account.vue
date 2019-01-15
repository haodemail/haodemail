<template>
  <div>
    <br>
    <Modal v-model="showAddForm" :loading="loading" :title="titleNew" @on-ok="submitAdd">
      <Form :label-width="120" ref='addForm' :model='addForm' :rules="ruleValidateAddForm">
        <Row>
          <Col :span="24">
            <FormItem label="Domain Name" prop="domain">
              <Input v-model="addForm.domain" :disabled="modifyDomain"></Input>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="24">
            <FormItem label="Postmaster Password" :label-width="150" prop="password">
              <Input v-model="addForm.password"></Input>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="12">
            <FormItem label="Max Users" prop="maxUserCount">
              <InputNumber :max="1000" :min="1" v-model="addForm.maxUserCount" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
          <Col :span="12">
            <FormItem label="Max Quota(M)" prop="maxUserQuota">
              <InputNumber :max="10000" :min="-1" :step="100" v-model="addForm.maxUserQuota" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
        </Row>
        <Row>
          <Col :span="12">
            <FormItem label="Max Mails" prop="maxMailCount">
              <InputNumber :max="10000" :min="-1" :step="100" v-model="addForm.maxMailCount" style="width: 100%"></InputNumber>
            </FormItem>
          </Col>
          <Col :span="12">
            <FormItem label="Expire Time" prop="expireTime">
              <DatePicker v-model="addForm.expireTime" type="date" :options="greatNow" placeholder=""></DatePicker>
            </FormItem>
          </Col>
        </Row>
      </Form>
    </Modal>
    <Form align="left" :label-width="80" ref='searchForm' :model='searchForm' inline>
      <Button icon="md-add" type="primary" @click="prepareAdd">Add Domain</Button>
      <FormItem :label-width="0">
        <Input v-model="searchForm.domain" search enter-button="Search" @on-search="getDomainList" placeholder="Input Domain Name"
               style="width: 380px"/>
      </FormItem>
    </Form>
    <Table stripe :columns="tableColumns" :data="tableData"></Table>
  </div>
</template>

<script>
  import APIManger from "../api/"

  let moment = require('moment');

  export default {
    name: 'Domain',
    components: {},
    data() {
      var validPassword = (rule, value, callback) => {
        let that = this
        if (!that.modifyDomain) {
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
        showAddForm: false,
        modifyDomain: false,
        loading: true,
        pageSize: 20,
        page: 1,
        allData: [],
        tableData: [],
        searchForm: {
          domain: "",
          orderBy: "createTime",
          orderSort: "desc",
        },
        titleNew: "Create A New Domain",
        addForm: {
          id: "",
          domain: "",
          password: "",
          maxUserCount: 10,
          maxUserQuota: 1000,
          maxMailCount: 1000,
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
        tableColumns: [
          {
            title: 'Domain Name',
            key: 'domain',
            sortable: true
          }, {
            title: 'Users',
            key: 'userCount',
            width: 120,
            sortable: true,
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
            sortable: true,
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
            sortable: true,
            render: (h, params) => {
              let c = params.row.maxMailCount
              if (c < 0) {
                c = "unlimited"
              }
              return h("span", params.row.mailCount + "/" + c)
            }
          }, {
            title: 'Create Time',
            key: 'createTime',
            sortable: true,
            render: (h, params) => {
              return h("span", moment(params.row.createTime).format("YYYY-MM-DD HH:MM"))
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
                    type: 'ios-create'
                  },
                  style: {
                    fontSize: '20px',
                    color: '#559DF9',
                    cursor: "pointer",
                  },
                  on: {
                    click: () => {
                      let that = this
                      that.showAddForm = true
                      that.modifyDomain = true
                      that.titleNew = "Modify Domain Property"
                      that.addForm.domain = params.row.domain
                      that.addForm.maxUserCount = params.row.maxUserCount
                      that.addForm.maxUserQuota = params.row.maxUserQuota
                      that.addForm.maxMailCount = params.row.maxUserQuota
                      that.addForm.expireTime = params.row.expireTime
                    }
                  }
                }),
                h('Icon', {
                  props: {
                    type: 'ios-trash'
                  },
                  style: {
                    fontSize: '20px',
                    color: '#ed4014',
                    cursor: "pointer",
                  },
                  on: {
                    click: () => {
                      // let para = {
                      //   url: params.row.url
                      // }
                      // deleteBlackURLAPI(para).then((res) => {
                      //   if (res.data.ok == true) {
                      //     that.$Message.success(res.data.info)
                      //     that.getBlackURLList()
                      //   } else {
                      //     that.$Message.info(res.data.info)
                      //   }
                      // }).catch(function (err) {
                      //   that.$Message.error({
                      //     content: err.toString(),
                      //     duration: 5
                      //   });
                      // });
                    }
                  }
                }),
              ])
            },
          }],
      };
    },
    methods:
      {
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
        submitAdd() {
          let that = this
          that.loading = false;
          that.$refs["addForm"].validate(valid => {
            if (valid) {
              APIManger.createDomain(that.addForm).then((res) => {
                if (res.data.ok == true) {
                  that.showAddForm = false;
                  that.$Message.success(res.data.info)
                } else {
                  that.$Message.error({
                    content: res.data.info,
                    duration: 5
                  })
                }
              }).catch(function (err) {
                that.$Message.error({
                  content: err.toString(),
                  duration: 5
                });
              });
            } else {
              setTimeout(() => {
                that.loading = false;
              }, 0);
            }
          });
        }
        ,

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
        }
        ,

      }
    ,
    mounted() {
      this.getDomainList();
    }
    ,
  }
</script>
