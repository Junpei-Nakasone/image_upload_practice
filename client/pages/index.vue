<template>
  <div class="container">
    <h2 class="text-center page-title">商品登録</h2>
    <b-form enctype="multipart/form-data"/>
    <!-- <b-form-group
      id="exampleInputGroup1"
      label="商品名"
      class="form-margin"
      label-for="exampleInput1"
    >
      <b-form-input
        id="exampleInput1"
        v-model="name"
        required />
    </b-form-group>

    <b-form-group
      id="exampleInputGroup2"
      label="価格"
      label-for="exampleInput2"
      class="form-margin"
    >
      <b-form-input
        id="exampleInput2"
        type="number"
        v-model="price"
        required />
    </b-form-group>

    <b-form-group id="exampleInputGroup3" label="商品について" label-for="exampleInput3" class="form-margin">
      <b-form-textarea
      id="textarea"
      rows="3"
      v-model="description"
      max-rows="6"
    />
    </b-form-group>

    <b-form-group
      id="exampleInputGroup4"
      label="在庫数"
      label-for="exampleInput4"
      class="form-margin"
    >
      <b-form-input
        id="exampleInput4"
        type="number"
        class="form-width"
        v-model="itemCount"
        required />
    </b-form-group> -->

    <b-form-group
      id="exampleInputGroup5"
      label="商品画像"
      label-for="exampleInput5"
      class="form-margin"
    >

  <b-form-file class="mt-3" @change="onUpload()" ref="files" multiple plain />
    </b-form-group>
    <div class="text-center">
      <b-button type="submit" variant="outline-info" @click="createProduct()">登録する</b-button>
    </div>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      name: "",
      price: 0,
      description: "",
      itemCount: 0,
      images: "",
    }
  },
  methods: {
    onUpload: function() {
      this.images = event.target.files;
    },
    createProduct: function() {
      let formData = new FormData
      formData.append('name', this.name)
      formData.append('price', this.price)
      formData.append('description', this.description)
      formData.append('item_count', this.itemCount)
      formData.append('shop_id', this.shopId)
      // for( let i = 0; i < this.images.length; i++) {
      //   let image = this.images[i];
      //   formData.append('images[]', image);
      // }
      formData.append('image', this.images[0])
      this.$axios.$post('http://localhost:9999/upload',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(response => {
        console.log(response.data.status);
      }).catch(error => {
        console.log(error);
      })
    }
  }
}
</script>
