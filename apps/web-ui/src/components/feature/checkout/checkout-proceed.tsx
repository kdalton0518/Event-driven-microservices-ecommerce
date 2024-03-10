"use client";

import Link from "next/link";

import { Button } from "@/components/ui/button";
import { useCartStore } from "@/store/cart";
import { useCheckoutStore } from "@/store/checkout";
import { useOrderStore } from "@/store/order";

export function ProceedCheckout() {
  const { cart } = useCartStore();
  const { initCheckout } = useCheckoutStore();
  const { clearOrder, setProductList, setCustomerId } = useOrderStore();

  const handleProceedCheckout = () => {
    initCheckout(cart);
    clearOrder();
    setProductList(cart);
    setCustomerId("2f1134ee-7403-48da-b92c-6b3b8aac3708");
  };

  if (!cart.length) {
    return null;
  }

  return (
    <Link href="/checkout">
      <Button onClick={handleProceedCheckout}>Proceed to Checkout</Button>
    </Link>
  );
}
