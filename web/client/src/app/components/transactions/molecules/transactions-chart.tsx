import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  ChartLegend,
  ChartLegendContent,
  type ChartConfig,
} from "@/components/ui/chart";
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";
// import { getMonth, format } from "date-fns";

const TransactionsChart = () => {
  // const { data } = useGetTransactions();

  //TODO: Make this dynamic by creating a list of objects for each month with deposits = total and withdrawals = total

  const chartData = [
    { month: "January", deposit: 186, withdrawal: 80 },
    { month: "February", deposit: 305, withdrawal: 200 },
    { month: "March", deposit: 237, withdrawal: 120 },
    { month: "April", deposit: 73, withdrawal: 190 },
    { month: "May", deposit: 209, withdrawal: 130 },
    { month: "June", deposit: 214, withdrawal: 140 },
  ];

  const chartConfig = {
    deposit: {
      label: "Deposit",
      color: "hsl(var(--chart-2))",
    },
    withdrawal: {
      label: "Withdrawal",
      color: "hsl(var(--chart-1))",
    },
  } satisfies ChartConfig;

  return (
    <ChartContainer config={chartConfig} className="min-h-[200px] w-full">
      <BarChart accessibilityLayer data={chartData}>
        <CartesianGrid vertical={false} />

        <XAxis
          dataKey="month"
          tickLine={false}
          tickMargin={10}
          axisLine={false}
          tickFormatter={(value) => value.slice(0, 3)}
        />
        <ChartTooltip content={<ChartTooltipContent />} />
        <ChartLegend content={<ChartLegendContent />} />

        <Bar dataKey="deposit" fill="var(--color-deposit)" radius={4} />
        <Bar dataKey="withdrawal" fill="var(--color-withdrawal)" radius={4} />
      </BarChart>
    </ChartContainer>
  );
};

export default TransactionsChart;
