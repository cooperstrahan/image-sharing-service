<template>

    <div class="contain">
       <h1>Upload</h1>
      <div class="row mb-3">
        <label for="title" class="col-sm-2 col-form-label">Title</label>
        <div class="col-sm-10">
          <input v-model="title" type="text" class="form-control" id="title">
        </div>
      </div>
      <div class="row mb-3">
        <label for="tags" class="col-sm-2 col-form-label">Tags</label>
        <div class="col-sm-10">
          <input v-model="tags" type="text" class="form-control" id="tags">
        </div>
      </div>
      <div class="row mb-3">
        <label for="description" class="col-sm-2 col-form-label">Description</label>
        <div class="col-sm-10">
          <textarea v-model="description" class="form-control" id="description" rows="3"></textarea>
        </div>
        
      </div>
      <div class="row mb-3">
        <label>File
          <input type="file" id="file" ref="file" v-on:change="handleFileUpload()"/>
        </label>
      </div>

      <button class="btn btn-primary" v-on:click="submitFile()">Submit</button>
    </div>



</template>

<script>
import axios from 'axios';

export default {
  title: 'Upload',
  name: 'Delete',
  props: {
    msg: String
  },
  data () {
    return {
      file: '',
      title: '',
      tags: '',
      description: '',
    }
  },

  methods: {
    handleFileUpload() {
      this.file = this.$refs.file.files[0];
    },
    submitFile() {
      let formData = new FormData();
      formData.append('myFile', this.file);

      formData.append('title', this.title)
      formData.append('tags', this.tags)
      formData.append('description', this.description)

      axios
        .post("http://localhost:5000/images/", 
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }
      )
        .then(response => console.log(response))
        .catch(function(){
          console.log("FAILURE!!")
        });

        this.file= '';
        this.title='';
        this.tags='';
        this.description='';

      this.$router.push('/');
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
.contain {
  width: 50%;
  margin: auto;
  padding: 1rem;
}
</style>
