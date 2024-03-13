"use client";

import { PackageOpenIcon } from "lucide-react";
import Link from "next/link";

import { SignInButton } from "@/components/feature/auth/signin-header-button";
import { CartButton } from "@/components/feature/cart/cart-button";
import { OrderButton } from "@/components/feature/order/order-button";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { useUserStore } from "@/store/user";

export default function Header() {
  const { user } = useUserStore();

  const [first, last] = user?.customer.name.toUpperCase().split(" ") ?? ["U"];

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

        {user ? (
          <Avatar>
            <AvatarImage
              src="https://cdn-icons-png.flaticon.com/512/149/149071.png"
              alt={user.customer.name}
            />
            <AvatarFallback>
              {first[0]}
              {last ? last[0] : ""}
            </AvatarFallback>
          </Avatar>
        ) : (
          <Link href="/auth/signin">
            <SignInButton />
          </Link>
        )}
      </div>
    </div>
  );
}
