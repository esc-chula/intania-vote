"use client";

import { useEffect, useRef, useState } from "react";

import Image from "next/image";
import { useRouter } from "next/navigation";

import { User, Vote } from "lucide-react";
import { calculateTime } from "~/lib/vote";

import { cn, Skeleton } from "@intania-vote/shadcn";

interface VoteCardProps {
  name: string;
  slug: string;
  image?: string;
  owner: string;
  isEligible: boolean;
  startAt: Date;
  endAt: Date;
  choices: {
    name: string;
  }[];
  totalBallots: number;
}

const VoteCard: React.FC<VoteCardProps> = ({
  name,
  slug,
  owner,
  isEligible,
  startAt,
  endAt,
  choices,
  totalBallots,
}) => {
  const router = useRouter();

  const topContainerRef = useRef<HTMLDivElement>(null);
  const [initialTopContainerHeight, setInitialTopContainerHeight] = useState<
    number | null
  >(null);

  useEffect(() => {
    if (topContainerRef.current) {
      setInitialTopContainerHeight(topContainerRef.current.clientHeight - 32);
    }
  }, []);

  const isVoteActive = () => {
    const now = new Date();
    return startAt <= now && endAt >= now;
  };

  const [time, setTime] = useState<string | null>(null);

  useEffect(() => {
    setTime(calculateTime(startAt, endAt));

    const interval = setInterval(() => {
      const calculatedTime = calculateTime(startAt, endAt);
      setTime(calculatedTime);
    }, 1000);

    return () => clearInterval(interval);
  }, [endAt, startAt]);

  const isDislaying = initialTopContainerHeight && time ? true : false;

  const ownerImageSrc = (): string | null => {
    switch (owner) {
      case "User":
        return null;
      case "ESC":
        return "/assets/esc.jpg";
      case "ISESC":
        return "/assets/isesc.jpg";
      default:
        return null;
    }
  };

  const now = new Date();

  return (
    <div className="flex flex-col">
      <div
        className={cn(
          "w-full cursor-pointer overflow-hidden",
          isDislaying ? "h-auto" : "h-0",
        )}
        onClick={() => {
          router.push(`/vote/${slug}`);
        }}
      >
        <div ref={topContainerRef} className="relative z-10">
          <div
            className={cn(
              "flex flex-col gap-4 p-5 text-white",
              now > startAt && isEligible ? "text-white" : "text-neutral-800",
            )}
            style={{
              height:
                now > startAt && isEligible
                  ? (initialTopContainerHeight ?? "auto")
                  : "auto",
            }}
          >
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-1.5">
                {ownerImageSrc() ? (
                  <div className="relative aspect-square w-4 overflow-hidden rounded-full">
                    <Image
                      src={ownerImageSrc() || ""}
                      alt={owner}
                      fill
                      className="object-contain"
                      sizes="100%"
                    />
                  </div>
                ) : (
                  <div className="grid aspect-square w-5 place-content-center rounded-full bg-white text-neutral-800">
                    <User size={12} />
                  </div>
                )}

                <span className="text-xs">
                  {owner === "ESC" ? "กวศ." : owner}
                </span>
              </div>
              {isVoteActive() ? (
                <div className="flex items-center gap-1.5 rounded-full bg-white py-1 pl-2.5 pr-3 text-xs font-bold text-neutral-700">
                  <span className="aspect-square w-2 animate-pulse rounded-full bg-red-500" />
                  <span>{time}</span>
                </div>
              ) : (
                <div
                  className={cn(
                    "text-xs font-bold",
                    now > startAt && isEligible
                      ? "text-white"
                      : "text-neutral-600",
                  )}
                >
                  <span>{time}</span>
                </div>
              )}
            </div>
            <h2 className="text-3xl font-bold">{name}</h2>
            {!isEligible && now > endAt ? (
              <div className="flex h-12 w-full items-center justify-center rounded-full border border-neutral-200 bg-white font-semibold text-neutral-600 hover:bg-white/90">
                ดูผล
              </div>
            ) : null}
          </div>
          <div
            className={cn(
              "absolute inset-0 -z-10",
              now > startAt && isEligible
                ? "bg-primary rounded-t-4xl rounded-br-4xl"
                : "rounded-4xl border border-neutral-200 bg-neutral-50",
            )}
          />
        </div>
        {now > startAt && isEligible ? (
          <div className="flex">
            <div className="bg-primary rounded-b-4xl flex flex-grow -translate-y-px items-end gap-2 px-5 pb-5 pt-10 text-white">
              {now > endAt ? (
                <div className="text-primary flex h-12 w-full items-center justify-center rounded-full bg-white font-semibold hover:bg-white/90">
                  ดูผล
                </div>
              ) : (
                <div className="text-primary flex h-12 w-full items-center justify-center rounded-full bg-white font-semibold hover:bg-white/90">
                  โหวตเลย
                </div>

                // choices.map((choice, index) => (
                //   <div
                //     key={index}
                //     className="flex h-12 w-full items-center justify-center rounded-full bg-black/15 px-2 text-center hover:bg-black/10"
                //   >
                //     <span className="text-sm">{choice.name}</span>
                //   </div>
                // ))
              )}
              {}
            </div>
            <div className="relative flex rounded-3xl bg-white pl-2 pt-2 text-neutral-700">
              <div className="z-20 flex flex-col items-center justify-center gap-1.5 rounded-2xl border-2 border-neutral-400 bg-white p-4 text-center">
                <Vote />
                <span className="text-xs">{totalBallots}</span>
              </div>
              <div className="bg-primary rounded-br-4xl rounded-tr-4xl rounded-bl-4xl absolute left-0 top-0 -z-10 aspect-square w-[4rem] -translate-x-px -translate-y-px" />
            </div>
          </div>
        ) : null}
      </div>
      {isDislaying ? null : <Skeleton className="rounded-4xl h-44 w-full" />}
    </div>
  );
};

export default VoteCard;
