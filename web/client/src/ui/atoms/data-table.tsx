"use client";

import TransactionDrawer from "@/app/components/transactions/organisms/transaction-drawer";
import { transactionKeys } from "@/app/hooks/requests";
import { axAPI } from "@/infrastructure/api";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCaption,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  ColumnDef,
  ColumnFiltersState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  Table as TanstackTableType,
  useReactTable,
} from "@tanstack/react-table";
import React, { SetStateAction, useEffect, useState } from "react";
import Paginator from "@/app/components/transactions/atoms/paginator";

type PaginationType = {
  pageIndex: number;
  pageSize: number;
};

type TableType<TData> = {
  data: TData[] | undefined;
  columns: ColumnDef<TData, any>[];
  pagination?: PaginationType;
  setPagination?: React.Dispatch<SetStateAction<PaginationType>>;
  sorting?: SortingState;
  setSorting?: React.Dispatch<SetStateAction<SortingState>>;
  isManualPagination?: boolean;
  pageSize?: number;
  isLoading: boolean;
  tableCaption: string;
  createTable: (table: TanstackTableType<TData>) => void;
  filtering: string;
  setFiltering: React.Dispatch<SetStateAction<string>>;
};

const DataTable = <T,>({
  data,
  isLoading,
  tableCaption,
  createTable,
  filtering,
  setFiltering,
  columns,
}: TableType<T>) => {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 10,
  });

  const queryClient = useQueryClient();
  const deleteTransaction = useMutation({
    mutationFn: (id: string) => axAPI.delete(`/api/deleteTransaction/${id}`),
    onSuccess() {
      queryClient.invalidateQueries({ queryKey: transactionKeys.all });
    },
  });

  const tableInstance = useReactTable({
    columns,
    data: data || [],
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    state: {
      sorting,
      globalFilter: filtering,
      columnFilters,
      pagination,
    },
    onSortingChange: setSorting,
    onGlobalFilterChange: setFiltering,
    onColumnFiltersChange: setColumnFilters,
    onPaginationChange: setPagination,
  });

  const rowData = () => data && tableInstance.getRowModel().rows;

  useEffect(() => {
    createTable(tableInstance);
  }, [tableInstance, createTable]);

  const row = rowData();

  if (isLoading) {
    return <div>...loading</div>;
  }

  // const pageCount = tableInstance.getPageCount();

  // const total = row?.length;

  return (
    <div className="flex flex-col ">
      <header className="text-center text-2xl font-bold">Transactions</header>

      <div className="flex space-x-2 justify-end">
        <Paginator data={data} table={tableInstance} />
      </div>

      <Table className="w-full ml-4 border-collapse divide-y divide-gray-400 text-xs">
        <TableCaption>{tableCaption}</TableCaption>

        <TableHeader>
          {tableInstance?.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <th
                  className="text-left cursor-pointer"
                  key={header.id}
                  onClick={header.column.getToggleSortingHandler()}
                >
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                  {{
                    asc: "🔼",
                    desc: "🔽",
                  }[header.column.getIsSorted() as string] ?? null}
                </th>
              ))}
            </TableRow>
          ))}
        </TableHeader>

        <TableBody>
          {row?.map((r) => (
            <React.Fragment key={r.id}>
              <TableRow key={r.id}>
                {r.getVisibleCells().map((cell) => {
                  return (
                    <td key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </td>
                  );
                })}
                <td>
                  <Button
                    onClick={() => deleteTransaction.mutate(r.getValue("ID"))}
                  >
                    Delete
                  </Button>
                </td>
              </TableRow>
            </React.Fragment>
          ))}
        </TableBody>
      </Table>
      <TransactionDrawer />
    </div>
  );
};

export default DataTable;
