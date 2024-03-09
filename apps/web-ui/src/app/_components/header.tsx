import { ThemeToggle } from "@/components/theme/theme-toggle";
import Link from "next/link";

export default function Header() {
  return (
    <div className="p-4 flex justify-between items-center border-b border-zinc-200 dark:border-zinc-900">
      <Link href="/">
        <h1 className="text-xl">D-Commerce</h1>
      </Link>
      <ThemeToggle />
    </div>
  );
}
