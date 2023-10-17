// eslint-disable-next-line import/no-unresolved
import 'virtual:uno.css'
import 'vuetify/styles'
import { createApp } from 'vue'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import App from './App.vue'
import { router } from './router'

const vuetify = createVuetify()

createApp(App)
  .use(router)
  .use(vuetify)
  .mount('#app')