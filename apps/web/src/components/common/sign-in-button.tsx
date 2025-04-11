"use client";

import { signIn } from "next-auth/react";

import { Cog } from "lucide-react";

import { Button } from "@intania-vote/shadcn";

const SignInButton: React.FC = () => {
  return (
    <Button
      className="font-semibold"
      size="lg"
      onClick={() => {
        signIn("intania");
      }}
    >
      <Cog />
      Sign in with Intania
    </Button>
  );
};

export default SignInButton;
