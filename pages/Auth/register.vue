<template>
  <div>
    <div class="register-form">
      <div class="for-shadow">
        <div class="register-input">
          <form class="register" @submit.prevent="register">
            <label for="email">Почтовый ящик</label>
            <input id="email" v-model="mail" class="input-register" type="email" required placeholder="Введите ваш email">
            
            <label for="name">Имя</label>
            <input id="name" v-model="name" class="input-register" required placeholder="Введите ваше имя">
            
            <label for="password">Пароль</label>
            <input id="password" v-model="password" class="input-register" type="password" required placeholder="Введите пароль">
            
            <label for="confirm-password">Повтор пароля</label>
            <input id="confirm-password" v-model="confirmpassword" class="input-register" type="password" required placeholder="Повторите пароль">
          </form>
        </div>

        <div class="button-container">
          <Enterbutton
            :mails="mail"
            :passwords="password"
            :bools="bool"
            :confirmpasswords="confirmpassword"
            :names="name"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'RegIster',

  data () {
    return {
      mail: '',
      password: '',
      confirmpassword: '',
      bool: false,
      name: ''
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
  padding-bottom: 10px;
}

.register-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 93vh; 
  background: rgba(26, 26, 26, 0.97);
  border-radius: 10px;
}

.register {
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: 400px;
}

.input-register {
  background-color: #333;
  border: 1px solid #0BDA51;
  color: #FFF;
  padding: 10px;
  width: 100%;
  border-radius: 5px;
  margin-bottom: 15px;
  transition: border-color 0.3s;
}

.input-register::placeholder {
  color: #00FF7F;
}

.input-register:hover {
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

form button {
  background-color: #00FF7F;
  color: #0A0A0A;
  padding: 10px 20px;
  font-weight: bold;
  border-radius: 5px;
  border: none;
}

form button:hover {
  background-color: #0BDA51;
  color: #FFF;
}
</style>
