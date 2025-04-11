"use server";

import { z } from "zod";
import { getSession } from "~/lib/auth";
import { actionClient } from "~/lib/safe-action";

import { Owner } from "@intania-vote/grpc-ts";

import { grpc } from "./grpc";

export const createVote = actionClient
  .schema(
    z.object({
      name: z.string().max(100),
      description: z.string().max(300),
      slug: z.string().max(100),
      image: z.string().max(100).optional(),
      owner: z.enum(["USER", "ESC", "ISESC"]),
      eligibleStudentId: z.string().max(50),
      eligibleDepartment: z.string().max(50),
      eligibleYear: z.string().max(50),
      isPrivate: z.boolean(),
      isRealTime: z.boolean(),
      isAllowMultiple: z.boolean(),
      startAt: z.date(),
      endAt: z.date(),
      choices: z.array(
        z.object({
          number: z.number(),
          name: z.string().max(50),
          description: z.string().max(50),
          information: z.string().max(700).optional(),
          image: z.string().max(100).optional(),
        }),
      ),
    }),
  )
  .action(
    async ({
      parsedInput: {
        name,
        description,
        slug,
        image,
        owner,
        eligibleStudentId,
        eligibleDepartment,
        eligibleYear,
        isPrivate,
        isRealTime,
        isAllowMultiple,
        startAt,
        endAt,
        choices,
      },
    }) => {
      try {
        const session = await getSession();
        const oidcId = session?.user?.oidcId;
        if (!oidcId) {
          return { failure: "User is not authenticated" };
        }

        let ownerEnum: Owner;
        switch (owner) {
          case "USER":
            ownerEnum = Owner.USER;
            break;
          case "ESC":
            ownerEnum = Owner.ESC;
            break;
          case "ISESC":
            ownerEnum = Owner.ISESC;
            break;
          default:
            return { failure: "Invalid owner type" };
        }

        const newVote = await grpc.vote.CreateVote({
          oidcId,
          vote: {
            name,
            description,
            slug,
            image,
            owner: ownerEnum,
            eligibleStudentId,
            eligibleDepartment,
            eligibleYear,
            isPrivate,
            isRealTime,
            isAllowMultiple,
            startAt: startAt.toISOString(),
            endAt: endAt.toISOString(),
          },
          choices,
        });
        if (newVote) {
          return { success: "Successfully created vote" };
        }
      } catch (error) {
        console.error("Error in create vote action:", error);
        return { failure: "An error occurred during create vote" };
      }
    },
  );

export const getVoteBySlug = actionClient
  .schema(z.object({ slug: z.string().max(100) }))
  .action(async ({ parsedInput: { slug } }) => {
    try {
      const vote = await grpc.vote.GetVoteBySlug({ slug });

      return {
        success: "Successfully fetched vote",
        vote,
      };
    } catch (error) {
      console.error("Error in get vote by slug action:", error);
      return {
        failure: "An error occurred during get vote by slug",
      };
    }
  });

export const getVotes = actionClient.action(async () => {
  try {
    const votes = await grpc.vote.GetVotes({});

    return {
      success: "Successfully fetched votes",
      votes,
    };
  } catch (error) {
    console.error("Error in get votes action:", error);
    return {
      failure: "An error occurred during get votes",
    };
  }
});

export const getVotesByUserEligibility = actionClient.action(async () => {
  try {
    const session = await getSession();
    const oidcId = session?.user?.oidcId;
    if (!oidcId) {
      return { failure: "User is not authenticated" };
    }

    const votes = await grpc.vote.GetVotesByUserEligibility({ oidcId });

    return {
      success: "Successfully fetched votes",
      votes,
    };
  } catch (error) {
    console.error("Error in get votes by user eligibility action:", error);
    return {
      failure: "An error occurred during get votes by user eligibility",
    };
  }
});

export const hasUserVoted = actionClient
  .schema(z.object({ slug: z.string().max(100) }))
  .action(async ({ parsedInput: { slug } }) => {
    try {
      const session = await getSession();
      const oidcId = session?.user?.oidcId;
      if (!oidcId) {
        return { failure: "User is not authenticated" };
      }

      const hasVoted = await grpc.vote.HasUserVoted({ slug, oidcId });

      return {
        success: "Successfully checked if user has voted",
        hasVoted,
      };
    } catch (error) {
      console.error("Error in has user voted action:", error);
      return {
        failure: "An error occurred during check if user has voted",
      };
    }
  });
