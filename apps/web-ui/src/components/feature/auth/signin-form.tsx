"use client";

import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { FormEvent, useState } from "react";

import { loginCustomer } from "@/api/customer";
import { NotificationToast } from "@/components/common/notification-toast";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/components/ui/use-toast";
import { useUserStore } from "@/store/user";
import { ILoginCustomerIn, ILoginCustomerOut } from "@/types/customer";

export function SigninForm() {
  const router = useRouter();
  const { toast } = useToast();
  const { setUser } = useUserStore();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { isPending, mutateAsync, data } = useMutation({
    mutationFn: (data: ILoginCustomerIn) => loginCustomer(data),
    onSuccess: (data: ILoginCustomerOut) => {
      setUser(data);
      router.push("/");
    },
    onError: () => {
      toast({
        title: "Failed to login.",
        className: "bg-red-500 text-white",
        action: (
          <NotificationToast
            id="failed-login-customer"
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
    mutateAsync({ email, password });
  };

  return (
    <form className="space-y-2" onSubmit={handleSubmit}>
      <div className="space-y-2">
        <Label htmlFor="email">Email</Label>
        <Input
          id="email"
          placeholder="m@example.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
      </div>
      <div className="space-y-2">
        <Label htmlFor="password">Password</Label>
        <Input
          id="password"
          type="password"
          placeholder="************"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
      </div>

      <Button className="w-full" type="submit" disabled={isPending}>
        {isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
        Sign In
      </Button>

      <div className="flex items-center gap-2">
        <span>Not registered yet?</span>
        <Link className="inline-block text-sm underline" href="/auth/signup">
          Create an account.
        </Link>
      </div>
    </form>
  );
}
