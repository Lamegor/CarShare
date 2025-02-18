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
.register-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  align-items: center;
  justify-content: center;
  height: 93vh; 
  z-index: 5;
  background-image: url('https://zastavki.gas-kvas.com/uploads/posts/2024-09/thumbs/zastavki-gas-kvas-com-jop5-p-zastavki-na-telefon-dlya-muzhchin-krutie-m-5.jpg');
}
.register {
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: 400px;
}

.label {
  margin-bottom: 5px;
  font-weight: bold;
}

.input-register {
  border: 1px solid #fffefe;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 15px;
  transition: border-color 0.3s;
}

.input-register::placeholder {
  color: rgb(59, 85, 214);
}

.input-register:hover {
  border-color: #007bff;
}

.for-shadow {
  padding: 25px;
  border-radius: 15px;
  -webkit-box-shadow: 0px 1px 8px 6px rgba(0, 0, 0, 0.2);
  -moz-box-shadow: 0px 1px 8px 6px rgba(0, 0, 0, 0.2);
  box-shadow: 0px 1px 8px 6px rgba(255, 255, 255, 0.2);
  background-color: rgb(24, 88, 224);
}
</style>
