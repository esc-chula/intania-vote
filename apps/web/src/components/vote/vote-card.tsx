"use client";

import { useEffect, useRef, useState } from "react";

import Link from "next/link";

import { Vote } from "lucide-react";

import { cn, Skeleton } from "@intania-vote/shadcn";

interface VoteCardProps {
  name: string;
  slug: string;
  image?: string;
  owner: string;
  startAt: Date;
  endAt: Date;
  choices: {
    name: string;
  }[];
}

const VoteCard: React.FC<VoteCardProps> = ({
  name,
  slug,
  owner,
  startAt,
  endAt,
  choices,
}) => {
  const topContainerRef = useRef<HTMLDivElement>(null);
  const [initialTopContainerHeight, setInitialTopContainerHeight] = useState<
    number | null
  >(null);

  useEffect(() => {
    if (topContainerRef.current) {
      setInitialTopContainerHeight(topContainerRef.current.clientHeight - 32);
    }
  }, []);

  return (
    <>
      <Link
        href={`/vote/${slug}`}
        className={cn(
          "flex w-fit overflow-hidden",
          initialTopContainerHeight ? "h-auto" : "h-0",
        )}
      >
        <div className="max-w-md">
          <div ref={topContainerRef} className="relative z-10">
            <div
              className="flex flex-col gap-2 p-5 text-white"
              style={{ height: initialTopContainerHeight ?? "auto" }}
            >
              <div className="flex items-center justify-between">
                <div>
                  <span className="text-xs">{owner}</span>
                </div>
              </div>
              <h2 className="text-3xl font-bold">{name}</h2>
            </div>
            <div className="bg-primary rounded-t-4xl rounded-br-4xl absolute inset-0 -z-10" />
          </div>
          <div className="flex">
            <div className="bg-primary rounded-b-4xl flex flex-grow -translate-y-px items-end gap-2 p-5 text-white">
              {choices.map((choice, index) => (
                <div
                  key={index}
                  className="flex h-10 w-full items-center justify-center rounded-full bg-black/15 px-2 text-center hover:bg-black/10"
                >
                  <span className="text-sm">{choice.name}</span>
                </div>
              ))}
            </div>
            <div className="relative flex rounded-3xl bg-white pl-2 pt-2 text-neutral-700">
              <div className="z-20 flex flex-col items-center justify-center gap-1.5 rounded-3xl border-2 border-neutral-400 bg-white p-4 text-center">
                <Vote />
                <span className="text-xs">129</span>
              </div>
              <div className="bg-primary rounded-br-4xl rounded-tr-4xl rounded-bl-4xl absolute left-0 top-0 -z-10 aspect-square w-[4rem] -translate-x-px -translate-y-px" />
            </div>
          </div>
        </div>
      </Link>
      {initialTopContainerHeight ? null : (
        <Skeleton className="rounded-4xl h-44 w-full" />
      )}
    </>
  );
};

export default VoteCard;
