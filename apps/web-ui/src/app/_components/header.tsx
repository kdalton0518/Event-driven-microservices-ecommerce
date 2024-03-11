import Link from "next/link";
import { PackageOpenIcon } from "lucide-react";

import { CartButton } from "@/components/feature/cart/cart-button";
import { OrderButton } from "@/components/feature/order/order-button";
import { ThemeToggle } from "@/components/theme/theme-toggle";

export default function Header() {
  return (
    <div className="p-4 flex justify-between items-center border-b border-zinc-200 dark:border-zinc-900">
      <Link href="/" className="flex gap-2 items-center">
        <PackageOpenIcon />
        <h1 className="text-xl font-semibold">D-Commerce</h1>
      </Link>
      <div className="flex gap-2 items-center">
        <Link href="/order">
          <OrderButton />
        </Link>
        <Link href="/cart">
          <CartButton />
        </Link>
        <ThemeToggle />
      </div>
    </div>
  );
}
