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

  methods: {
    Listener () {
      this.GetCSRFToken()
      if (this.bools) {
        if (this.mails && this.passwords) {
          this.GetCSRFToken()
          this.apisLog()
        } else {
          alert('Остались незаполненные поля')
        };
      } else {
        if (this.passwords !== this.confirmpasswords) {
          alert('Пароли не совпадают')
        } else if (this.passwords && this.confirmpasswords && this.mails && this.names) {
          this.apisReg()
        } else {
          alert('Остались незаполненные поля')
        };
      };
    },
    // async GetCSRFToken () {
    //   let response = await fetch('http://localhost:8080/auth/csrf-token');

    //   if (response.ok) { // если HTTP-статус в диапазоне 200-299
    //     let json = await response.json();
    //   } else {
    //     alert("Ошибка HTTP: " + response.status);
    //   }
    //   // const csrfToken = await this.$axios.$get('http://localhost:8080/auth/csrf-token');
    //   // this.csrfToken = csrfToken.X-Csrf-Token;
    //   this.csrfToken = json.X-Csrf-Token;
    // },
    apisReg () {
      fetch('http://localhost:8080/auth/register', {
        method: 'POST',
        headers: {'Content-Type': 'application/json', 'X-Csrf-Token': csrfToken },
        body: JSON.stringify(
          {
            username: this.names,
            email: this.mails,
            '1_password': this.passwords,
            '2_password': this.confirmpasswords,
          })
      })
        .then((res) => {
          if (res.status === 201) {
            this.$router.push({ name: 'Auth-login' })
          } else {
            alert(res.status)
          }
        })
        .then((res) => {
          if (res.status === 400) {
            alert(res)
          }
        })
      },
    apisLog () {
      fetch('https://localhost:8080/auth/login', {
        method: 'POST',
        headers: { accept: 'application/json', 'Content-Type': 'application/json', 'X-Csrf-Token': this.csrfToken },
        body: JSON.stringify(
          {
            username: this.names,
            password: this.passwords,
          })
      })
        .then((res) => {
          if (res.status === 200) {
            this.$router.push({ name: 'index'})
          }
        })
        .then((res) => {
          if (res.status === 401) {
            alert(res)
          }
        })
      }
  },

  computed: {
    csrfToken() {
      return this.$store.state.csrfToken;
    }
  },
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
