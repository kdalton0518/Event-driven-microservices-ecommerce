"use client";

import Link from "next/link";

import { Card } from "@/components/ui/card";
import { IProduct } from "@/types/product";
import { currencyFormatter } from "@/utils/currency-formatter";

export default function ProductCard(props: IProduct) {
  return (
    <Card>
      <div className="grid gap-4 text-center p-4">
        <Link className="inline-block" href={`/product/${props.id}`}>
          <img
            alt="Thumbnail"
            className="aspect-square object-cover border border-gray-200 rounded-lg overflow-hidden hover:opacity-90 transition-opacity dark:border-gray-800"
            src={props.image_url}
            height={200}
            width={200}
          />
          <span className="sr-only">View Product</span>
        </Link>
        <div className="flex items-center gap-2">
          <Link
            className="font-semibold hover:underline dark:font-semibold-contrast"
            href="#"
          >
            {props.name}
          </Link>
        </div>
        <div className="text-sm font-semibold">
          {currencyFormatter.format(props.price / 100)}
        </div>
      </div>
    </Card>
  );
}
