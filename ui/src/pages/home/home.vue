<template>
  <component :is='layout' />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useDisplay } from 'vuetify'

import DesktopLayout from './desktop-layout.vue'
import MobileLayout from './mobile-layout.vue'

const { name } = useDisplay()

const layout = computed(() => {
  switch (name.value) {
    case 'xs':
      return MobileLayout
  }
  return DesktopLayout
})
</script>