import {
  IGetProductListIn,
  IGetProductListOut,
  IProduct,
} from "@/types/product";

export async function getProductList({
  page,
  items,
}: IGetProductListIn): Promise<IGetProductListOut> {
  const res = await fetch(
    `http://localhost:8080/api/products?page=${page}&items=${items}`
  );

  if (!res.ok) {
    throw new Error("Failed to fetch product list");
  }

  return res.json();
}

export async function getProduct(id: number): Promise<IProduct> {
  const res = await fetch(`http://localhost:8080/api/products/${id}`);

  if (!res.ok) {
    throw new Error("Failed to fetch product");
  }

  return res.json();
}
