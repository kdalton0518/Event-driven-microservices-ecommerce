"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { ProceedFinish } from "@/components/feature/checkout/checkout-finish";
import { CheckoutPaymentInformation } from "@/components/feature/checkout/checkout-payment-information";
import { CheckoutProductReview } from "@/components/feature/checkout/checkout-product-review";

const queryClient = new QueryClient();

export default function Page() {
  return (
    <main className="p-10 space-y-4">
      <QueryClientProvider client={queryClient}>
        <div className="flex flex-col md:flex-row justify-between">
          <h1 className="text-2xl">Checkout</h1>
          <ProceedFinish />
        </div>
        <div className="w-full flex flex-col-reverse md:flex-row gap-2">
          <div className="w-full md:w-1/2">
            <CheckoutPaymentInformation />
          </div>
          <div className="w-full md:w-1/2">
            <CheckoutProductReview />
          </div>
        </div>
      </QueryClientProvider>
    </main>
  );
}
