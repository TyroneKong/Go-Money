"use client";
import {
  axAPI,
  getAllExpenses,
  getAllTransactions,
  getCurrentUser,
} from "@/infrastructure/api";
import {
  useQuery,
  queryOptions,
  useMutation,
  UseMutationOptions,
} from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import Params from "../constants/params";
import { useRouter } from "next/navigation";
import { Inputs } from "../components/transactions/transaction-props";

type RegisterParamsMappedType<T> = { [k in keyof T]: string };

export type RegisterParams = RegisterParamsMappedType<typeof Params>;

type MyPick<T, K extends keyof T> = { [p in K]: T[p] };

type LoginParams = MyPick<
  RegisterParams,
  "username" | "password" | "confirmPassword"
>;

export const transactionKeys = {
  all: ["transactions"],
};

const Transactions = {
  all: transactionKeys.all,
  list: () =>
    queryOptions({
      queryKey: Transactions.all,
      queryFn: () => getAllTransactions(),
    }),
};

export const useGetTransactions = () => {
  return useQuery(Transactions.list());
};

const Expenses = {
  all: ["expenses"],
  list: () =>
    queryOptions({
      queryKey: [...Expenses.all],
      queryFn: () => getAllExpenses(),
    }),
};

export const useGetExpenses = () => {
  return useQuery(Expenses.list());
};

const currentUserKeys = {
  key: ["me"],
};

export const useQueryCurrentUser = () => {
  return useQuery({
    queryKey: currentUserKeys.key,
    queryFn: () => getCurrentUser(),
  });
};

const query = async (body: LoginParams) => {
  const response = await axAPI.post("/api/login", body, {
    withCredentials: true,
  });

  return response;
};

const loginMutationKey = {
  key: ["me"],
};

export const useLoginMutation = (
  mutationProps: Omit<
    UseMutationOptions<AxiosResponse, Error, LoginParams, unknown>,
    "mutationFn" | "mutationKey"
  >
) => {
  const { onSuccess, onError } = mutationProps;
  const { mutate, ...rest } = useMutation({
    mutationKey: loginMutationKey.key,
    mutationFn: (body: LoginParams) => query(body),
    onSuccess,
    onError,
  });

  return {
    mutate,
    ...mutationProps,
    ...rest,
  };
};

const register = {
  key: ["register"],
};

const registerQuery = (body: RegisterParams) =>
  axAPI.post("/api/register", body, {
    withCredentials: true,
  });

export const useRegisterMutation = (
  mutationProps: Omit<
    UseMutationOptions<AxiosResponse, Error, RegisterParams, unknown>,
    "mutationFn" | "mutationKey"
  >
) => {
  const { onSuccess, onError } = mutationProps;
  const { mutate, ...rest } = useMutation({
    mutationKey: register.key,
    mutationFn: (body: RegisterParams) => registerQuery(body),
    onSuccess,
    onError,
  });

  return {
    mutate,
    ...mutationProps,
    ...rest,
  };
};

export const useLogoutMutation = () => {
  const router = useRouter();

  const { mutate, ...rest } = useMutation({
    mutationKey: ["logout"],
    mutationFn: () => axAPI.post("/api/logout"),
    onSuccess: () => {
      router.push("/");
    },
  });

  return {
    mutate,
    ...rest,
  };
};

export const useCreateNewTransaction = (
  mutationProps: Omit<
    UseMutationOptions<AxiosResponse, Error, unknown>,
    "mutationFn" | "mutationKey"
  >
) => {
  const { onSuccess, onSettled } = mutationProps;

  const { mutate, ...rest } = useMutation({
    mutationFn: (body: Inputs) => axAPI.post("/api/createTransaction", body),
    onSuccess,
    onSettled,
  });

  return {
    mutate,
    ...rest,
  };
};
