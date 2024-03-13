export interface ICustomer {
  id: string;
  name: string;
  email: string;
  password: string;
}

export interface ICreateCustomerIn {
  name: string;
  email: string;
  password: string;
}

export interface ILoginCustomerIn {
  email: string;
  password: string;
}

export interface ILoginCustomerOut {
  access_token: string;
  customer: Omit<ICustomer, "password">;
}

export interface ICustomerAuth extends ILoginCustomerOut {}
