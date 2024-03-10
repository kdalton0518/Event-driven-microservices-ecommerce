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
    setCustomerId("haha");
  };

  return (
    <Link href="/checkout">
      <Button onClick={handleProceedCheckout}>Proceed to Checkout</Button>
    </Link>
  );
}
