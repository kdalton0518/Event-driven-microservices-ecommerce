import { ICustomerAuth } from "@/types/customer";
import { create } from "zustand";
import { createJSONStorage, devtools, persist } from "zustand/middleware";

type UserStore = {
  user: ICustomerAuth | null;
  setUser: (auth: ICustomerAuth) => void;
  logoutUser: () => void;
};

export const useUserStore = create<UserStore>()(
  devtools(
    persist(
      (set) => ({
        user: null,
        setUser: (user: ICustomerAuth) => {
          set(() => ({ user }));
        },
        logoutUser: () => {
          set(() => ({ user: null }));
        },
      }),
      {
        name: "user-storage",
        storage: createJSONStorage(() => localStorage),
      }
    )
  )
);
