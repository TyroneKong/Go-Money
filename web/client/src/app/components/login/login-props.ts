import { z } from "zod";

export const formSchema = z.object({
  username: z.string().min(4),
  password: z.string().min(8),
});

export type Inputs = z.infer<typeof formSchema>;
