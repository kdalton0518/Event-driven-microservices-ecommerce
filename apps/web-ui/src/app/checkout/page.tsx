import { ProceedFinish } from "@/components/feature/checkout/checkout-finish";
import { CheckoutPaymentInformation } from "@/components/feature/checkout/checkout-payment-information";
import { CheckoutProductReview } from "@/components/feature/checkout/checkout-product-review";

export default function Page() {
  return (
    <main className="p-10 space-y-4">
      <div className="flex justify-between">
        <h1 className="text-2xl">Checkout</h1>
        <ProceedFinish />
      </div>
      <div className="w-full flex gap-2">
        <div className="w-1/2">
          <CheckoutPaymentInformation />
        </div>
        <div className="w-1/2">
          <CheckoutProductReview />
        </div>
      </div>
    </main>
  );
}
