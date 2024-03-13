"use client";

import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { FormEvent, useState } from "react";

import { registerCustomer } from "@/api/customer";
import { NotificationToast } from "@/components/common/notification-toast";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/components/ui/use-toast";
import { ICreateCustomerIn } from "@/types/customer";

export function SignupForm() {
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
        action: (
          <NotificationToast
            id="success-register-customer"
            altText="Signup success"
            type="success"
            children={null}
          />
        ),
      });
      router.push("/auth/signin");
    },
    onError: () => {
      toast({
        title: "Failed to sign up.",
        action: (
          <NotificationToast
            id="failed-register-customer"
            altText="Try again"
            type="error"
            children={<span onClick={() => router.refresh()}>Try again</span>}
          />
        ),
      });
    },
  });

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    mutateAsync({ name, email, password });
  };

  return (
    <form className="space-y-2" onSubmit={handleSubmit}>
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
      <Button className="w-full" type="submit" disabled={isPending}>
        {isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
        Sign Up
      </Button>

      <div className="flex items-center gap-2">
        <span>Already registered?</span>
        <Link className="inline-block text-sm underline" href="/auth/signin">
          Sign In with your account
        </Link>
      </div>
    </form>
  );
}
