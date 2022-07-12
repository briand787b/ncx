<template>
  <div class="bg-gray-50 md:h-full">
    <div class="mx-auto max-w-screen-xl px-4 py-12 sm:px-6 
        lg:flex lg:justify-between lg:py-16 lg:px-8 border-solid border-2
      border-indigo-600 lg:h-full">
      <div class="w-full text-center border-2 border-solid 
        border-slate-500 lg:mr-1 lg:h-full lg:overflow-scroll">
        <h2>Items List</h2>
        <div class="grid grid-cols-1 md:grid-cols-2">
          <div v-for="item in forSaleItems" :key="item.id" @click="addToCart" :id="item.id" class="bg-gray-100 px-1 py-4 m-1 border-2 border-solid border-lime-500 
              hover:bg-sky-700 cursor-pointer active:bg-indigo-800">
            {{ item.displayName
            }}</div>
        </div>
      </div>
      <div class="w-full text-center border-2 border-solid border-slate-500 
          mt-5 lg:mt-0 lg:ml-1 lg:h-full md:flex md:flex-col md:justify-between">
        <h2 class="bg-sky-600 h-14 flex-none pt-2">Total</h2>
        <ul class="flex flex-col grow justify-start overflow-scroll">
          <li v-for="item in cartItems">{{ item.displayName }} {{ item.msrpCents }}</li>
        </ul>
        <div class="bg-indigo-500 py-2 mt-2 h-36 md:h-24 flex-none">
          <h3>Subtotal: {{ subtotal }}</h3>
          <div class="mt-2 flex flex-col items-center md:flex-row md:justify-center">
            <button class="bg-red-500 p-2 m-1 w-40" @click="cancel">Cancel</button>
            <button class="bg-lime-500 p-2 m-1 w-40">Pay</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    addToCart({ target: { id } }) {
      console.log('item ID: ', id);
      const item = this.forSaleItems.filter(fsItem => fsItem.id === id)[0]
      this.cartItems.push(item);
    },
    cancel() {
      this.cartItems = [];
    },
  },
  computed: {
    subtotal() {
      const subTCents = this.cartItems.reduce(((prev, curr) => curr.msrpCents + prev), 0);
      return "$" + (subTCents / 100).toFixed(2);
    },
  },
  mounted() {
    this.forSaleItems = [
      { id: "0", displayName: 'Canned Beans', msrpCents: 1000, },
      { id: "1", displayName: 'Tomatoes', msrpCents: 1000, },
      { id: "2", displayName: 'Tilapia', msrpCents: 1000, },
      { id: "3", displayName: 'Hot Sauce', msrpCents: 1000, },
      { id: "4", displayName: 'Sour Cream', msrpCents: 1000, },
      { id: "5", displayName: 'Ground Beef', msrpCents: 1000, },
      { id: "6", displayName: 'Shredded Chicken', msrpCents: 1000, },
      { id: "7", displayName: 'Cheese', msrpCents: 1000, },
      { id: "8", displayName: 'Jalapenos', msrpCents: 1000, },
      { id: "9", displayName: 'Onions', msrpCents: 1000, },
    ]
  },
  data() {
    return {
      cartItems: [],
      forSaleItems: [],
    };
  },
};
</script>