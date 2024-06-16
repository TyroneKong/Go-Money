"use client";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import Image from "next/image";
import { useRegisterMutation } from "@/app/hooks/requests";
import { Card } from "@/components/ui/card";
import RegisterForm from "./organisms/register-form";
import { SubmitHandler } from "react-hook-form";
import { Inputs } from "./register-props";

const Register = () => {
  const router = useRouter();

  // const { toast } = useToast();

  const register = useRegisterMutation({
    onSuccess: () => {
      router.push("/login");
    },
  });

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    register.mutate(data);
    register.onSuccess;
  };

  return (
    <div>
      <Button
        variant="default"
        className="mt-5 w-full"
        onClick={() => router.push("/login")}
      >
        Login
      </Button>
      <h1 className="text-4xl font-bold text-center mb-3">GoMoney</h1>
      <Card className="mb-5">
        <Image src="/piggybank.jpg" alt="piggybank" width={300} height={300} />
      </Card>

      <RegisterForm onSubmit={onSubmit} />
    </div>
  );
};

export default Register;
