import Link from "next/link";

import { House, SquarePen, UserRound } from "lucide-react";

import { cn } from "@intania-vote/shadcn";

const Navigation = () => {
  return (
    <nav className="fixed bottom-0 left-0 right-0 flex items-center justify-around border-t border-neutral-200 bg-neutral-50 p-3">
      <Link href="/" className={cn("bg-primary rounded-full p-3 text-white")}>
        <House />
      </Link>
      <Link href="/create" className={cn("rounded-full p-3 text-neutral-400")}>
        <SquarePen />
      </Link>
      <Link href="/profile" className={cn("rounded-full p-3 text-neutral-400")}>
        <UserRound />
      </Link>
    </nav>
  );
};

export default Navigation;
