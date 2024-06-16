import DataTable from "@/ui/atoms/data-table";
import UseTransactionColumns from "../atoms/columns";
import { SetStateAction, useCallback, useState } from "react";
import Transaction from "@/types/transactions";
import { Table } from "@tanstack/react-table";
import { useGetTransactions } from "@/app/hooks/requests";
import SelectComponent from "../molecules/transactions";
import Search from "../atoms/search";

type TransactionTableProps = {
  table: Table<Transaction> | null;
  setTableInstance: React.Dispatch<SetStateAction<Table<Transaction> | null>>;
};

const TransactionTable = ({
  setTableInstance,
  table,
}: TransactionTableProps) => {
  const { data, isLoading } = useGetTransactions();

  const [filtering, setFiltering] = useState("");

  const columns = UseTransactionColumns();
  const createTable = useCallback(
    (table: Table<Transaction>) => {
      setTableInstance(table);
    },
    [setTableInstance]
  );

  const filterColumnById = (id: string, value: string | string[]) => {
    const type = table?.getAllColumns().filter((x) => x.id === id)[0];
    type?.setFilterValue(value);
  };

  return (
    <>
      <Search filter={filtering} setFilter={setFiltering} />
      <SelectComponent
        columnFilter={filterColumnById}
        id="type"
        // clearFilters={table?.resetColumnFilters}
      />

      <DataTable
        createTable={createTable}
        filtering={filtering}
        setFiltering={setFiltering}
        data={data}
        isLoading={isLoading}
        tableCaption="Transactions"
        columns={columns}
      />
    </>
  );
};

export default TransactionTable;
