"use client";

import { Button } from "@/components/ui/button";
import { useCartStore } from "@/store/cart";
import { IProduct } from "@/types/product";

export function CartAddButton(props: IProduct) {
  const { addItem } = useCartStore();
  const handleAdd = () => addItem(props);

  return (
    <Button
      className="bg-zinc-700 text-white hover:bg-zinc-800"
      size="lg"
      onClick={handleAdd}
    >
      Add to cart
    </Button>
  );
}
