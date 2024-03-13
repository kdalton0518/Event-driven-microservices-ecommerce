"use client";

import Link from "next/link";

import { Button } from "@/components/ui/button";
import { useCartStore } from "@/store/cart";
import { useCheckoutStore } from "@/store/checkout";
import { useOrderStore } from "@/store/order";
import { useUserStore } from "@/store/user";

export function ProceedCheckout() {
  const { user } = useUserStore();
  const { cart } = useCartStore();
  const { initCheckout } = useCheckoutStore();
  const { clearOrder, setProductList, setCustomerId } = useOrderStore();

  const handleProceedCheckout = () => {
    initCheckout(cart);
    clearOrder();
    setProductList(cart);
    setCustomerId(user?.customer.id ?? "");
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
