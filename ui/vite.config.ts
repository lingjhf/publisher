import path from 'path'

import Vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import { defineConfig } from 'vite'
import vuetify from 'vite-plugin-vuetify'

function resolve(p: string) {
  return path.resolve(__dirname, p)
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [Vue(), UnoCSS({}), vuetify({ autoImport: true })],
  resolve: {
    alias: {
      '@': resolve('./src'),
    },
  },
  build: {
    target: 'esnext',
    rollupOptions: {
      output: {
        dir: '../static',
        manualChunks(id) {
          if (id.includes('vuetify')) {
            return 'vuetify'
          }
          if (id.includes('vue-router')) {
            return 'vue-router'
          }
          if (id.includes('vue')) {
            return 'vue'
          }
        }
      }
    }
  },
  server: {
    host: '0.0.0.0'
  }
})
