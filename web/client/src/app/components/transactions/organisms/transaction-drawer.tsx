import { zodResolver } from "@hookform/resolvers/zod";
import { useQueryClient } from "@tanstack/react-query";
import { transactionKeys, useCreateNewTransaction } from "@/app/hooks/requests";
// import { toast } from '@/components/ui/use-toast';
import Transaction from "@/types/transactions";
import TransactionForm from "../molecules/transaction-form";
import { Inputs, transactionDrawerSchema } from "../transaction-props";
import { SubmitHandler, useForm } from "react-hook-form";
import GeneralDrawerForm from "@/ui/atoms/general-drawer-form";

function TransactionDrawer() {
  const form = useForm<Inputs>({
    resolver: zodResolver(transactionDrawerSchema),
    defaultValues: {
      userId: "",
      name: "",
      amount: "",
      currency: "",
      type: "",
      bank: "",
    },
  });

  const queryClient = useQueryClient();

  const createOnSuccess = async (newTransaction: Inputs) => {
    const previousData = queryClient.getQueryData(transactionKeys.all);

    await queryClient.setQueryData(
      transactionKeys.all,
      (oldData: Transaction[]) => [...oldData, newTransaction]
    );
    return { previousData };
  };

  const createTransaction = useCreateNewTransaction({
    onSuccess: ({ data }) => {
      createOnSuccess(data);
    },
    onSettled: () =>
      queryClient.invalidateQueries({ queryKey: transactionKeys.all }),
  });

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    createTransaction.mutate(data);
  };

  return (
    <GeneralDrawerForm
      form={form}
      onSubmit={onSubmit}
      title="Create a new Transaction"
    >
      <TransactionForm />
    </GeneralDrawerForm>
  );
}

export default TransactionDrawer;
