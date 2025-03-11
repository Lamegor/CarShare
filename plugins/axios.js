export default function ({ $axios, store }) {
  $axios.onRequest((config) => {
    const csrfToken = store.state.csrfToken;
    if (csrfToken) {
      config.headers['X-Csrf-Token'] = csrfToken;
    }
  });
}
