import Home from './views/Home.vue'
import NotFound from './views/NotFound.vue'
import IDEHome from './views/IDEHome.vue'
import BusinessRulesIDE from './views/BusinessRulesIDE.vue'
import EntitiesIDE from './views/EntitiesIDE.vue'

/** @type {import('vue-router').RouterOptions['routes']} */
export const routes = [
  { path: '/', component: Home, meta: { title: 'Home' } },
  { path: '/ide', component: IDEHome, meta: { title: 'IDE'} },
  { path: '/ide/entities', component: EntitiesIDE, meta: { title: 'Entities'} },
  { path: '/ide/rules', component: BusinessRulesIDE, meta: { title: 'Business Rules'} },
  { path: '/:path(.*)', component: NotFound },
]
