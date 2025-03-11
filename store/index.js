export const state = () => ({
  csrfToken: null,
});

export const mutations = {
  setCsrfToken(state, token) {
    state.csrfToken = token;
  }
};
