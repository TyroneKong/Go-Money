import { Button } from "@/components/ui/button";
import FormInput from "@/ui/atoms/form-input";
import { FormProvider, SubmitHandler, useForm } from "react-hook-form";
import { formSchema, Inputs } from "../login-props";
import { zodResolver } from "@hookform/resolvers/zod";

type LoginFormProps = {
  onSubmit: SubmitHandler<Inputs>;
};

const LoginForm = ({ onSubmit }: LoginFormProps) => {
  const form = useForm({
    reValidateMode: "onChange",
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormInput label="username" name="username" type="text" />
        <FormInput label="password" name="password" type="text" />
        <Button disabled={!form?.formState.isValid} type="submit">
          Submit
        </Button>
      </form>
    </FormProvider>
  );
};

export default LoginForm;
