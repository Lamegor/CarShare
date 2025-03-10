<template>
  <div id="app">
    <!-- Прелоадер -->
    <div v-if="loading" class="preloader">
      <div class="spinner"></div>
    </div>

    <header>
      <nav>
        <ul>
          <li><a href="#">Домой</a></li>
          <li><a href="#services">Услуги</a></li>
          <li><a href="#pricing">Цены</a></li>
          <li><a href="#contact">Контакты</a></li>
        </ul>
        <a href="/Auth/login" class="login-btn">Вход</a>
      </nav>
    </header>


    <main>
      <!-- Hero секция -->
      <section class="hero">
        <div class="animation">
          <img src="/img/sti4k31.png" alt="Бубару" />
        </div>
        <div class="hero-content">
          <h2>Добро пожаловать в</h2>
          <h1>EGK Cars</h1>
          <p>Ваш каршеринг</p>
          <a href="/Auth/register" class="cta-btn">Начать</a>
        </div>
      </section>

      <!-- Раздел Услуги -->
      <section id="services" class="services">
        <h2>Наши Услуги</h2>
        <div class="service-options">
          <div v-for="service in services" :key="service.title" class="service-option">
            <h3>{{ service.title }}</h3>
            <p>{{ service.description }}</p>
          </div>
        </div>
      </section>

      <!-- Раздел Цены -->
      <section id="pricing" class="pricing">
        <h2>Цены</h2>
        <div class="pricing-options">
          <div class="pricing-option">
            <img src="img/carN.png" alt="Стандарт" />
            <h3>Стандарт</h3>
            <p>От 1000 рублей в сутки</p>
            <p>Для краткосрочных поездок и аренды автомобилей на один день.</p>
          </div>
          <div class="pricing-option">
            <img src="img/carE.png" alt="Эконом" />
            <h3>Эконом</h3>
            <p>От 5000 рублей в неделю</p>
            <p>Для аренды на неделю с выгодными условиями.</p>
          </div>
          <div class="pricing-option">
            <img src="img/carP.png" alt="Бизнес" />
            <h3>Бизнес</h3>
            <p>От 20000 рублей в месяц</p>
            <p>Для долгосрочной аренды с гарантией лучших автомобилей и обслуживания.</p>
          </div>
        </div>
      </section>

      <!-- Раздел Контакты -->
      <section id="contact" class="contact">
        <h2>Контакты</h2>
        <p>Свяжитесь с нами для получения дополнительной информации.</p>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label for="name">Имя</label>
            <input type="text" id="name" v-model="form.name" placeholder="Ваше имя" required>
          </div>
          <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" v-model="form.email" placeholder="Ваш email" required>
          </div>
          <div class="form-group">
            <label for="message">Сообщение</label>
            <textarea id="message" v-model="form.message" placeholder="Ваше сообщение" required></textarea>
          </div>
          <button type="submit" class="cta-btn">Отправить</button>
        </form>
      </section>
    </main>

    <footer>
      <p>&copy; 2025 EGK Brothers. Все права защищены.</p>
    </footer>
  </div>
</template>

<script setup>
  import {
    ref,
    onMounted
  } from 'vue';

  const loading = ref(true);
  const services = ref([{
      title: 'Каршеринг',
      description: 'Мы предоставляем удобные автомобили для аренды.'
    },
    {
      title: 'Долгосрочная аренда',
      description: 'Автомобили на длительный срок с возможностью выкупа.'
    },
    {
      title: 'Сервисное обслуживание',
      description: 'Техническое обслуживание автомобилей.'
    }
  ]);

  const form = ref({
    name: '',
    email: '',
    message: ''
  });

  const submitForm = () => {
    alert(`Спасибо, ${form.value.name}! Мы свяжемся с вами.`);
    form.value = {
      name: '',
      email: '',
      message: ''
    };
  };

  onMounted(() => {
    setTimeout(() => {
      loading.value = false;
    }, 1000);
  });

</script>

