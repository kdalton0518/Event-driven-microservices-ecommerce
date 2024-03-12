"use client";

import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

import { registerCustomer } from "@/api/customer";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ToastAction } from "@/components/ui/toast";
import { useToast } from "@/components/ui/use-toast";
import { ICreateCustomerIn } from "@/types/customer";

export default function Component() {
  const router = useRouter();
  const { toast } = useToast();

  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { isPending, mutateAsync, data } = useMutation({
    mutationFn: (data: ICreateCustomerIn) => registerCustomer(data),
    onSuccess: () => {
      toast({
        title: "Successfully signed up",
        className:
          "bg-emerald-50 text-black dark:bg-transparent dark:text-white dark:border-emerald-950",
      });
      router.push("/auth/signin");
    },
    onError: () => {
      toast({
        title: "Failed to login.",
        className:
          "bg-red-50 text-black dark:bg-transparent dark:text-white dark:border-red-950",
        action: (
          <ToastAction id={"failed-login-customer"} altText="Try again">
            <span onClick={() => router.refresh()}>Try again</span>
          </ToastAction>
        ),
      });
    },
  });

  const handleSubmit = () => mutateAsync({ name, email, password });

  return (
    <div className="mx-auto mt-32 max-w-[400px] space-y-6">
      <div className="space-y-2 text-center">
        <h1 className="text-3xl font-bold">Sign Up</h1>
        <p className="text-gray-500 dark:text-gray-400">
          Enter your information to create an account
        </p>
      </div>
      <div className="space-y-2">
        <div className="space-y-2">
          <Label htmlFor="singup-name">Name</Label>
          <Input
            id="singup-name"
            placeholder="John Doe"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="signup-email">Email</Label>
          <Input
            id="signup-email"
            placeholder="m@example.com"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="signup-password">Password</Label>
          <Input
            id="signup-password"
            type="password"
            placeholder="************"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <Button
          className="w-full"
          type="submit"
          disabled={isPending}
          onClick={handleSubmit}
        >
          {isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          Sign Up
        </Button>

        <div className="flex items-center gap-2">
          <span>Already registered?</span>
          <Link className="inline-block text-sm underline" href="/auth/signin">
            Sign In with your account
          </Link>
        </div>
      </div>
    </div>
  );
}
