"use client";

import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react";
import Link from "next/link";

import { createOrder } from "@/api/order";
import { Button } from "@/components/ui/button";
import { ToastAction } from "@/components/ui/toast";
import { useToast } from "@/components/ui/use-toast";
import { useOrderStore } from "@/store/order";
import { IOrder } from "@/types/order";

export function ProceedFinish() {
  const { toast } = useToast();
  const { order, clearOrder } = useOrderStore();

  const { isPending, mutateAsync } = useMutation({
    mutationFn: (data: IOrder) => createOrder(data),
    onSuccess: () => {
      clearOrder();
      toast({
        title: "Successfully created order.",
        className:
          "bg-emerald-50 text-black dark:bg-transparent dark:text-white dark:border-emerald-950",
      });
    },
    onError: () => {
      toast({
        title: "Failed to create order.",
        className:
          "bg-red-50 text-black dark:bg-transparent dark:text-white dark:border-red-950",
        action: (
          <ToastAction id={"failed-create-order"} altText="Try again">
            <a href={"/"}>Back to home page</a>
          </ToastAction>
        ),
      });
    },
  });

  const handleFinishCheckout = () => mutateAsync(order);

  const disable =
    !order.customer_id || !order.payment_method || !order.product_list.length;

  if (disable) {
    return null;
  }

  return (
    <Link href="/checkout/completed">
      <Button onClick={handleFinishCheckout} disabled={disable}>
        {isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
        Finish Checkout
      </Button>
    </Link>
  );
}
