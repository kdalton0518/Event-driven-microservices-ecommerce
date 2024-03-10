import { IProduct } from "./product";

export interface IOrder {
  customer_id: string;
  product_list: IProduct[];
  payment_method: string;
}

export interface IGetOrder extends IOrder {
  id: string;
  status: string;
  total_price: number;
  created_at: Date;
  updated_at: Date;
}
