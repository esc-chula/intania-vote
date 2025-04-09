"use client";

import { signIn } from "next-auth/react";
import Link from "next/link";

import { Button } from "@intania-vote/shadcn";

const Page: React.FC = () => {
  return (
    <>
      <Button
        onClick={() => {
          signIn("intania");
        }}
      >
        Login with Intania
      </Button>
      <Link href="/vote/test">Test Election</Link>
    </>
  );
};

export default Page;
