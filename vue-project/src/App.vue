<template>
  <div id="app">
    <form v-on:submit.prevent="create" >
          <div>
            <label v-if="this.isEdit == false"><h1>Add New user</h1></label>
            <label v-else><h1>Edit user</h1></label>
          </div>
          <input v-model="first_name" type="text" class="form-control" id="fnameinput" placeholder="first name">
          <input v-model="last_name" type="text" class="form-control" id="lnameinput" placeholder="lasr name">
          <input v-model="email" type="text" class="form-control" id="emailinput" placeholder="e-mail">
          <input v-model="gender" type="text" class="form-control" id="genderinput" placeholder="gender">
          <input v-model="age" type="text" class="form-control" id="ageinput" placeholder="age">
          <button v-if="this.isEdit == false" type="submit" class="btn btn-success btn-block mt-3">
            Add user
          </button>
          <button v-else v-on:click="updateUser()" type="button" class="btn btn-primary btn-block mt-3">
            Update
          </button>
          <button v-if="this.isEdit == true" v-on:click="cancelUpdate()" type="button" class="btn btn-primary btn-block mt-3 btn btn-danger">
            Cancel
          </button>
        </form>

    <table class="table">
            <tr>
              <th scope="col">id</th>
              <th scope="col">first name</th>
              <th scope="col">last name</th>
              <th scope="col">email</th>
              <th scope="col">gender</th>
              <th scope="col">age</th>
              <th scope="col">modify</th>
            </tr>
          <tr v-for="user in users" v-bind:key="user.id">
            <td v-for="value in user" v-bind:key="value.id">{{value}}</td>
            <td>
              <button v-on:click="editUser(user)" class="btn btn-info" >Edit</button>
              <button v-on:click="deleteUser(user[0])" class="btn btn-danger">Delete</button>
            </td>
          </tr>
        </table>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data: function(){
    return{
      users: [],
      id: '',
      first_name: '',
      last_name: '',
      email: '',
      gender: '',
      age: '',
      isEdit: false
    }
  },mounted(){
    this.showAll()
  },
  methods: {
      showAll: function() {
        axios.get('http://localhost:8000/showAll')
        .then((response) => {
          // handle success
          this.users = response.data;
          console.log(this.users);
        })
      },
      create: function() {
        axios.post('http://localhost:8000/create',{
          First_name : this.first_name,
          Last_name : this.last_name,
          Email : this.email,
          Gender : this.gender,
          Age : this.age,
        })
        
      }
  },
  name: 'app'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
