<template>
  <section id="ncx-entities" class="my-4 max-w-full">
    <header id="ncx-entities-header" class="lg:flex lg:justify-between max-w-full">
      <h3 class="sm:text-3xl">Entities Editor</h3>
      <button @click="addNewEntity" class="
              mt-4
              lg:mt-0
              mb-4
              w-20
              h-12
              rounded-md
              flex
              justify-center
              bg-green-600
              items-center
              hover:bg-green-400
              focus:outline-none focus:ring focus:ring-green-300
            ">
        <DocumentAddIcon class="h-8 w-8 text-slate-100" />
      </button>
    </header>
    <section>
      <ul>
        <li v-for="entity in flattenedEntityMap" :key="entity.id">
          <div class="w-full px-4">
            <div class="mx-auto w-full rounded-2xl bg-white p-2">
              <Disclosure v-slot="{ open }">
                <DisclosureButton class="
                        flex
                        w-full
                        justify-between
                        rounded-lg
                        bg-slate-100
                        px-4
                        py-2
                        text-left text-sm
                        font-medium
                        text-slate-900
                        hover:bg-slate-200
                        focus:outline-none
                        focus-visible:ring
                        focus-visible:ring-slate-500
                        focus-visible:ring-opacity-75
                      ">
                  <span>{{ entity.name }}</span>
                  <ChevronUpIcon :class="open ? 'rotate-180 transform' : ''" class="h-5 w-5 text-slate-500" />
                </DisclosureButton>
                <DisclosurePanel class="px-4 pt-4 pb-2 text-sm text-gray-500">

                </DisclosurePanel>
              </Disclosure>
            </div>
          </div>
        </li>
      </ul>
    </section>
    <footer>
      <a v-if="entities.length > 0" href="/ide#ncx-entities-header" class="focus:ring focus:ring-slate-700 text-sm">
        Back to top
      </a>
    </footer>
  </section>
</template>

<script>
import { DocumentAddIcon } from '@heroicons/vue/outline'
import { Disclosure, DisclosureButton, DisclosurePanel } from "@headlessui/vue";
import { ChevronUpIcon, PencilIcon } from "@heroicons/vue/solid";

export default {
  components: {
    ChevronUpIcon,
    DocumentAddIcon,
    Disclosure,
    DisclosureButton,
    DisclosurePanel,
    PencilIcon,
  },
  data() {
    return {
      entities: [],
    };
  },
  methods: {
    addNewEntity(e) {
      console.log("new-entity-generator clicked!");
    },
    editEntity(e) {
      console.log("entity editor clicked");
    },
  },
  computed: {
    flattenedEntityMap() {
      const flattenedEntityTree = {};
      const chainsaw = (woodbin, forest) => {
        forest.forEach(tree => {
          woodbin[tree.id] = {
            id: tree.id,
            conceptLabel: tree.conceptLabel,
            name: tree.name,
            initialValue: tree.initialValue,
          };

          if (tree?.attributes?.length > 0) {
            chainsaw(woodbin, tree.attributes);
          }
        });
      }

      // recursive arborist
      chainsaw(flattenedEntityTree, this.entities);
      return flattenedEntityTree;
    },
  },
  mounted() {
    // just load dummy data for now
    (this.entities = [
      {
        id: 1,
        conceptLabel: "Rental",
        name: "Rental",
        attributes: [
          {
            id: 2,
            conceptLabel: "Data Attribute",
            name: "Rental Period",
            type: "Date Range",
          },
          {
            id: 3,
            conceptLabel: "Data Attribute",
            name: "Rental Price Before Discount",
            type: "Amount",
            initialValue: {
              id: 4,
              conceptLabel: "Number",
              value: 0.0,
            },
          },
          {
            id: 5,
            conceptLabel: "Data Attribute",
            name: "Discount",
            type: "Percentage",
            initialValue: {
              id: 6,
              conceptLabel: "Number",
              value: 0,
            },
          },
          {
            id: 7,
            conceptLabel: "Data Attribute",
            name: "Rental Price After Discount",
            type: "Amount",
            initialValue: {
              id: 8,
              conceptLabel: "Attribute Reference",
              refID: 3,
            },
          },
        ],
      },
    ])
  },
}

</script>