import { getOrder } from "@/api/order";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { OrderStatus } from "@/utils/constant-mapper";

export default async function Page({ params }: { params: { id: string } }) {
  const order = await getOrder(params.id);

  return (
    <main className="p-4">
      <Card>
        <CardHeader>
          <CardTitle>Order Details</CardTitle>
          <CardDescription>
            Your order is not confirmed until payment is received.
          </CardDescription>
        </CardHeader>
        <CardContent className="grid gap-8">
          <div className="grid gap-2">
            <div className="text-sm">Order ID</div>
            <div className="font-medium">{order.id}</div>
          </div>
          <div className="grid gap-2">
            <div className="text-sm">Order Status</div>
            <div className="font-medium">{OrderStatus[order.status]}</div>
          </div>
          <div className="grid gap-2">
            <div className="text-sm">Order Date</div>
            <div>
              {new Date(order.created_at).toDateString()},{" "}
              {new Date(order.created_at).toLocaleTimeString()}
            </div>
          </div>
        </CardContent>
      </Card>
    </main>
  );
}
