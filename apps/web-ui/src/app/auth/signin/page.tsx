import Link from "next/link";

import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function Component() {
  return (
    <div className="m-auto max-w-[400px] space-y-6">
      <div className="space-y-2 text-center">
        <h1 className="text-3xl font-bold">Sign In</h1>
        <p className="text-gray-500 dark:text-gray-400">
          Enter your credentials to access your account
        </p>
      </div>
      <div className="space-y-2">
        <div className="space-y-2">
          <Label htmlFor="email">Email or Username</Label>
          <Input id="email" placeholder="m@example.com" required />
        </div>
        <div className="space-y-2">
          <div className="flex items-center">
            <Label htmlFor="password">Password</Label>
            <Link className="ml-auto inline-block text-sm underline" href="#">
              Forgot Password
            </Link>
          </div>
          <Input id="password" required type="password" />
        </div>
        <Button className="w-full">Sign In</Button>
      </div>
    </div>
  );
}
