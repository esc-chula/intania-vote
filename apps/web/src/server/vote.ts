"use server";

import { z } from "zod";
import { getSession } from "~/lib/auth";
import { actionClient } from "~/lib/safe-action";

import { Owner } from "@intania-vote/grpc-ts";

import { grpc } from "./grpc";

export const createVote = actionClient.action(async () => {
  try {
    const session = await getSession();
    const oidcId = session?.user?.oidcId;
    if (!oidcId) {
      return { failure: "User is not authenticated" };
    }

    const newVote = await grpc.vote.CreateVote({
      oidcId,
      vote: {
        name: "Test Name",
        description: "Test Description",
        slug: "test-name",
        image: "/mock.jpg",
        owner: Owner.STUDENT,
        eligibleStudentId: "*",
        eligibleDepartment: "*",
        eligibleYear: "*",
        isPrivate: false,
        isRealTime: false,
        isAllowMultiple: false,
        startAt: new Date().toISOString(),
        endAt: new Date(new Date().getTime() + 1000 * 60 * 60).toISOString(),
      },
      choices: [
        {
          number: "1",
          name: "Choice 1",
          description: "Choice 1 Description",
          information: "Choice 1 Information",
        },
      ],
    });
    if (newVote) {
      return { success: "Successfully created vote" };
    }
  } catch (error) {
    console.error("Error in create vote action:", error);
    return { failure: "An error occurred during create vote" };
  }
});
