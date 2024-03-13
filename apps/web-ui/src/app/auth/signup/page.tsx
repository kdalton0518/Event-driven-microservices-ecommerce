import { SignupForm } from "@/components/feature/auth/signup-form";

export default function Component() {
  return (
    <div className="mx-auto mt-32 max-w-[400px] space-y-6">
      <div className="space-y-2 text-center">
        <h1 className="text-3xl font-bold">Sign Up</h1>
        <p className="text-gray-500 dark:text-gray-400">
          Enter your information to create an account
        </p>
      </div>
      <SignupForm />
    </div>
  );
}
