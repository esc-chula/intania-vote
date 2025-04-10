"use server";

import { z } from "zod";
import { getSession } from "~/lib/auth";
import { actionClient } from "~/lib/safe-action";

import { grpc } from "./grpc";

export const createBallotProof = actionClient
  .schema(
    z.object({
      choiceNumber: z.string(),
      voteSlug: z.string(),
    }),
  )
  .action(async ({ parsedInput: { choiceNumber, voteSlug } }) => {
    try {
      const session = await getSession();
      const oidcId = session?.user?.oidcId;
      if (!oidcId) {
        return { failure: "User is not authenticated" };
      }

      const proof = await grpc.ballot.CreateBallotProof({
        oidcId,
        voteSlug,
        choiceNumber,
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

export const verifyBallot = actionClient
  .schema(
    z.object({
      ballotKey: z.string(),
    }),
  )
  .action(async ({ parsedInput: { ballotKey } }) => {
    try {
      const session = await getSession();
      const oidcId = session?.user?.oidcId;
      if (!oidcId) {
        return { failure: "User is not authenticated" };
      }

      const verifiedBallot = await grpc.ballot.VerifyBallot({
        ballotKey,
        oidcId,
      });

      return { success: "Successfully verified ballot", verifiedBallot };
    } catch (error) {
      console.error("Error in verify ballot action:", error);
      return {
        failure: "An error occurred during verify ballot",
      };
    }
  });
