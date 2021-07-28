<template>
  <div class="hello">
    <h1>All images</h1>

  <div>

    

    <li 
    v-for='image in info' 
    :key='image.id'
    >

      <div class="card image-card">
      <img class="card-img-top" v-bind:src="require('../../../api/cmd/main/images/' + image.filename)" />
      <div class="card-body">
        <h5 class="card-title">{{image.title}}</h5>
        <p class="card-text">{{image.description}}</p>
        <router-link
        class="btn btn-primary"
        aria-current="page"
        v-bind:to="'/images/' + image.id"
      >View Image</router-link>
        </div>
      </div>
    
      
     
    </li>


  
    </div>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Images',
  title: 'View Images',
  props: {
    msg: String
  },
  data () {
    return {
      info: null,
    }
  },

  mounted () {
    axios
      .get("http://localhost:5000/images")
      .then(response => (this.info = response.data))    
  },

  computed: {
      // generateImageStrings: function () {
      //   console.log("generate Image strings called")
      //   return this.info.map(function (imageObject) {
      //     imageObject.filename = "../../../api/cmd/main/images/" + imageObject.filename   
      //     return imageObject
      //   }) 
          
    // }
  }
}


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.hello {
  margin: 15px;
}
/* .small-image {
  max-width: rem;
} */
.image-card {
  /* height: 20rem; */
  width: 18rem;
  margin: 1rem;
}
</style>
