import Transaction from "@/types/transactions";
import { Row } from "@tanstack/react-table";
import { z } from "zod";

const dateRangeFilter = (
  row: Row<Transaction>,
  columnId: string,
  filterValue: string[]
) => {
  const rowValue = row.getValue(columnId);

  return !!filterValue.find((x) => x === rowValue);
};

export default dateRangeFilter;

export const transactionDrawerSchema = z.object({
  userId: z.string().optional(),
  name: z.string(),
  amount: z.string(),
  currency: z.string(),
  type: z.string(),
  bank: z.string(),
});

export type Inputs = z.infer<typeof transactionDrawerSchema>;

export const fieldsArray = [
  "userId",
  "name",
  "amount",
  "currency",
  "bank",
] as const;
