"use server";

import { z } from "zod";
import { getSession } from "~/lib/auth";
import { actionClient } from "~/lib/safe-action";

import { grpc } from "./grpc";

export const createBallotProof = actionClient
  .schema(
    z.object({
      choiceId: z.number(),
    }),
  )
  .action(async ({ parsedInput: { choiceId } }) => {
    try {
      const session = await getSession();
      const oidcId = session?.user?.oidcId;
      if (!oidcId) {
        return { failure: "User is not authenticated" };
      }

      const proof = await grpc.ballot.CreateBallotProof({
        choiceId,
        oidcId,
      });

      return { success: "Successfully created ballot proof", proof };
    } catch (error) {
      console.error("Error in create ballot proof action:", error);
      return {
        failure: "An error occurred during create ballot proof",
      };
    }
  });

export const createBallot = actionClient
  .schema(
    z.object({
      voteSlug: z.string(),
      proof: z.object({
        commitment: z.string(),
        blindingFactor: z.string(),
        challenge: z.string(),
        response: z.string(),
        nullifier: z.string(),
        receipt: z.string(),
      }),
    }),
  )
  .action(async ({ parsedInput: { voteSlug, proof } }) => {
    try {
      const session = await getSession();
      const oidcId = session?.user?.oidcId;
      if (!oidcId) {
        return { failure: "User is not authenticated" };
      }

      const ballot = await grpc.ballot.CreateBallot({
        oidcId,
        voteSlug,
        proof,
      });

      return { success: "Successfully created ballot", ballot };
    } catch (error) {
      console.error("Error in create ballot action:", error);
      return {
        failure: "An error occurred during create ballot",
      };
    }
  });
