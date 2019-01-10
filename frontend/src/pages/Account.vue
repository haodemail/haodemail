<template>
    <div>
        <br>
        <Modal v-model="showAddForm" :loading="loading" title="Add Domain" @on-ok="submitAdd">
            <Form :label-width="85" ref='addForm' :model='addForm' :rules="ruleValidateAddForm">
                <Row>
                    <Col :span="24">
                        <FormItem label="Domain Name" prop="domain">
                            <Input v-model="addForm.domain"></Input>
                        </FormItem>
                    </Col>
                </Row>
                <Row>
                    <Col :span="12">
                        <FormItem label="Max Users" prop="maxUsers">
                            <Input :max="100" :min="1" v-model="addForm.maxUsers"></Input>
                        </FormItem>
                    </Col>
                    <Col :span="12">
                        <FormItem label="Max Quota" prop="maxQuota">
                            <Input :max="102400" :min="1024" v-model="addForm.maxQuota"></Input>
                        </FormItem>
                    </Col>
                </Row>
                <Row>
                    <Col :span="24">
                        <FormItem label="Expire Time" prop="expireTime">
                            <DatePicker v-model="addForm.expireTime" type="date" placeholder=""></DatePicker>
                        </FormItem>
                    </Col>
                </Row>
            </Form>
        </Modal>
        <Form align="left" :label-width="80" ref='searchForm' :model='searchForm' inline>
            <Button type="primary" @click="showAddForm=true">Add Domain</Button>
            <Input v-model="searchForm.keyword" placeholder="Input Domain Name" style="width: 350px"/>
            <Button type="success" @click="getDomainList">Search</Button>
        </Form>
        <br>
        <Table stripe :columns="tableColumns" :data="tableData"></Table>
    </div>
</template>

<script>
    var moment = require('moment');

    export default {
        name: 'Domain',
        components: {},
        data() {
            return {
                showAddForm: false,
                loading: false,
                pageSize:15,
                page:1,
                searchForm: {
                    keyword: "",
                },
                addForm: {
                    domain:"",
                    maxUsers:10,
                    maxQuota:1024000,
                    expireTime:new Date(2019,12,30),
                },
                ruleValidateAddForm:{},
                tableColumns: [
                    {
                        title: '',
                        key: 'id',
                        width: 50,
                    },
                    {
                        title: 'Domain Name',
                        key: 'domain',
                        sortable: true
                    },
                    {
                        title: 'Total Users',
                        key: 'totalUsers',
                        width: 120,
                        sortable: true
                    },
                    {
                        title: 'Total Quota',
                        key: 'totalQuota',
                        width: 120,
                        sortable: true
                    },
                    {
                        title: 'Total Mails',
                        key: 'totalMails',
                        width: 120,
                        sortable: true
                    },
                    {
                        title: 'Create Time',
                        key: 'createTime',
                        sortable: true
                    },
                    {
                        title: 'Operation',
                        key: '',
                        width: 200,
                        align: 'center',
                        render: (h, params) => {
                            return h("div", [
                                h(
                                    "Button", {
                                        props: {
                                            type: "info",
                                            size: "small",
                                        },
                                        style: {
                                            marginRight: '5px'
                                        },
                                        on: {
                                            click: () => {
                                            }
                                        }
                                    },
                                    "Configure"
                                ),
                                h(
                                    "Button", {
                                        props: {
                                            type: "error",
                                            size: "small",
                                        },
                                        on: {
                                            click: () => {
                                            }
                                        }
                                    },
                                    "Delete"
                                ),

                            ]);
                        }
                    }
                ],
                allData:[],
                tableData: [
                    {
                        id: 1,
                        domain: 'haodemail.com',
                        totalUsers: 100,
                        totalQuota: 10240000,
                        totalMails: 30000,
                        createTime: new Date(),
                    },
                    {
                        id: 2,
                        domain: 'haodemail.com',
                        totalUsers: 100,
                        totalQuota: 10240000,
                        totalMails: 30000,
                        createTime: new Date(),
                    },
                    {
                        id: 3,
                        domain: 'haodemail.com',
                        totalUsers: 100,
                        totalQuota: 10240000,
                        totalMails: 30000,
                        createTime: new Date(),
                    },
                    {
                        id: 4,
                        domain: 'haodemail.com',
                        totalUsers: 100,
                        totalQuota: 10240000,
                        totalMails: 30000,
                        createTime: new Date(),
                    },
                    {
                        id: 5,
                        domain: 'haodemail.com',
                        totalUsers: 100,
                        totalQuota: 10240000,
                        totalMails: 30000,
                        createTime: new Date(),
                    },
                ]
            }
        },
        methods: {
            submitAdd() {
                let that = this
                that.loading = false;
                that.$refs["addForm"].validate(valid => {
                    if (valid) {
                      $http.get(`/api`, { params: params }).then((res) => {
                            if (res.data.ok == true) {
                                let domainData = res.data.data
                                that.AllData.unshift(domainData)
                                that.tableData.unshift(domainData)
                                that.tableData.pop()
                                that.showAddForm = false;
                                that.$Message.success(res.data.info)
                            } else {
                                that.$Message.info(res.data.info)
                            }
                        }).catch(function (err) {
                            self.$Message.error({
                                content: resp.data.info,
                                duration: 5
                            });
                        });
                    } else {
                        setTimeout(() => {
                            that.loading = true;
                        }, 0);
                    }
                });
            },
            
            getDomainList() {
                let that = this
                let para = {
                    keyword: this.searchForm.keyword,
                    pagesize: that.pageSize,
                    page: that.page,
                }
                listDomainAPI(para).then((res) => {
                    if (res.data.ok == true) {
                        if (res.data.data) {
                            that.alldata = res.data.data
                            that.tableData = that.allData.slice(0, that.pageSize);
                            that.page = 1
                        } else {
                            this.allData = []
                            this.tableData = []
                            this.page = 1
                        }
                    }
                })
            },

        },
        mounted() {
            this.getDomainList();
        },
    }
</script>
