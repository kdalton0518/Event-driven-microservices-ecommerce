import { getProduct } from "@/api/product";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
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
        <form className="grid gap-4 md:gap-10">
          <div className="grid gap-2">
            <Label className="text-base" htmlFor="quantity">
              Select Quantity
            </Label>
            <Select defaultValue="1">
              <SelectTrigger className="w-24">
                <SelectValue placeholder="Select" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="1">1</SelectItem>
                <SelectItem value="2">2</SelectItem>
                <SelectItem value="3">3</SelectItem>
                <SelectItem value="4">4</SelectItem>
                <SelectItem value="5">5</SelectItem>
              </SelectContent>
            </Select>
          </div>
          <div className="flex flex-col gap-2">
            <Button size="lg">Buy</Button>
            <Button
              className="bg-zinc-700 text-white hover:bg-zinc-800"
              size="lg"
            >
              Add to cart
            </Button>
          </div>
        </form>
        <Separator />
        <div className="grid gap-4 text-sm leading-loose">
          <p>
            Introducing the Acme Prism T-Shirt, a perfect blend of style and
            comfort for the modern individual. This tee is crafted with a
            meticulous composition of 60% combed ringspun cotton and 40%
            polyester jersey, ensuring a soft and breathable fabric that feels
            gentle against the skin.
          </p>
        </div>
      </div>
      <div className="grid gap-3 items-start">
        <img
          className="aspect-square object-cover border border-gray-200 w-full rounded-lg overflow-hidden dark:border-gray-800"
          src={data.image_url}
          alt={data.name}
          height={600}
          width={600}
        />
      </div>
    </div>
  );
}
