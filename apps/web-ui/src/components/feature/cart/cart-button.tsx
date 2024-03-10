import { Button } from "@/components/ui/button";
import { ShoppingCart } from "lucide-react";

export function CartButton() {
  return (
    <Button variant="outline" size="icon">
      <ShoppingCart className="w-5" />
    </Button>
  );
}
