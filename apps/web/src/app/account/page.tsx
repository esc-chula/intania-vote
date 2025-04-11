"use client";

import { signOut, useSession } from "next-auth/react";

import { LogOut, User } from "lucide-react";
import Navigation from "~/components/common/navigation";

import { Button } from "@intania-vote/shadcn";

const Page: React.FC = () => {
  const session = useSession();

  return (
    <>
      <Button
        variant="outline"
        size="icon"
        className="fixed right-6 top-5 z-50 h-14 w-14 rounded-full"
        onClick={() => {
          signOut();
        }}
      >
        <LogOut />
      </Button>
      <div className="flex flex-grow flex-col gap-4 p-5">
        <div className="flex flex-shrink flex-col pt-2">
          <h1 className="pr-20 text-3xl font-bold text-neutral-800">Account</h1>
        </div>
        <div className="flex items-center gap-5 pt-4">
          <div className="grid aspect-square w-36 place-content-center rounded-full border border-neutral-200 bg-neutral-50 text-neutral-300">
            <User size={48} />
          </div>
          <p className="text-2xl font-bold text-neutral-900">
            {session.data?.user.studentId}
          </p>
        </div>
      </div>
      <Navigation />
    </>
  );
};

export default Page;
