<template>
  <nav class="navbar orange lighten-1">
    <div class="nav-wrapper">
      <div class="navbar-left">
        <a href="#" @click.prevent="$emit('click')">
          <i class="material-icons black-text">dehaze</i>
        </a>
        <span class="black-text"></span>
      </div>

      <ul class="right hide-on-small-and-down">
        <li>
          <i class="material-icons">account_box</i>
        </li>
        <li>
          <a class="dropdown-trigger black-text" href="#" data-target="dropdown" ref="dropdown">
            {{userName}}
            <i class="material-icons right">arrow_drop_down</i>
          </a>

          <ul id="dropdown" class="dropdown-content">
            <li>
              <a href="#" class="black-text" @click.prevent="logout">
                <i class="material-icons">assignment_return</i>Выйти
              </a>
            </li>
          </ul>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  data: () => ({
    date: new Date(),
    interval: null,
    dropdown: null
  }),
  methods: {
    logout() {
      this.$store
        .dispatch("logout")
        .then(() => {
          this.$router.push("/login");
        })
        .catch(err => {
          console.log(err);
        });
    }
  },
  computed: {
    userName() {
      return this.$store.getters.userName;
    }
  },
  mounted() {
    this.interval = setInterval(() => {
      this.date = new Date();
    }, 1000);
    this.dropdown = M.Dropdown.init(this.$refs.dropdown, {
      constrainWidth: false
    });
  },
  beforeDestroy() {
    clearInterval(this.interval);
    if (this.dropdown && this.dropdown.destroy) {
      this.dropdown.destroy();
    }
  }
};
</script>

<style lang="scss" scoped>
.dropdown-content {
  top: 100% !important;
}
</style>
