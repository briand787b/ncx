<template>
  <div class="bg-gray-50">
    <div class="mx-auto max-w-screen-xl px-4 py-12 sm:px-6">
      <h2 class="
          text-3xl
          font-extrabold
          leading-9
          tracking-tight
          text-gray-900
          sm:text-4xl sm:leading-10
        ">
        NCX
      </h2>
      <p><a href="/ide/entities">Edit Entities</a></p>
      <p><a href="/ide/rules">Edit Business Rules</a></p>
    </div>
  </div>
</template>

<script>
import { DocumentAddIcon } from "@heroicons/vue/outline";
import EntitiesIDE from "@/components/EntitiesIDE.vue";
import BusinessRulesIDE from "@/components/BusinessRulesIDE.vue";
import { Disclosure, DisclosureButton, DisclosurePanel } from "@headlessui/vue";
import { ChevronUpIcon, PencilIcon } from "@heroicons/vue/solid";

export default {
  components: {
    BusinessRulesIDE,
    EntitiesIDE,
    DocumentAddIcon,
    Disclosure,
    DisclosureButton,
    DisclosurePanel,
    ChevronUpIcon,
    PencilIcon,
  },
};
</script>