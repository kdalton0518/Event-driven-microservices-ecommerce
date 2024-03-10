"use client";

import Link from "next/link";

import { Button } from "@/components/ui/button";
import { useOrderStore } from "@/store/order";

export function ProceedFinish() {
  const { order } = useOrderStore();

  const handleFinishCheckout = () => {};

  const disable =
    !order.customer_id || !order.payment_method || !order.product_list.length;

  if (disable) {
    return null;
  }

  return (
    <Link href="/checkout/completed">
      <Button onClick={handleFinishCheckout} disabled={disable}>
        Finish Checkout
      </Button>
    </Link>
  );
}
