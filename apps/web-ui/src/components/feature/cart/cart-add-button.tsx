"use client";

import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { useCartStore } from "@/store/cart";
import { IProduct } from "@/types/product";

export function CartAddButton(props: IProduct) {
  const { toast } = useToast();
  const { addItem } = useCartStore();

  return (
    <Button
      className="bg-zinc-500 dark:bg-zinc-700 text-white hover:bg-zinc-600 dark:hover:bg-zinc-800"
      size="lg"
      onClick={() => {
        addItem(props);
        toast({
          className:
            "bg-emerald-50 text-black dark:bg-transparent dark:text-white dark:border-emerald-950",
          title: "Item added to shopping cart",
        });
      }}
    >
      Add to cart
    </Button>
  );
}
