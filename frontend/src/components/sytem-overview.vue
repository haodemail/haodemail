<template>
  <Row class="pt-state-overview" type="flex">
    <Col span="10" class="symbol" :style="{ backgroundColor: color }">
      <i :class="icon"></i>
    </Col>
    <Col span="14" class="value">
      <h2 class="count" :style="{color: color}">{{ count }}</h2>
      <p>{{ title }}</p>
    </Col>
  </Row>
</template>

<script>
export default {
  name: 'StateOverView',
  props: ['color', 'icon', 'title', 'count']
}
</script>

<style lang="less">
  .pt-state-overview {
    background-color: white;
    border-radius: 4px;
    min-height: 100px;
    overflow: hidden;
  }

  .symbol{
    color: white;
    font-size: 38px;
    background-color: white;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .value{
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .value p {
    font-size: 16px;
    color: #99a9c0;
  }
</style>