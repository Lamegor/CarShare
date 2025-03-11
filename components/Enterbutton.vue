<template>
  <div>
    <button type="button" class="vhod" @click="Listener">
      ВХОД
    </button>
  </div>
</template>

<script>
export default {
  name: 'EnterButton',

  props: {
    mails: {
      type: String,
      default: ''
    },
    passwords: {
      type: String,
      default: ''
    },
    confirmpasswords: {
      type: String,
      default: ''
    },
    bools: Boolean,
    names: {
      type: String,
      default: ''
    }
  },

  data() {
  return {
    csrfToken: '',
    // другие данные
  }
},
mounted() {
  // Запрашиваем CSRF токен при загрузке компонента
  // this.GetCSRFToken();
},
methods: {
  Listener() {
    if (this.bools) {
      if (this.mails && this.passwords) {
        this.apisLog();
      } else {
        alert('Остались незаполненные поля');
      }
    } else {
      if (this.passwords !== this.confirmpasswords) {
        alert('Пароли не совпадают');
      } else if (this.passwords && this.confirmpasswords && this.mails && this.names) {
        this.apisReg();
      } else {
        alert('Остались незаполненные поля');
      }
    }
  },
  // async GetCSRFToken() {
  //   try {
  //     let csrf = await fetch('http://localhost:8080/auth/csrf-token');

  //     // Проверяем, успешен ли ответ
  //     if (!csrf.ok) {
  //       throw new Error(`Ошибка сети: ${csrf.status} ${csrf.statusText}`);
  //     }

  //     // Получаем CSRF токен из заголовков
  //     const csrfToken = csrf.headers.get('X-Csrf-Token');
  //     if (!csrfToken) {
  //       throw new Error('CSRF токен не найден в заголовках ответа.');
  //     }

  //     // Сохраняем CSRF токен в состоянии компонента
  //     this.csrfToken = csrfToken;
  //   } catch (error) {
  //     console.error('Ошибка при получении CSRF токена:', error);
  //   }
  // },
  apisReg() {
    fetch('http://localhost:8080/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json'},
      body: JSON.stringify({
        username: this.names,
        email: this.mails,
        '1_password': this.passwords,
        '2_password': this.confirmpasswords,
      })
    })
    .then((res) => {
      if (res.status === 201) {
        this.$router.push({ name: 'Auth-login' });
      } else {
        alert(res.status);
      }
    })
    .catch((error) => {
      console.error('Ошибка при регистрации:', error);
    });
  },
  apisLog() {
    fetch('http://localhost:8080/auth/login', {
      method: 'POST',
      headers: { accept: 'application/json', 'Content-Type': 'application/json'},
      body: JSON.stringify({
        username: this.names,
        password: this.passwords,
      })
    })
    .then((res) => {
      if (res.status === 200) {
        const setCookieHeader = response.headers['set-cookie'];
        if (setCookieHeader) {
          // Устанавливаем куку с помощью cookie-universal-nuxt
          this.$cookies.set('myCookie', setCookieHeader, {
            path: '/', // Путь, на котором кука доступна
            maxAge: 60 * 60 * 24, // Время жизни куки в секундах (1 день)
            secure: false, // Установите true, если используете HTTPS
          });
        }
        this.$router.push({ name: 'index' });
      } else {
        alert(res.status);
      }
    })
    .catch((error) => {
      console.error('Ошибка при входе:', error);
    });
  }
  }
}
</script>

<style>
.vhod {
  text-align: center;
  margin-top: 15px;
  border: 1px solid rgba(0, 255, 127, 0.676);
  width: 200px;
  padding: 20px;
  border-radius: 7px;
  background-color: rgba(0, 255, 128, 0.676);
  color: white;
  font-size: medium;
  display: inline-block;
  cursor: pointer;
}
.vhod:hover {
  background-color: rgba(0, 255, 128, 0.276);
}
</style>
