import Link from "next/link";

import { Button } from "@/components/ui/button";
import Image from "next/image";

export default async function Home() {
  return (
    <main className="flex flex-col justify-center items-center">
      <div className="px-10 py-28 flex flex-col md:flex-row justify-evenly gap-8 bg-gradient-to-r from-cyan-400 to-blue-400 dark:from-cyan-800 dark:to-blue-800 w-full">
        <div className="flex flex-col gap-8">
          <h1 className="text-3xl font-bold tracking-tighter sm:text-5xl xl:text-6xl/none">
            Shop the Latest Trends
          </h1>
          <p className="max-w-[400px] text-zinc-800 dark:text-zinc-300 md:text-xl">
            Discover our hand-picked selection of high-quality products.
          </p>
          <Link href="/product">
            <Button>View products</Button>
          </Link>
        </div>

        <Image src="/home.svg" alt="home" width={300} height={200} />
      </div>
    </main>
  );
}
