"use client";

import { ChangeEvent, useState } from "react";

import { Input } from "@/components/ui/input";

export function CheckoutCreditCard() {
  const [cardNumber, setCardNumber] = useState("");
  const [cardValidDate, setCardValidDate] = useState("");
  const [cardCVV, setCardCVV] = useState("");

  console.log(cardValidDate);

  const handleCardNumber = (e: ChangeEvent<HTMLInputElement>) => {
    const sanitizedValue = e.target.value.replace(/\D/g, "");
    const maskedValue = sanitizedValue.replace(
      /^(\d{0,4})(\d{0,4})(\d{0,4})(\d{0,4}).*/,
      (match: string, p1: string, p2: string, p3: string, p4: string) => {
        return [p1, p2, p3, p4].filter((group) => !!group).join(" ");
      }
    );
    setCardNumber(maskedValue);
  };

  const handleCardValidDate = (e: ChangeEvent<HTMLInputElement>) => {
    let { value } = e.target;
    value = value.replace(/\D/g, "");
    let maskedValue = "";

    for (let i = 0; i < value.length; i++) {
      if (i === 2) {
        maskedValue += "/";
      }
      if (i === 4) {
        break;
      }
      maskedValue += value[i];
    }

    setCardValidDate(maskedValue);
  };

  const handleCardCvv = (e: ChangeEvent<HTMLInputElement>) => {
    let { value } = e.target;
    value = value.replace(/\D/g, "");
    value = value.slice(0, 3);
    setCardCVV(value);
  };

  return (
    <div className="flex gap-1">
      <Input
        className="w-2/3"
        type="text"
        max={16}
        value={cardNumber}
        onChange={handleCardNumber}
        placeholder="xxxx xxxx xxxx xxxx"
      />
      <Input
        className="w-1/3"
        type="text"
        max={5}
        value={cardValidDate}
        onChange={handleCardValidDate}
        placeholder="MM/YY"
      />
      <Input
        className="w-1/3"
        type="text"
        max={3}
        value={cardCVV}
        onChange={handleCardCvv}
        placeholder="CVV"
      />
    </div>
  );
}
