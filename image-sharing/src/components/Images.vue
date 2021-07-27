<template>
  <div class="hello">
    <h1>All images</h1>

  <div>
    <li 
    v-for='image in info' 
    :key='image.id'
    >
    
     <img class="small-image" v-bind:src="require('../../../api/cmd/main/images/' + image.filename)" /> 
     <router-link
        class="nav-link active"
        aria-current="page"
        v-bind:to="'/images/' + image.id"
      >
     <h3> {{image.title }} </h3>
     </router-link
                >
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
.small-image {
  max-width: 10rem;
}
</style>
