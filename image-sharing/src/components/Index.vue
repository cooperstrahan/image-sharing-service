<template>
  <div class="index">

    <!-- {{ info }} -->
    <div class="title">
      <h2 v-if="image"> {{ image.title }} </h2>
    

      <img class="card-img-top" v-bind:src="getImage(image.filename)" v-bind:alt="image.filename" />
        <div class="image-info">
        <p v-if="image">Tags: {{ image.tags }} </p> 
        <p v-if="image">Description: {{ image.description }}</p> 
      </div>
    </div>
    
  <div>

    <button type="button" class="btn btn-danger" v-on:click="deleteImage">Delete</button>
  </div>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Index',
  title: 'View Image',
  props: {
    msg: String
  },
  data () {
    return {
      image: null,
      isFetching: true,
    }
  },

  mounted () {
    // console.log(this.$route.path)
    const path = this.$route.path
    const index = path.replace('/images/','')

    let apiPath = "http://localhost:5000/images/" + index
    // console.log(apiPath)
    axios
      .get(apiPath)
      .then(response => {
        this.image = response.data
        this.isFetching = false
        })
  },
  methods :  {
    deleteImage: function() {
      const path = this.$route.path
      const index = path.replace('/images/','')
      let apiPath = "http://localhost:5000/images/" + index
      axios
      .delete(apiPath)
      .then(response => {
        console.log(response);
        const status = JSON.parse(response.status);

        if (status == '200') {
          this.$router.push('/');
        }
      })
      .catch(error => {
        console.log(error);
      });
    } ,

    getImage(image) {
      try {
        return require('../../../api/cmd/main/images/' + image)
      } catch (e) {
        console.log(e);
        return  
      }
    }
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
.middle-image {
  width: 20rem;
}
.title {
  padding: 10px;
  margin: auto;
  width: 20rem;
}
.image-info {
  margin-top: 20px;
  text-align: left;
}
</style>
