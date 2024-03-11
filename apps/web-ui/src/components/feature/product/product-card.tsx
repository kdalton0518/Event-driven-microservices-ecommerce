"use client";

import { Card } from "@/components/ui/card";
import { IProduct } from "@/types/product";
import { currencyFormatter } from "@/utils/currency-formatter";

export default function ProductCard(props: IProduct) {
  return (
    <Card className="hover:shadow-md w-60 h-full">
      <div className="flex flex-col justify-center items-center gap-4 p-4">
        <img
          className="aspect-square object-cover border border-zinc-200 rounded-lg overflow-hidden hover:opacity-90 transition-opacity dark:border-zinc-800"
          alt={props.name}
          src={props.image_url}
          height={200}
          width={200}
        />

        <span className="flex items-center gap-2 font-semibold dark:font-semibold-contrast text-wrap text-left">
          {props.name}
        </span>
        <span className="text-lg text-left font-semibold">
          {currencyFormatter.format(props.price / 100)}
        </span>
      </div>
    </Card>
  );
}
