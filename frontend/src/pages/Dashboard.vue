<template>
    <div class="dashboard">
        <Row :gutter="16">
            <Col :lg="6" :md="12">
                <SystemOverView color="#41b883" icon="fa fa-database" title="Total Domains" count="996"></SystemOverView>
            </Col>
            <Col :lg="6" :md="12">
                <SystemOverView color="#1d8ce0" icon="fa fa-user" title="Total Users" count="109009"></SystemOverView>
            </Col>
            <Col :lg="6" :md="12">
                <SystemOverView color="#ffa000" icon="fa fa-envelope" title="Total Mails" count="1010"></SystemOverView>
            </Col>
            <Col :lg="6" :md="12">
                <SystemOverView color="#f60000" icon="fa fa-floppy-o" title="Total Quota" count="9960"></SystemOverView>
            </Col>
        </Row>
        <Row style="padding-top: 15px; height: 300px">
            <Col :span="6" style="padding-right: 13px;">
                <QueueOverView></QueueOverView>
            </Col>
            <Col :span="18">
                <QueueView></QueueView>
            </Col>
        </Row>
        <Row style="padding-top: 15px;">
            <Col :span="12">
                <ChartBarLine></ChartBarLine>
            </Col>
            <Col :span="12">
                <ChartBarLine></ChartBarLine>
            </Col>
        </Row>
    </div>
</template>

<script>
    import SystemOverView from '@/components/sytem-overview.vue'
    import QueueView from '@/components/queue.vue'
    import QueueOverView from '@/components/queue-overview.vue'
    import ChartBarLine from '@/components/chart-bar-line.vue'

    export default {
        name: 'Dashboard',
        components: {
            SystemOverView,
            QueueView,
            QueueOverView,
            ChartBarLine,
        },
        data() {
            return {

            }
        }
    }
</script>

<style lang="less">

</style>
