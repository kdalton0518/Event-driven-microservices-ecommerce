import { getProduct } from "@/api/product";
import { ProductForm } from "@/components/feature/product/product-form";
import { Separator } from "@/components/ui/separator";
import { currencyFormatter } from "@/utils/currency-formatter";

export default async function Page({ params }: { params: { id: string } }) {
  const data = await getProduct(+params.id);

  return (
    <div className="grid md:grid-cols-2 items-start max-w-6xl px-4 mx-auto py-6 gap-6 lg:gap-12">
      <div className="grid gap-4 md:gap-10 items-start">
        <div className="flex flex-col gap-4 md:flex items-start">
          <h1 className="text-zinc-500 dark:text-zinc-300 font-bold text-3xl lg:text-4xl">
            {data.name}
          </h1>
          <span className="text-4xl font-bold">
            {currencyFormatter.format(data.price / 100)}
          </span>
        </div>
        <span>Quantity: {data.quantity}</span>

        <ProductForm {...data} />

        <Separator />
        <div className="grid gap-4 text-sm leading-loose">
          <p>{data.description}</p>
        </div>
      </div>
      <div className="grid gap-3 items-start">
        <img
          className="aspect-square object-cover border border-zinc-200 w-full rounded-lg overflow-hidden dark:border-zinc-800"
          src={data.image_url}
          alt={data.name}
          height={600}
          width={600}
        />
      </div>
    </div>
  );
}
