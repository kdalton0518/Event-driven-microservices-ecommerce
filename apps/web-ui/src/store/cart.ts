import { create } from "zustand";
import { createJSONStorage, devtools, persist } from "zustand/middleware";

import { IProduct } from "@/types/product";

type CartStore = {
  cart: IProduct[];
  addItem: (item: IProduct) => void;
  removeItem: (id: number) => void;
};

export const useCartStore = create<CartStore>()(
  devtools(
    persist(
      (set) => ({
        cart: [],
        addItem: (item) =>
          set((state) => {
            const itemIndex = state.cart.findIndex((i) => i.id === item.id);
            if (itemIndex !== -1) {
              state.cart[itemIndex].quantity += item.quantity;
              return { cart: state.cart };
            }

            return { cart: [...state.cart, item] };
          }),
        removeItem: (id) =>
          set((state) => ({
            cart: state.cart.filter((item) => item.id !== id),
          })),
      }),
      {
        name: "shopping-cart",
        storage: createJSONStorage(() => localStorage),
      }
    )
  )
);
