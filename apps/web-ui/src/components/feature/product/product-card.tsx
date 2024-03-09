import { currencyFormatter } from "@/utils/currency-formatter";
import Image from "next/image";

interface ProductCardProps {
  id: number;
  name: string;
  price: number;
  quantity: number;
  image_url: string;
}

export default function ProductCard(props: ProductCardProps) {
  return (
    <div className="flex flex-col bg-white dark:bg-zinc-900 border max-w-fit shadow-sm cursor-pointer">
      <Image
        className="p-4 pb-10"
        src={props.image_url}
        alt={props.name}
        width={250}
        height={250}
      />

      <div className="border-b border-zinc-200 dark:border-zinc-800" />

      <div className="flex flex-col p-4 pb-20">
        <span>{props.name}</span>
        <span className="text-xl">
          {currencyFormatter.format(props.price / 100)}
        </span>
      </div>
    </div>
  );
}
