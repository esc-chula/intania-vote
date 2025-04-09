import Link from "next/link";

import { House, SquarePen, UserRound } from "lucide-react";

import { cn } from "@intania-vote/shadcn";

import TabBarWrapper from "./tab-bar-wrapper";

const Navigation: React.FC = () => {
  return (
    <TabBarWrapper>
      <nav className="flex h-full w-full items-center justify-around">
        <Link href="/" className={cn("bg-primary rounded-full p-3 text-white")}>
          <House />
        </Link>
        <Link
          href="/create"
          className={cn("rounded-full p-3 text-neutral-400")}
        >
          <SquarePen />
        </Link>
        <Link
          href="/profile"
          className={cn("rounded-full p-3 text-neutral-400")}
        >
          <UserRound />
        </Link>
      </nav>
    </TabBarWrapper>
  );
};

export default Navigation;
