"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

import { House, SquarePen, UserRound } from "lucide-react";

import { cn } from "@intania-vote/shadcn";

import TabBarWrapper from "./tab-bar-wrapper";

const Navigation: React.FC = () => {
  const pathname = usePathname();

  return (
    <TabBarWrapper>
      <nav className="flex h-full w-full items-center justify-around">
        <Link
          href="/home"
          className={cn(
            "rounded-full p-3.5",
            pathname === "/home"
              ? "bg-primary text-white"
              : "texxt-neutral-400",
          )}
        >
          <House />
        </Link>
        <Link
          href="/new"
          className={cn(
            "rounded-full p-3.5",
            pathname === "/new" ? "bg-primary text-white" : "texxt-neutral-400",
          )}
        >
          <SquarePen />
        </Link>
        <Link
          href="/account"
          className={cn(
            "rounded-full p-3.5",
            pathname === "/account"
              ? "bg-primary text-white"
              : "texxt-neutral-400",
          )}
        >
          <UserRound />
        </Link>
      </nav>
    </TabBarWrapper>
  );
};

export default Navigation;
