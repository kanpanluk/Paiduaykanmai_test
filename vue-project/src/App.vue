<template>
  <div id="app">
    
          <div>
            <label><h1></h1></label>
          </div>
          <button type="button" class="btn btn-info btn-lg" data-toggle="modal" data-target="#createModal">Create New User</button>
          <!-- Modal -->
          <div class="modal fade" id="createModal" role="dialog">
            <div class="modal-dialog">
            
              <!-- Modal content-->
              <div class="modal-content">
                <div class="modal-header">
                  <button type="button" class="close" data-dismiss="modal">&times;</button>
                  <h4 >Fill Form</h4>
                </div>
                <div class="modal-body">
                  <input v-model="first_name" type="text" class="form-control" id="fnameinput" placeholder="first name">
                  <input v-model="last_name" type="text" class="form-control" id="lnameinput" placeholder="lasr name">
                  <input v-model="email" type="text" class="form-control" id="emailinput" placeholder="e-mail">
                  <input v-model="gender" type="text" class="form-control" id="genderinput" placeholder="gender">
                  <input v-model="age" type="text" class="form-control" id="ageinput" placeholder="age">
                </div>
                <div class="modal-footer">
                   <button v-on:click="create()" data-dismiss="modal" type="submit" class="btn btn-success btn-block mt-3">
                    Submit
                  </button>
                </div>
              </div>
              
            </div>
          </div>
   
    <div></div>
    <table class="table" align="center">
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
              <button v-on:click="edit(users)" data-toggle = "modal" data-target="#editModal" class="btn btn-info" >Edit</button>
              <div class="modal fade" id="editModal" role="dialog">
                <div class="modal-dialog">
                
                  <!-- Modal content-->
                  <div class="modal-content">
                    <div class="modal-header">
                      <button type="button" class="close" data-dismiss="modal">&times;</button>
                      <h4 >Edited</h4>
                    </div>
                    <div class="modal-body">
                      <input v-model="first_name" type="text" class="form-control" id="fnameinput" placeholder="first name">
                      <input v-model="last_name" type="text" class="form-control" id="lnameinput" placeholder="lasr name">
                      <input v-model="email" type="text" class="form-control" id="emailinput" placeholder="e-mail">
                      <input v-model="gender" type="text" class="form-control" id="genderinput" placeholder="gender">
                      <input v-model="age" type="text" class="form-control" id="ageinput" placeholder="age">
                    </div>
                    <div class="modal-footer">
                      <button v-on:click="update()" data-dismiss="modal" type="submit" class="btn btn-success btn-block mt-3">
                        Submit
                      </button>
                    </div>
                  </div>
                  
                </div>
              </div>
              <button v-on:click="del(user[0])" class="btn btn-danger">Delete</button>
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
          // console.log(this.users);
        })
      },
      create: function() {
        const params = new URLSearchParams();
        params.append('first_name', this.first_name);
        params.append('last_name', this.last_name);
        params.append('email', this.email);
        params.append('gender', this.gender);
        params.append('age', this.age);
        axios.post('http://localhost:8000/create',params)
        .then((response) => {
          this.first_name = ''
          this.last_name =''
          this.email = ''
          this.gender = ''
          this.age = ''
          // this.showAll()
          console.log(response)
        })
        .catch(err => {
          console.log(err)
        })  
      },
      edit(user){
        alert(this.isEdit)
        this.id = user[0]
        this.first_name = user[1]
        this.last_name = user[2]
        this.email = user[3]
        this.gender = user[4]
        this.age = user[5]
        this.isEdit = true
        window.scrollTo(0, 0);
        alert(this.isEdit)
      },
      update: function() {
        const params = new URLSearchParams();
        params.append('first_name', this.first_name);
        params.append('last_name', this.last_name);
        params.append('email', this.email);
        params.append('gender', this.gender);
        params.append('age', this.age);
        axios.post('http://localhost:8000/update',params)
        .then((response) => {
          this.first_name = ''
          this.last_name =''
          this.email = ''
          this.gender = ''
          this.age = ''
          // this.showAll()
          console.log(response)
        })
        .catch(err => {
          console.log(err)
        })  
      },
      del: function(id){
      axios.get(`http://localhost:8000/delete/${id}`)
      .then((response) => {
          this.showAll()
          alert(id);
          
        })
        .catch(err => {
          console.log(err)
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
