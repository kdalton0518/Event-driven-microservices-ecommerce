import { Button } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { ShoppingCart } from "lucide-react";

export function CartButton() {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button variant="outline" size="icon">
            <ShoppingCart className="w-5" />
          </Button>
        </TooltipTrigger>
        <TooltipContent>
          <p>Shopping Cart</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
