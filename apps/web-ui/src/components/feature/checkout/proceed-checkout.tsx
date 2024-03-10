"use client";

import Link from "next/link";

import { Button } from "@/components/ui/button";
import { useCartStore } from "@/store/cart";
import { useCheckoutStore } from "@/store/checkout";

export function ProceedCheckout() {
  const { cart } = useCartStore();
  const { initCheckout } = useCheckoutStore();

  return (
    <Link href="/checkout">
      <Button onClick={() => initCheckout(cart)}>Proceed to Checkout</Button>
    </Link>
  );
}
