"use client";

import { signIn } from "next-auth/react";
import Link from "next/link";

import { grpc } from "~/server/grpc";
import { createVote } from "~/server/vote";

import { Owner } from "@intania-vote/grpc-ts";
import { Button } from "@intania-vote/shadcn";

const Page: React.FC = () => {
  return (
    <>
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
        <Link href="/vote/test">
          <Button>Test Election</Button>
        </Link>
        <div>
          <Button
            onClick={() => {
              createVote().catch((err) => {
                console.error(err);
              });
            }}
          >
            Create Vote
          </Button>
        </div>
      </div>
    </>
  );
};

export default Page;
