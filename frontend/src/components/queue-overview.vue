<template>
    <Card style="height: 285px">
        <p slot="title" style="text-align: left">
            <Icon type="ios-film-outline"></Icon>
            Queue Statics
        </p>
        <p v-for="q in queueCount" class="queuename">
            {{ q.name }}
            <span>
                {{ q.count }}
            </span>
        </p>
    </Card>
</template>
<script>
    export default {
        data () {
            return {
                name: "queue-overview",
                queueCount: [
                    {
                        "name": "local",
                        "count": 10,
                    }, {
                        "name": "remote",
                        "count": 10,
                    }, {
                        "name": "defer",
                        "count": 10,
                    }, {
                        "name":"bounce",
                        "count": 10,
                    },
                ],
            }
        },
        methods: {

        },
        mounted () {
        }
    }
</script>

<style lang="less" scoped>
    .queuename {
        font-size: 16px;
        font-color: #19be6b;
        line-height: 2.5;
        text-align: left;
        padding-left: 50px;

    }
</style>