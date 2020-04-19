<template>
  <ul class="sidenav app-sidenav" :class="{open: value}">

    <router-link
      v-for="link in links"
      :key="link.url"
      tag="li"
      active-class="active"
      :to="link.url"
      :exact="link.exact"
    >
      <a href="#" class="waves-effect waves-orange pointer">
        <i class="material-icons black-text">{{link.icon}}</i>
        {{link.title}}
      </a>
    </router-link>

  </ul>
</template>

<script>
export default {
  props: ['value'],
  data: () => ({
    links: [
      {title: 'Задачи', url: '/tasks', icon: 'assignment', exact: true},
      {title: 'Настройки', url: '/settings', icon: 'settings', exact: true}
    ]
  })
}
</script>



