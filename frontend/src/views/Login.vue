<template>
  <form class="card auth-card" @submit.prevent="submitHandler">
    <div class="card-content">
      <div class="row">
        <span class="card-title">3D Mesh computation</span>
      </div>
      <div class="input-field">
        <input
          id="email"
          type="text"
          v-model.trim="email"
          :class="{invalid: ($v.email.$dirty && !$v.email.required)}"
        />
        <label for="email">Email</label>
        <small
          class="helper-text invalid"
          v-if="$v.email.$dirty && !$v.email.required"
        >Поле Email не должно быть пустым</small>
      </div>
      <div class="input-field">
        <input
          id="password"
          type="password"
          v-model.trim="password"
          :class="{invalid: ($v.password.$dirty && !$v.password.required) || ($v.password.$dirty && !$v.password.minLength)}"
        />
        <label for="password">Пароль</label>
        <small
          class="helper-text invalid"
          v-if="$v.password.$dirty && !$v.password.required"
        >Поле Пароль не должно быть пустым</small>
      </div>
    </div>
    <div class="card-action">
      <div>
        <button class="btn waves-effect waves-light auth-submit" type="submit">
          Войти
          <i class="material-icons right">send</i>
        </button>
        <small v-if="loginError" class="helper-text invalid">Неверный логин или пароль</small>
      </div>
      <p class="center">
        Нет учетной записи?
        <router-link to="/register">Зарегистрироваться</router-link>
      </p>
    </div>
  </form>
</template>
<script>
import { required, minLength } from "vuelidate/lib/validators";

export default {
  name: "login",
  data() {
    return {
      email: "",
      password: ""
    };
  },
  validations: {
    email: { required },
    password: { required }
  },
  computed: {
    loginError() {
      return this.$store.getters.authStatus === "error";
    }
  },
  methods: {
    async submitHandler() {
      try {
        if (this.$v.$invalid) {
          this.$v.$touch();
          return;
        }

        const { email, password } = this;
        await this.$store.dispatch("login", { email, password });
        this.$router.push("/tasks");

      } catch(e) {
          console.log("Login error was caused")
          console.log(e);
      }
    }
  }
};
</script>