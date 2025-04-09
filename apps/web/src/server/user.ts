"use server";

import { z } from "zod";
import { actionClient } from "~/lib/safe-action";

import { grpc } from "./grpc";

export const login = actionClient
  .schema(
    z.object({
      oidcId: z.string(),
      studentId: z.string(),
    }),
  )
  .action(async ({ parsedInput: { oidcId, studentId } }) => {
    try {
      const newUser = await grpc.user.CreateUser({
        oidcId,
        studentId,
      });
      if (newUser) {
        return { success: "Successfully logged in" };
      }
    } catch (error) {
      console.error("Error in login action:", error);
      return { failure: "An error occurred during login" };
    }
  });
