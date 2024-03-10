"use client";

import { useState } from "react";

import { Card } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useOrderStore } from "@/store/order";
import { CheckoutCreditCard } from "./checkout-credit-card";

const paymentMethodList = ["PIX", "CREDIT_CARD"];

export function CheckoutPaymentInformation() {
  const { setPaymentMethod } = useOrderStore();
  const [paymentMethod, setMethod] = useState("");

  return (
    <Card className="border rounded-md p-4">
      <form className="space-y-4">
        <div className="space-y-4">
          <Label className="text-base" htmlFor="quantity">
            Select payment method
          </Label>
          <Select
            defaultValue="1"
            onValueChange={(value) => {
              const payment = paymentMethodList[Number(value) - 1];
              setMethod(payment);
              setPaymentMethod(payment);
            }}
          >
            <SelectTrigger className="w-36">
              <SelectValue placeholder="Select" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="1">Pix</SelectItem>
              <SelectItem value="2">Credit Card</SelectItem>
            </SelectContent>
          </Select>
        </div>

        {paymentMethod === "CREDIT_CARD" && <CheckoutCreditCard />}
      </form>
    </Card>
  );
}
