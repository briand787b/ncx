<template>
  <div>
    <header class="bg-white shadow" v-if="$route.meta.title">
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <a href="/">
          <h1 class="text-3xl font-bold leading-tight text-gray-900">
            NCR Cloud
          </h1>
        </a>
      </div>
    </header>
    <main>
      <router-view />
    </main>
  </div>
</template>

<script>
import { Menu, MenuButton, MenuItems, MenuItem } from "@headlessui/vue";

export default {};
</script>