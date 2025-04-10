"use client";

import { useEffect, useRef, useState } from "react";

import Image from "next/image";
import Link from "next/link";

import { User, Vote } from "lucide-react";
import moment from "moment";

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

  const isVoteActive = () => {
    const now = new Date();
    return startAt <= now && endAt >= now;
  };

  const calculateTime = (startAt: Date, endAt: Date): string | null => {
    const now = new Date();

    if (now > startAt && now < endAt) {
      const timeDiff = endAt.getTime() - now.getTime();
      const hours = Math.floor(
        (timeDiff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60),
      );
      const minutes = Math.floor((timeDiff % (1000 * 60 * 60)) / (1000 * 60));
      const seconds = Math.floor((timeDiff % (1000 * 60)) / 1000);

      if (timeDiff > 0) {
        if (timeDiff < 24 * 60 * 60 * 1000) {
          return `${String(hours).padStart(2, "0")}:${String(minutes).padStart(
            2,
            "0",
          )}:${String(seconds).padStart(2, "0")}`;
        } else {
          const days = Math.floor(timeDiff / (1000 * 60 * 60 * 24));
          return `เหลือเวลาอีก ${days} วัน`;
        }
      }
    } else if (now < startAt) {
      const timeDiff = startAt.getTime() - now.getTime();
      const hours = Math.floor(
        (timeDiff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60),
      );
      const minutes = Math.floor((timeDiff % (1000 * 60 * 60)) / (1000 * 60));
      const seconds = Math.floor((timeDiff % (1000 * 60)) / 1000);

      if (timeDiff > 0) {
        if (timeDiff < 24 * 60 * 60 * 1000) {
          return `เริ่มในอีก ${String(hours).padStart(2, "0")}:${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}`;
        } else {
          const days = Math.floor(timeDiff / (1000 * 60 * 60 * 24));
          if (days === 1) {
            return "เริ่มพรุ่งนี้";
          } else if (days > 1) {
            return `เริ่มในอีก ${days} วัน`;
          }
        }
      }
    } else if (now > endAt) {
      return moment(endAt).format("DD/MM/YYYY");
    }

    return null;
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
        return "/images/esc.jpg";
      case "ISESC":
        return "/images/isesc.jpg";
      default:
        return null;
    }
  };

  return (
    <>
      <Link
        href={`/vote/${slug}`}
        className={cn("w-full overflow-hidden", isDislaying ? "h-auto" : "h-0")}
      >
        <div ref={topContainerRef} className="relative z-10">
          <div
            className="flex flex-col gap-4 p-5 text-white"
            style={{ height: initialTopContainerHeight ?? "auto" }}
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
                  <span className="aspect-square w-2 rounded-full bg-red-500" />
                  <span>{time}</span>
                </div>
              ) : (
                <div className="text-xs font-bold text-white">
                  <span>{time}</span>
                </div>
              )}
            </div>
            <h2 className="text-3xl font-bold">{name}</h2>
          </div>
          <div className="bg-primary rounded-t-4xl rounded-br-4xl absolute inset-0 -z-10" />
        </div>
        <div className="flex">
          <div className="bg-primary rounded-b-4xl flex flex-grow -translate-y-px items-end gap-2 px-5 pb-5 pt-10 text-white">
            {choices.map((choice, index) => (
              <div
                key={index}
                className="flex h-12 w-full items-center justify-center rounded-full bg-black/15 px-2 text-center hover:bg-black/10"
              >
                <span className="text-sm">{choice.name}</span>
              </div>
            ))}
          </div>
          <div className="relative flex rounded-3xl bg-white pl-2 pt-2 text-neutral-700">
            <div className="z-20 flex flex-col items-center justify-center gap-1.5 rounded-2xl border-2 border-neutral-400 bg-white p-4 text-center">
              <Vote />
              <span className="text-xs">129</span>
            </div>
            <div className="bg-primary rounded-br-4xl rounded-tr-4xl rounded-bl-4xl absolute left-0 top-0 -z-10 aspect-square w-[4rem] -translate-x-px -translate-y-px" />
          </div>
        </div>
      </Link>
      {isDislaying ? null : <Skeleton className="rounded-4xl h-44 w-full" />}
    </>
  );
};

export default VoteCard;
