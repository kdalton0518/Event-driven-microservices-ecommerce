import { IGetMeta } from "./common";

export interface IProduct {
  id: number;
  name: string;
  price: number;
  quantity: number;
  image_url: string;
}

export interface IGetProductListIn {
  page: number;
  items: number;
}

export interface IGetProductListOut {
  product_list: IProduct[];
  meta: IGetMeta;
}
