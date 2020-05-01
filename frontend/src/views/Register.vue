<template>
  <form class="card auth-card" @submit.prevent="submitHandler">
    <div class="card-content">
      <div class="row">
        <span class="card-title">3D Mesh computation</span>
      </div>

      <div class="input-field">
        <input id="username" type="text" v-model.trim="username" />
        <label for="username">Логин</label>
      </div>
      <div class="input-field">
        <input id="email" type="text" v-model.trim="email" />
        <label for="email">Email</label>
      </div>
      <div class="input-field">
        <input id="password" type="password" v-model.trim="password" />
        <label for="password">Пароль</label>
      </div>
    </div>

    <div class="card-action">
        <div>
            <button class="btn waves-effect waves-light auth-submit" type="submit">
                Зарегистрироваться
                <i class="material-icons right">send</i>
            </button>
        </div>
        <p class="center">
            Есть учетная запись?
        <router-link to="/login">Войти</router-link>
      </p>
    </div>
  </form>
</template>

<script>
import { required, minLength } from "vuelidate/lib/validators";

export default {
  name: "register",
  data() {
    return {
      username: "",
      email: "",
      password: ""
    };
  },
  validations: {
    username: { required },
    email: { required },
    password: { required }
  },
  methods: {
    async submitHandler() {
      try {
        if (this.$v.$invalid) {
          this.$v.$touch();
          return;
        }

        const { username, email, password } = this;
        await this.$store.dispatch("register", { username, email, password });

      } catch (e) {
        console.log("Register error was caused");
        console.log(e);
      }
    }
  }
};
</script>