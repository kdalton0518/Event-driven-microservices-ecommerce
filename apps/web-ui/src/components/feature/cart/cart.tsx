"use client";

import { Cross1Icon } from "@radix-ui/react-icons";

import { useCartStore } from "@/store/cart";
import { currencyFormatter } from "@/utils/currency-formatter";
import Link from "next/link";

export function Cart() {
  const { cart, removeItem } = useCartStore();

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
              onClick={() => removeItem(item.id)}
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
                <span className="text-xl hover:underline">{item.name}</span>
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