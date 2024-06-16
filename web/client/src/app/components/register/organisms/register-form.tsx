import { Button } from "@/components/ui/button";
import FormInput from "@/ui/atoms/form-input";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormProvider, SubmitHandler, useForm } from "react-hook-form";
import { formSchema, Inputs } from "../register-props";

type RegisterFormProps = {
  onSubmit: SubmitHandler<Inputs>;
};

const RegisterForm = ({ onSubmit }: RegisterFormProps) => {
  const form = useForm<Inputs>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      name: "",
      password: "",
      confirmPassword: "",
    },
  });

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormInput label="user name" name="username" type="text" />
        <FormInput label="name" name="name" type="text" />

        <FormInput label="password" name="password" type="text" />
        <FormInput
          label="confirm password"
          name="confirmPassword"
          type="text"
        />
        <Button type="submit">Submit</Button>
      </form>
    </FormProvider>
  );
};

export default RegisterForm;
