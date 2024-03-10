import Link from "next/link";

import { CartButton } from "@/components/feature/cart/cart-button";
import { ThemeToggle } from "@/components/theme/theme-toggle";

export default function Header() {
  return (
    <div className="p-4 flex justify-between items-center border-b border-zinc-200 dark:border-zinc-900">
      <Link href="/">
        <h1 className="text-2xl font-semibold">D-Commerce</h1>
      </Link>
      <div className="flex gap-2 items-center">
        <Link href="/cart">
          <CartButton />
        </Link>
        <ThemeToggle />
      </div>
    </div>
  );
}
