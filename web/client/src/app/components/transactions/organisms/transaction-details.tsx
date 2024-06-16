import { Button } from "@/components/ui/button";
import { useLogoutMutation } from "@/app/hooks/requests";
import { LogOut } from "lucide-react";
import { useState } from "react";
import { Table } from "@tanstack/react-table";
import Transaction from "@/types/transactions";
import TransactionTable from "./transaction-table";
import TransactionsChart from "../molecules/transactions-chart";

const TransactionDetails = () => {
  const { mutate: logout } = useLogoutMutation();
  const [tableInstance, setTableInstance] = useState<Table<Transaction> | null>(
    null
  );

  return (
    <div className="p-4 ">
      <div className="w-full flex justify-end">
        <Button onClick={() => logout()}>
          <LogOut />
        </Button>
      </div>

      <TransactionTable
        table={tableInstance}
        setTableInstance={setTableInstance}
      />
      <TransactionsChart/>
    </div>
  );
};

export default TransactionDetails;
