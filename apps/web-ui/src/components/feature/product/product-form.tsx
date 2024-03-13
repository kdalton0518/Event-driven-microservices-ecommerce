"use client";

import Link from "next/link";
import { useState } from "react";

import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useCheckoutStore } from "@/store/checkout";
import { useOrderStore } from "@/store/order";
import { useUserStore } from "@/store/user";
import { IProduct } from "@/types/product";
import { CartAddButton } from "../cart/cart-add-button";

export function ProductForm(props: IProduct) {
  const { user } = useUserStore();
  const { clearOrder, setProductList, setCustomerId } = useOrderStore();
  const { initCheckout } = useCheckoutStore();
  const [quantity, setQuantity] = useState(1);

  const handleBuyProduct = () => {
    const product = [{ ...props, quantity }];
    initCheckout(product);
    clearOrder();
    setProductList(product);
    setCustomerId(user?.customer.id ?? "");
  };

  return (
    <form className="grid gap-4 md:gap-10" onSubmit={(e) => e.preventDefault()}>
      <div className="grid gap-2">
        <Label className="text-base" htmlFor="quantity">
          Select Quantity
        </Label>
        <Select defaultValue="1" onValueChange={(value) => setQuantity(+value)}>
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

      <div className="flex flex-col gap-2 w-full">
        <Link href={`/checkout`}>
          <Button className="w-full" onClick={handleBuyProduct}>
            Buy
          </Button>
        </Link>
        <CartAddButton
          id={props.id}
          name={props.name}
          description={props.description}
          price={props.price}
          quantity={quantity}
          image_url={props.image_url}
        />
      </div>
    </form>
  );
}
