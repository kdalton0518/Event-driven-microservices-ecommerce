"use client";

import { useState } from "react";

import { Card } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useOrderStore } from "@/store/order";
import { CheckoutCreditCard } from "./checkout-credit-card";

export function CheckoutPaymentInformation() {
  const { setPaymentMethod, order } = useOrderStore();
  const [paymentMethod, setMethod] = useState(order.payment_method);

  return (
    <Card className="border rounded-md p-4">
      <form className="space-y-4">
        <div className="space-y-4">
          <Label className="text-base" htmlFor="quantity">
            Select payment method
          </Label>
          <Select
            defaultValue={paymentMethod}
            onValueChange={(value) => {
              setMethod(value);
              setPaymentMethod(value);
            }}
          >
            <SelectTrigger className="w-36">
              <SelectValue placeholder="Select method" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem value="PIX">Pix</SelectItem>
                <SelectItem value="CREDIT_CARD">Credit Card</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>

        {paymentMethod === "CREDIT_CARD" && <CheckoutCreditCard />}
      </form>
    </Card>
  );
}
