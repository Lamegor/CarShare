export default async function ({ $axios, store }) {
  if (!store.state.csrfToken) {
    try {
      const response = await $axios.get('http://localhost:8080/auth/csrf-token'); // Запрос токена
      const csrfToken = response.headers['X-Csrf-Token']; // Получаем токен из заголовка

      if (csrfToken) {
        store.commit('setCsrfToken', csrfToken);
      }
    } catch (error) {
      console.error('Ошибка получения CSRF-токена:', error);
    }
  }
}
