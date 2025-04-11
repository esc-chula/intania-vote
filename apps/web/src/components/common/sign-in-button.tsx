"use client";

import { signIn } from "next-auth/react";

import { Button } from "@intania-vote/shadcn";

const SignInButton: React.FC = () => {
  return (
    <Button
      onClick={() => {
        signIn("intania");
      }}
    >
      Login with Intania
    </Button>
  );
};

export default SignInButton;
