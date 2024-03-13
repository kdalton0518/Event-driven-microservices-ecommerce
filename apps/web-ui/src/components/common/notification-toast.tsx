import { ReactNode } from "react";

import { ToastAction } from "../ui/toast";

interface NotificationToastProps {
  id: string;
  altText: string;
  type: "success" | "error";
  children?: ReactNode;
}

export function NotificationToast({
  id,
  altText,
  type,
  children,
}: NotificationToastProps) {
  const className =
    type === "success"
      ? "border-0 bg-emerald-700 text-white hover:bg-emerald-800"
      : "border-0 bg-red-700 text-white hover:bg-red-800";

  return (
    <ToastAction id={id} altText={altText} className={className}>
      {children}
    </ToastAction>
  );
}
