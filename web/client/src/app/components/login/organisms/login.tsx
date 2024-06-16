"use client";

import { useRouter } from "next/navigation";
import { Card } from "@/components/ui/card";
import Image from "next/image";
import { useLoginMutation } from "@/app/hooks/requests";
import LoginForm from "../molecules/login-form";
import Link from "next/link";

function Login() {
  const router = useRouter();

  // const { toast } = useToast();

  const { mutate: login } = useLoginMutation({
    onSuccess: () => {
      router.push("/transactions");
    },
  });

  const onSubmit = (data: any) => {
    login(data);
  };
  // ...

  return (
    <div>
      <h1 className="text-4xl font-bold text-center mb-3">Login</h1>
      <Card className="mb-5">
        <Image src="/piggybank.jpg" alt="piggybank" width={400} height={400} />
      </Card>

      <LoginForm onSubmit={onSubmit} />
      <div className="flex space-x-2 mt-2">
        <p>Not yet registered?</p>
        <Link className="text-blue-500 underline " href="/register">
          Create an account
        </Link>
      </div>
    </div>
  );
}

export default Login;
