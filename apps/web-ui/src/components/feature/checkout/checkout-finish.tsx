"use client";

import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";

import { createOrder } from "@/api/order";
import { Button } from "@/components/ui/button";
import { ToastAction } from "@/components/ui/toast";
import { useToast } from "@/components/ui/use-toast";
import { useOrderStore } from "@/store/order";
import { IGetOrder, IOrder } from "@/types/order";

export function ProceedFinish() {
  const router = useRouter();
  const { toast } = useToast();
  const { order, clearOrder } = useOrderStore();

  const { isPending, mutateAsync, data } = useMutation({
    mutationFn: (data: IOrder) => createOrder(data),
    onSuccess: (data: IGetOrder) => {
      clearOrder();
      toast({
        title: "Successfully created order.",
        className:
          "bg-emerald-50 text-black dark:bg-transparent dark:text-white dark:border-emerald-950",
      });

      router.push(`/order/${data.id}`);
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
    <Button onClick={handleFinishCheckout} disabled={disable}>
      {isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
      Finish Checkout
    </Button>
  );
}
