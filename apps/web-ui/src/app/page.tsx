import Link from "next/link";

import { getProductList } from "@/api/product";
import ProductCard from "@/components/feature/product/product-card";

export default async function Home() {
  const data = await getProductList({ page: 1, items: 10 });

  return (
    <main className="dark:bg-zinc-950 p-4 flex justify-center">
      <div className="max-w-fit min-h-fit grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 p-2 gap-6">
        {data.product_list.map((p) => (
          <Link href={`/product/${p.id}`} key={p.id}>
            <ProductCard {...p} />
          </Link>
        ))}
      </div>
    </main>
  );
}
