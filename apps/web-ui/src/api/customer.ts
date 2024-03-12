import {
  ICreateCustomerIn,
  ILoginCustomerIn,
  ILoginCustomerOut,
} from "@/types/customer";

export async function registerCustomer(
  props: ICreateCustomerIn
): Promise<void> {
  const res = await fetch(`http://localhost:8080/api/auth/signup`, {
    method: "POST",
    body: JSON.stringify(props),
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to register customer");
  }
  return res.json();
}

export async function loginCustomer(
  props: ILoginCustomerIn
): Promise<ILoginCustomerOut> {
  const res = await fetch(`http://localhost:8080/api/auth/signin`, {
    method: "POST",
    body: JSON.stringify(props),
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to login customer");
  }

  return res.json();
}
