import { create } from "zustand";
import { createJSONStorage, devtools, persist } from "zustand/middleware";

import { IOrder } from "@/types/order";
import { IProduct } from "@/types/product";

type OrderStore = {
  order: IOrder;
  setCustomerId: (id: string) => void;
  setPaymentMethod: (paymentMethod: string) => void;
  setProductList: (productList: IProduct[]) => void;
  clearOrder: () => void;
};

export const useOrderStore = create<OrderStore>()(
  devtools(
    persist(
      (set) => ({
        order: {
          customer_id: "",
          payment_method: "",
          product_list: [],
        },
        setCustomerId: (id) =>
          set((state) => ({
            order: { ...state.order, customer_id: id },
          })),
        setPaymentMethod: (paymentMethod) =>
          set((state) => ({
            order: { ...state.order, payment_method: paymentMethod },
          })),
        setProductList: (productList) =>
          set((state) => ({
            order: { ...state.order, product_list: productList },
          })),
        clearOrder: () =>
          set(() => ({
            order: { customer_id: "", payment_method: "", product_list: [] },
          })),
      }),
      {
        name: "order",
        storage: createJSONStorage(() => localStorage),
      }
    )
  )
);
