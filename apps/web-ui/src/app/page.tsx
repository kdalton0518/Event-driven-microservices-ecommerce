import { getProductList } from "@/api/product";
import ProductCard from "@/components/feature/product/product-card";
import Link from "next/link";

export default async function Home() {
  const data = await getProductList({ page: 1, items: 10 });

  return (
    <main className="bg-zinc-100 dark:bg-zinc-950 p-4 flex justify-center">
      <div className="max-w-fit min-h-fit grid grid-cols-3 p-2 gap-4">
        {data.product_list.map((p) => (
          <Link href={`/product/${p.id}`} key={p.id}>
            <ProductCard {...p} />
          </Link>
        ))}
      </div>
    </main>
  );
}
