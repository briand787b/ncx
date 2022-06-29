import Home from './views/Home.vue'
// import Scan from './views/Scan.vue'
import NotFound from './views/NotFound.vue'

/** @type {import('vue-router').RouterOptions['routes']} */
export const routes = [
  { path: '/', component: Home, meta: { title: 'Home' } },
  // { path: '/scan', component: Scan, meta: { title: 'Scan' } },
  { path: '/:path(.*)', component: NotFound },
]
