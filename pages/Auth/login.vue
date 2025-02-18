<template>
  <div class="login-form">
    <div class="for-shadow">    
      <form class="login" @submit.prevent="login">
        Логин
        <input v-model="email" class="input-login" placeholder="demo@email" required>
        Пароль
        <input v-model="password" class="input-login" placeholder="password" type="password" required>
      </form>
      <Enterbutton :mails="email" :passwords="password" :bools="bool" />
      <TheFooter />
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
.login-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  align-items: center;
  justify-content: center;
  height: 93vh; 
}

.login {
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: 200px;
}

.input-login {
  border: 1px solid #ccc;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 15px;
  transition: border-color 0.3s;
}

.input-login:hover {
  border-color: #007bff;
}

.input-login::placeholder {
  color: rgb(185, 193, 231);
}

.for-shadow {
  padding: 25px;
  border-radius: 15px;
  -webkit-box-shadow: 0px 1px 8px 6px rgba(0, 0, 0, 0.2);
  -moz-box-shadow: 0px 1px 8px 6px rgba(0, 0, 0, 0.2);
  box-shadow: 0px 1px 8px 6px rgba(0, 0, 0, 0.2);
}
</style>
