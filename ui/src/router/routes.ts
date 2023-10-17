import { RouteRecordRaw } from 'vue-router'

export abstract class RouteName {
  static get home() {
    return 'home'
  }

  static get notFound() {
    return 'notFound'
  }
}

export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: RouteName.home,
    component: () => import('@/pages/home'),
  },
  {
    path: '/:pathMatch(.*)*',
    name: RouteName.notFound,
    component: () => import('@/pages/not-found.vue'),
  },
]
