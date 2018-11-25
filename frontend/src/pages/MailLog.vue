<template>
  <VmImageList :data="dataImageList" @delete-ok="deletefn" class="vm-margin"></VmImageList>
</template>

<script>
  import VmImageList from '@/components/vm-image-list'
  export default {
    name: 'ImageList',
    components: {
      VmImageList
    },
    methods: {
      deletefn: function (data) {
        for (let i = 0; i < this.dataImageList.length; i++) {
          if (this.dataImageList[i].id === data.id) {
            this.dataImageList.splice(i, 1)
          }
        }
      }
    },
    data: function () {
      return {
        dataImageList: [
          {
            id: '201707101552',
            title: 'Title1',
            img: require('@/assets/img/photo.jpg'),
            desc: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry,Lorem Ipsum has been the industry\'s standard dummy text ever sincc the 1500s ly dummy tly dummy tly dummy tly dummy tly dummy tly dummy t',
            detailUrl: '#',
            editUrl: '#'
          },
          {
            id: '201707101553',
            title: 'Title2',
            img: require('@/assets/img/photo.jpg'),
            desc: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry,Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s',
            detailUrl: '#',
            editUrl: '#'
          },
          {
            id: '201707101554',
            title: 'Title3',
            img: require('@/assets/img/photo.jpg'),
            desc: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry,Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s',
            detailUrl: '#',
            editUrl: '#'
          },
          {
            id: '201707101555',
            title: 'Title4',
            img: require('@/assets/img/photo.jpg'),
            desc: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry,Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s',
            detailUrl: '#',
            editUrl: '#'
          },
          {
            id: '201707101556',
            title: 'Title5',
            img: require('@/assets/img/photo.jpg'),
            desc: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry,Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s',
            detailUrl: '#',
            editUrl: '#'
          },

        ]
      }
    }
  }
</script>
