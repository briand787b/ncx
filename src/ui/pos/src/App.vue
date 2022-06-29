<template>
  <div class="h-screen">
    <header class="bg-white shadow md:h-1/5" v-if="$route.meta.title">
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bo ld leading-tight text-gray-900">
          POS
        </h1>
      </div>
    </header>
    <main class="md:h-4/5">
      <router-view />
    </main>
  </div>
</template>