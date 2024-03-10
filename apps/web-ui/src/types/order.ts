import { IProduct } from "./product";

export interface IOrder {
  customer_id: string;
  product_list: IProduct[];
  payment_method: string;
}
