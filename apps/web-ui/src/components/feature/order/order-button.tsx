import { Button } from "@/components/ui/button";
import { ShoppingBag } from "lucide-react";

export function OrderButton() {
  return (
    <Button variant="outline" size="icon">
      <ShoppingBag className="w-5" />
    </Button>
  );
}
