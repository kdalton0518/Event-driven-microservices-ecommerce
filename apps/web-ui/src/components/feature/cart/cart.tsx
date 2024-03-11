"use client";

import { Cross1Icon } from "@radix-ui/react-icons";
import Link from "next/link";

import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { useCartStore } from "@/store/cart";
import { currencyFormatter } from "@/utils/currency-formatter";

export function Cart() {
  const { toast } = useToast();
  const { cart, removeItem } = useCartStore();

  if (!cart.length) {
    return (
      <div className="pt-10 space-y-8 flex flex-col">
        <span>
          It seems like you don't have any products on your shopping cart yet...
        </span>
        <Link href="/">
          <Button>Go back to Home page</Button>
        </Link>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {cart.map((item) => (
        <div
          key={item.id}
          className="flex justify-between items-center gap-4 border border-zinc-100 dark:border-zinc-900 p-4 rounded-md"
        >
          <div className="flex items-center gap-4">
            <Cross1Icon
              className="cursor-pointer"
              onClick={() => {
                removeItem(item.id);
                toast({
                  className:
                    "bg-emerald-50 text-black dark:bg-transparent dark:text-white dark:border-emerald-950",
                  title: "Item removed from cart!",
                });
              }}
            />
            <Link href={`/product/${item.id}`}>
              <img
                className="aspect-square object-cover border border-zinc-200 rounded-lg overflow-hidden hover:opacity-90 transition-opacity dark:border-zinc-800"
                alt={item.name}
                src={item.image_url}
                height={100}
                width={100}
              />
            </Link>

            <div className="flex flex-col">
              <Link href={`/product/${item.id}`}>
                <span className="text-base md:text-xl hover:underline">
                  {item.name}
                </span>
              </Link>
              <div>Quantity: {item.quantity}</div>
            </div>
          </div>

          <span className="text-xl">
            {currencyFormatter.format(item.price / 100)}
          </span>
        </div>
      ))}
    </div>
  );
}
