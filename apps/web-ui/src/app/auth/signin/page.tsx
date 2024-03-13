import { SigninForm } from "@/components/feature/auth/signin-form";

export default function Component() {
  return (
    <main className="mx-auto mt-32 max-w-[400px] space-y-6">
      <div className="space-y-2 text-center">
        <h1 className="text-3xl font-bold">Sign In</h1>
        <p className="text-gray-500 dark:text-gray-400">
          Enter your credentials to access your account
        </p>
      </div>

      <SigninForm />
    </main>
  );
}
