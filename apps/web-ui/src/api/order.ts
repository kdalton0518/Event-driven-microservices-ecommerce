import { IGetOrder, IOrder } from "@/types/order";

export async function createOrder(props: IOrder): Promise<IGetOrder> {
  const res = await fetch(`http://localhost:8080/api/orders`, {
    method: "POST",
    body: JSON.stringify(props),
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to create order");
  }
  return res.json();
}

export async function getOrder(id: string): Promise<IGetOrder> {
  const res = await fetch(`http://localhost:8080/api/orders/${id}`);

  if (!res.ok) {
    throw new Error("Failed to fetch order");
  }

  return res.json();
}
