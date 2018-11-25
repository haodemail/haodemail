<template>
    <Table border :columns="queueCol" :data="queueData" height="285"></Table>
</template>
<script>
    export default {
        data () {
            return {
                queueCol: [
                    {
                        title: 'Sender',
                        key: 'sender',
                        render: (h, params) => {
                            return h('div', [
                                h('Icon', {
                                    props: {
                                        type: 'fa fa-user'
                                    }
                                }),
                                h('strong', params.row.name)
                            ]);
                        }
                    },
                    {
                        title: 'Reciept',
                        key: 'rcpt'
                    },
                    {
                        title: 'Subject',
                        key: 'subject'
                    },
                    {
                        title: 'Time',
                        key: 'mailtime'
                    },
                    {
                        title: 'Action',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            return h('div', [
                                h('Button', {
                                    props: {
                                        type: 'primary',
                                        size: 'small'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.show(params.index)
                                        }
                                    }
                                }, 'View'),
                                h('Button', {
                                    props: {
                                        type: 'error',
                                        size: 'small'
                                    },
                                    on: {
                                        click: () => {
                                            this.remove(params.index)
                                        }
                                    }
                                }, 'Delete')
                            ]);
                        }
                    }
                ],
                queueData: [
                    {
                        name: 'John Brown',
                        rcpt: 18,
                        subject: 'New York No. 1 Lake Park'
                    },
                    {
                        name: 'Jim Green',
                        rcpt: 24,
                        subject: 'London No. 1 Lake Park'
                    },
                    {
                        name: 'Joe Black',
                        rcpt: 30,
                        subject: 'Sydney No. 1 Lake Park'
                    },
                    {
                        name: 'Jon Snow',
                        rcpt: 26,
                        subject: 'Ottawa No. 2 Lake Park'
                    }
                ]
            }
        },
        methods: {
            show (index) {
                this.$Modal.info({
                    title: 'User Info',
                    content: `Name：${this.queueData[index].name}<br>Age：${this.queueData[index].age}<br>Address：${this.queueData[index].address}`
                })
            },
            remove (index) {
                this.queueData.splice(index, 1);
            }
        }
    }
</script>