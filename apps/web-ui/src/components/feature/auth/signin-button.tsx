import { Button } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { LogInIcon } from "lucide-react";

export function SignInButton() {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button variant="outline" size="icon">
            <LogInIcon className="w-5" />
          </Button>
        </TooltipTrigger>
        <TooltipContent>
          <p>Sign In</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
