import { Cart } from "@/components/feature/cart/cart";
import { ProceedCheckout } from "@/components/feature/checkout/checkout-proceed";

export default function Page() {
  return (
    <main className="p-10 space-y-4">
      <div className="flex flex-col gap-2 justify-between md:flex-row">
        <h1 className="text-2xl">Shopping Cart</h1>
        <ProceedCheckout />
      </div>
      <Cart />
    </main>
  );
}
