import { create } from "zustand";
import { createJSONStorage, devtools, persist } from "zustand/middleware";

import { IProduct } from "@/types/product";

type CheckoutStore = {
  checkout: IProduct[];
  initCheckout: (cart: IProduct[]) => void;
  clearCheckout: () => void;
};

export const useCheckoutStore = create<CheckoutStore>()(
  devtools(
    persist(
      (set) => ({
        checkout: [],
        initCheckout: (cart) => set((state) => ({ checkout: cart })),
        clearCheckout: () => set((state) => ({ checkout: [] })),
      }),
      {
        name: "checkout",
        storage: createJSONStorage(() => localStorage),
      }
    )
  )
);
