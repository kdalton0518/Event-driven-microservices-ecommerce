import ProductCard from "@/components/feature/product/product-card";
import { IProduct } from "@/types/product";

const mock: IProduct[] = [
  {
    id: 1,
    name: "iPhone 13 Pro",
    price: 99900,
    quantity: 50,
    image_url:
      "https://assets.nintendo.com/image/upload/f_auto/q_auto/dpr_2.0/c_scale,w_300/ncom/en_US/switch/refresh/switchindock",
  },
  {
    id: 2,
    name: "iPhone 13 Pro",
    price: 99900,
    quantity: 50,
    image_url:
      "https://assets.nintendo.com/image/upload/f_auto/q_auto/dpr_2.0/c_scale,w_300/ncom/en_US/switch/refresh/switchindock",
  },
  {
    id: 3,
    name: "iPhone 13 Pro",
    price: 99900,
    quantity: 50,
    image_url:
      "https://assets.nintendo.com/image/upload/f_auto/q_auto/dpr_2.0/c_scale,w_300/ncom/en_US/switch/refresh/switchindock",
  },
  {
    id: 4,
    name: "iPhone 13 Pro",
    price: 99900,
    quantity: 50,
    image_url:
      "https://assets.nintendo.com/image/upload/f_auto/q_auto/dpr_2.0/c_scale,w_300/ncom/en_US/switch/refresh/switchindock",
  },
  {
    id: 5,
    name: "iPhone 13 Pro",
    price: 99900,
    quantity: 50,
    image_url:
      "https://assets.nintendo.com/image/upload/f_auto/q_auto/dpr_2.0/c_scale,w_300/ncom/en_US/switch/refresh/switchindock",
  },
];

export default function Home() {
  return (
    <main className="w-screen h-screen bg-zinc-100 dark:bg-zinc-950 p-4">
      <div className="max-w-fit grid grid-cols-3 p-2 gap-4">
        {mock.map((p) => (
          <ProductCard {...p} />
        ))}
      </div>
    </main>
  );
}
