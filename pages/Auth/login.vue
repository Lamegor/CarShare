<template>
  <div>
    <div class="login-form">
      <div class="for-shadow">    
        <form class="login" @submit.prevent="login">
          <label for="email">Логин</label>
          <input id="email" v-model="email" class="input-login" type="email" required placeholder="demo@email">
          
          <label for="password">Пароль</label>
          <input id="password" v-model="password" class="input-login" type="password" required placeholder="password">
        </form>
        
        <div class="button-container">
          <Enterbutton :mails="email" :passwords="password" :bools="bool" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LogIn',

  data () {
    return {
      email: '',
      password: '',
      bool: true
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: Arial, sans-serif;
  background-color: #0A0A0A;
  background-image: url('/static/img/BUSHIDO.png');
  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
  color: #FFF;
  overflow-x: hidden;
}
label {
  padding-bottom: 5px;
}

.login-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 93vh; 
  background: rgba(26, 26, 26, 0.97);
  border-radius: 10px;
}

.login {
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: 400px;
}

.input-login {
  background-color: #333;
  border: 1px solid #0BDA51;
  color: #FFF;
  padding: 10px;
  width: 100%;
  border-radius: 5px;
  margin-bottom: 15px;
  transition: border-color 0.3s;
}

.input-login::placeholder {
  color: #00FF7F;
}

.input-login:hover {
  border-color: #00FF7F;
}

.for-shadow {
  padding: 25px;
  border-radius: 15px;
  box-shadow: 0px 1px 8px 6px rgba(0, 255, 127, 0.2);
  background-color: #1A1A1A;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.button-container {
  display: flex;
  justify-content: center;
  width: 100%;
  margin-top: 15px;
}
</style>
