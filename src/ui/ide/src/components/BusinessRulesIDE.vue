<template>
  <section id="ncx-business-rules" class="my-4 max-w-full">
    <header id="ncx-business-rules-header" class="lg:flex lg:justify-between max-w-full">
      <h3 class="sm:text-3xl">Business Rules Editor</h3>
      <button @click="addNewBizRule" class="
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
    <section class="">
      <ul>
        <li v-for="rule in bizRules" :key="rule.id">
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
                  <span>{{ rule.name }}</span>
                  <ChevronUpIcon :class="open ? 'rotate-180 transform' : ''" class="h-5 w-5 text-slate-500" />
                </DisclosureButton>
                <DisclosurePanel class="px-4 pt-4 pb-2 text-sm text-gray-500">
                  <div class="lg:flex lg:justify-between">
                    <p>When {{ rule.trigger }} {{ rule.action }}</p>
                    <button :rule-id="rule.id" @click="editBizRule(rule.id)" class="
                            my-2
                            p-2
                            hover:bg-slate-400
                            focus:outline-none focus:ring focus:ring-slate-300
                          ">
                      <PencilIcon class="h-5 w-5" />
                    </button>
                  </div>
                  <div v-if="rule.stateBeingEdited">
                    <div class="
                            max-w-fit
                            lg:flex lg:justify-start lg:items-baseline
                          ">
                      <div>
                        <label :for="rule.id">When</label>
                      </div>
                      <select :name="rule.trigger" :id="rule.id" class="mt-2 lg:mx-10">
                        <option value="" selected disabled hidden>
                          choose entity
                        </option>
                      </select>
                    </div>
                    <div class="
                            mt-4
                            lg:flex lg:justify-start lg:items-baseline
                          ">
                      <div>
                        <label :for="rule.id">Then</label>
                      </div>
                      <select :name="rule.action" :id="rule.id" class="mt-2 lg:mx-10">
                        <option value="" selected disabled hidden>
                          choose entity
                        </option>
                      </select>
                    </div>
                    <div class="mt-4 lg:flex lg:justify-end">
                      <button @click="editBizRule(rule.id)" class="
                              w-20
                              h-10
                              bg-red-600
                              text-white
                              ml-4
                              rounded-md
                            ">
                        Cancel
                      </button>
                      <button class="
                              w-20
                              h-10
                              bg-green-600
                              text-white
                              ml-4
                              rounded-md
                            ">
                        Save
                      </button>
                    </div>
                  </div>
                </DisclosurePanel>
              </Disclosure>
            </div>
          </div>
        </li>
      </ul>
    </section>
    <footer>
      <a v-if="bizRules.length > 0" href="/ide#ncx-business-rules-header"
        class="focus:ring focus:ring-slate-700 text-sm">Back to top</a>
    </footer>
  </section>
</template>

<script>
import { DocumentAddIcon } from "@heroicons/vue/outline";
import { Disclosure, DisclosureButton, DisclosurePanel } from "@headlessui/vue";
import { ChevronUpIcon, PencilIcon } from "@heroicons/vue/solid";

export default {
  components: {
    DocumentAddIcon,
    Disclosure,
    DisclosureButton,
    DisclosurePanel,
    ChevronUpIcon,
    PencilIcon,
  },
  data() {
    return {
      bizRules: [],
    };
  },
  methods: {
    addNewBizRule(e) {
      console.log("new-biz-rule-generator clicked!");
    },
    editBizRule(ruleID) {
      console.log("edit-biz-rule clicked!");
      const rule = this.bizRules.find((r) => r.id === ruleID);
      if (!rule) {
        return;
      }

      rule.stateBeingEdited = !rule.stateBeingEdited;
    },
  },
  mounted() {
    // just load dummy data for now
    (this.bizRules = [
      {
        id: 1,
        name: "Alert Manager When Cash Drawer Empty",
        trigger: "cash drawer is empty",
        action: "text 'drawer empty' to (705)-755-8822",
        stateBeingEdited: false,
      },
      {
        id: 2,
        name: "Page Attendant When SCO Line > 3 Groups",
        trigger: "sco line > 3 groups",
        action: "page @attendant",
        stateBeingEdited: false,
      },
      {
        id: 3,
        name: "Page Attendant When SCO Line > 3 Groups",
        trigger: "sco line > 3 groups",
        action: "page @attendant",
        stateBeingEdited: false,
      },
      {
        id: 4,
        name: "Page Supervisor When SCO Line > 5 Groups",
        trigger: "sco line > 5 groups",
        action: "page @supervisor",
        stateBeingEdited: false,
      },
      {
        id: 5,
        name: "Page Supervisor When SCO Down > 30 Minutes",
        trigger: "sco disabled > 30 minutes",
        action: "page @supervisor",
        stateBeingEdited: false,
      },
      {
        id: 6,
        name: "Page Manager When SCO Down > 1 Hour",
        trigger: "sco down > 1 hour",
        action: "page @manager",
        stateBeingEdited: false,
      },
      {
        id: 7,
        name: "Email District Manager When SCO Down > 24 Hours",
        trigger: "sco down > 24 hour",
        action: "email district_manager@store.com",
        stateBeingEdited: false,
      },
    ]);
  },
};
</script>