<style>
  /* General Styles */
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  /* Кастомная прокрутка */
  ::-webkit-scrollbar {
    width: 12px;
    height: 12px;
  }

  ::-webkit-scrollbar-thumb {
    background-color: #00FF7F;
    border-radius: 10px;
    border: 3px solid #1A1A1A;
  }

  ::-webkit-scrollbar-thumb:hover {
    background-color: #0BDA51;
  }

  ::-webkit-scrollbar-track {
    background: #333;
    border-radius: 10px;
  }

  ::-webkit-scrollbar-track:hover {
    background: #444;
  }

  /* Для горизонтальной полосы прокрутки */
  ::-webkit-scrollbar-horizontal {
    height: 12px;
  }

  ::-webkit-scrollbar-thumb.horizontal {
    background-color: #00FF7F;
    border-radius: 10px;
  }

  /* Общие стили для прелоадера */
  .preloader {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.8);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
    visibility: visible;
    opacity: 1;
    transition: opacity 0.5s ease, visibility 0.5s ease;
  }

  .preloader.hidden {
    visibility: hidden;
    opacity: 0;
  }

  /* Анимация вращающегося кольца */
  .spinner {
    width: 80px;
    height: 80px;
    border: 8px solid #f3f3f3;
    border-top: 8px solid #00FF7F;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }

    100% {
      transform: rotate(360deg);
    }
  }

  /* Подключение шрифта */
  @font-face {
    font-family: 'Dancing Script';
    src: url('/font/DancingScript.ttf') format('truetype');
    font-weight: normal;
    font-style: normal;
  }

  /* Общие стили для body */
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

  /* Стили для хедера */
  header {
    background-color: #1A1A1A;
    padding: 15px;
  }

  nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .login-btn {
    margin-left: auto;
    background-color: #00FF7F;
    color: #0A0A0A;
    padding: 10px 20px;
    text-decoration: none;
    font-weight: bold;
    border-radius: 5px;
  }

  .login-btn:hover {
    background-color: #0BDA51;
    color: #FFF;
  }


  nav ul {
    display: flex;
    justify-content: center;
    list-style-type: none;
  }

  nav ul li {
    margin: 0 20px;
  }

  nav ul li a {
    text-decoration: none;
    color: #0BDA51;
    font-size: 18px;
  }

  nav ul li a:hover {
    color: #00FF7F;
  }

  /* Стили для hero секции */
  .hero {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 100vh;
    padding: 0 20px;
    position: relative;
    overflow: hidden;
  }

  .hero-content {
    flex: 1;
    max-width: 50%;
    padding-left: 20px;
    text-align: left;
  }

  .hero h1 {
    font-family: 'Dancing Script';
    font-size: 80px;
    margin-bottom: 20px;
  }

  .hero h2 {
    font-size: 50px;
    margin-bottom: 20px;
  }

  .hero p {
    font-size: 20px;
    margin-bottom: 30px;
  }

  .animation {
    position: absolute;
    top: 0;
    right: 0;
    height: 100%;
    width: auto;
    transform: translateX(100%);
    animation: moveRight 1s ease-out forwards;
  }

  .animation img {
    height: 100%;
    width: auto;
    object-fit: cover;
  }

  @keyframes moveRight {
    0% {
      transform: translateX(100%);
    }

    100% {
      transform: translateX(70%);
    }
  }

  .cta-btn {
    background-color: #00FF7F;
    color: #0A0A0A;
    padding: 10px 20px;
    text-decoration: none;
    font-weight: bold;
    border-radius: 5px;
  }

  .cta-btn:hover {
    background-color: #0BDA51;
    color: #FFF;
  }

  /* Стили для блока features */
  .features {
    display: flex;
    justify-content: space-around;
    padding: 50px 20px;
    text-align: center;
  }

  .feature {
    background-color: #1A1A1A;
    padding: 20px;
    border-radius: 10px;
    width: 30%;
  }

  .feature h2 {
    font-size: 24px;
    margin-bottom: 10px;
    color: #00FF7F;
  }

  .feature p {
    font-size: 16px;
    color: #FFF;
  }

  /* Стили для раздела услуги */
  section {
    padding: 50px 20px;
    text-align: center;
    background-color: rgba(26, 26, 26, 0.97);
    margin: 20px 0;
    border-radius: 10px;
  }

  section h2 {
    font-size: 40px;
    margin-bottom: 30px;
    color: #00FF7F;
  }

  section h3 {
    font-size: 30px;
    margin-bottom: 10px;
    color: #0BDA51;
  }

  section p {
    font-size: 18px;
    color: #FFF;
  }

  .service-options {
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
  }

  .service-option {
    background-color: #333;
    padding: 20px;
    margin: 10px;
    width: 250px;
    border-radius: 10px;
    text-align: center;
  }

  /* Стили для раздела цен */
  .pricing-options {
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
    margin-top: 20px;
  }

  .pricing-option {
    background-color: #333;
    padding: 20px;
    margin: 10px;
    width: 250px;
    border-radius: 10px;
    text-align: center;
  }

  .pricing-option h3 {
    font-size: 24px;
    margin-bottom: 10px;
    color: #00FF7F;
  }

  .pricing-option p {
    font-size: 16px;
    color: #FFF;
  }

  .pricing-option img {
    max-width: 200px;
    filter: invert(1) sepia(1) saturate(5000) hue-rotate(70deg);
  }

  /* Стили для контактов */
  .contact form {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .form-group {
    width: 100%;
    margin-bottom: 20px;
    text-align: left;
  }

  .form-group label {
    font-size: 18px;
    color: #FFF;
    margin-bottom: 5px;
  }

  .form-group input,
  .form-group textarea {
    background-color: #333;
    border: 1px solid #0BDA51;
    color: #FFF;
    padding: 10px;
    width: 100%;
    border-radius: 5px;
  }

  .contact p {
    margin-bottom: 20px;
  }

  .contact .form-container {
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
    box-sizing: border-box;
  }

  .contact form .form-group.name-email {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
  }

  .contact form .form-group.name-email .name,
  .contact form .form-group.name-email .email {
    flex: 1;
  }

  .contact form .form-group.name-email input {
    width: 100%;
  }

  .contact form .form-group textarea {
    width: 100%;
    min-height: 120px;
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

  /* Стили для футера */
  footer {
    background-color: #333;
    color: white;
    text-align: center;
    padding: 20px 0;
    margin-top: auto;
  }

  /* Основной блок */
  main {
    flex: 1;
  }

  /* Медиа-запросы для адаптации на мобильных устройствах */
  @media (max-width: 768px) {
    .hero {
      flex-direction: column;
      height: auto;
      padding: 20px;
    }

    .hero-content {
      max-width: 100%;
      text-align: center;
    }

    .animation {
      position: relative;
      transform: none;
      animation: none;
      width: 100%;
      height: auto;
    }

    .animation img {
      width: 100%;
      height: auto;
    }

    .service-options,
    .pricing-options {
      flex-direction: column;
      align-items: center;
    }

    .service-option,
    .pricing-option {
      width: 100%;
    }
  }

</style>
