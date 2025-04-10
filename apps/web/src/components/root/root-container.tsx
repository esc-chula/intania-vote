"use client";

import { signIn } from "next-auth/react";

import { createVote } from "~/server/vote";

import { Button } from "@intania-vote/shadcn";

const RootContainer: React.FC = () => {
  return (
    <div className="flex flex-col gap-4 p-4">
      <div>
        <Button
          onClick={() => {
            signIn("intania");
          }}
        >
          Login with Intania
        </Button>
      </div>
      <div>
        <Button
          onClick={() => {
            createVote({
              name: "Test Vote",
              description: "Test Description",
              slug: "test",
              image: "/mock.jpg",
              owner: "USER",
              eligibleStudentId: "*",
              eligibleDepartment: "*",
              eligibleYear: "*",
              isPrivate: false,
              isRealTime: false,
              isAllowMultiple: false,
              startAt: new Date(),
              endAt: new Date(new Date().getTime() + 1000 * 60 * 60),
              choices: [
                {
                  number: "1",
                  name: "Choice 1",
                  description: "Choice 1 Description",
                  information: "Choice 1 Information",
                },
              ],
            }).catch((err) => {
              console.error(err);
            });
          }}
        >
          Create Vote
        </Button>
      </div>
    </div>
  );
};

export default RootContainer;
