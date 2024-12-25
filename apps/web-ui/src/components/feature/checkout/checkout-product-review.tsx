"use client";

import { Card } from "@/components/ui/card";
import { useCheckoutStore } from "@/store/checkout";
import { currencyFormatter } from "@/utils/currency-formatter";
import Image from "next/image";

export function CheckoutProductReview() {
  const { checkout } = useCheckoutStore();

  const total = checkout.reduce(
    (acc, crr) => acc + crr.price * crr.quantity,
    0
  );

  return (
    <Card className="border rounded-md p-4 space-y-4">
      <h2 className="text-xl">Products Review</h2>

      <div className="border-b border-zinc-200 dark:border-zinc-800" />

      {checkout.map((c) => (
        <div
          className="flex flex-col md:flex-row items-center justify-between"
          key={c.id}
        >
          <div className="flex justify-between md:justify-normal items-center w-full">
            {/* Add width and height properties to Image */}
            <Image
              src={c.image_url}
              alt={c.name}
              width={200}  // Adjust the width
              height={200} // Adjust the height
              className="w-20"
            />
            <div>
              <strong>{c.name}</strong> x {c.quantity}
            </div>
          </div>

          <span className="w-full text-right">
            {currencyFormatter.format(c.price / 100)}
          </span>
        </div>
      ))}

      <div className="border-b border-zinc-200 dark:border-zinc-800" />

      <div className="flex justify-between items-center">
        <span className="text-xl">Total (USD)</span>
        <span className="text-2xl font-bold">
          {currencyFormatter.format(total / 100)}
        </span>
      </div>
    </Card>
  );
}
