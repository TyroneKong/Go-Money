import { z } from "zod";

export const formSchema = z
  .object({
    username: z.string().min(2, {
      message: "Username must be at least 2 characters.",
    }),
    name: z.string(),
    password: z.string(),
    confirmPassword: z.string(),
  })
  .refine((x) => x.password === x.confirmPassword, {
    path: ["confirmPassword"],
    message: "passwords do not match",
  });

export type Inputs = z.infer<typeof formSchema>;
