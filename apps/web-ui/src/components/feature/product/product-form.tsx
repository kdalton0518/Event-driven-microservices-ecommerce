"use client";

import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { IProduct } from "@/types/product";
import { useState } from "react";
import { CartAddButton } from "../cart/cart-add-button";

export function ProductForm(props: IProduct) {
  const [quantity, setQuantity] = useState(1);

  const handleFieldValue = (value: string) => setQuantity(+value);

  return (
    <form className="grid gap-4 md:gap-10" onSubmit={(e) => e.preventDefault()}>
      <div className="grid gap-2">
        <Label className="text-base" htmlFor="quantity">
          Select Quantity
        </Label>
        <Select defaultValue="1" onValueChange={handleFieldValue}>
          <SelectTrigger className="w-24">
            <SelectValue placeholder="Select" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="1">1</SelectItem>
            <SelectItem value="2">2</SelectItem>
            <SelectItem value="3">3</SelectItem>
            <SelectItem value="4">4</SelectItem>
            <SelectItem value="5">5</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <div className="flex flex-col gap-2">
        <Button size="lg">Buy</Button>
        <CartAddButton
          id={props.id}
          name={props.name}
          price={props.price}
          quantity={quantity}
          image_url={props.image_url}
        />
      </div>
    </form>
  );
}
