"use client";

import { useState } from "react";

import { Tally } from "@intania-vote/grpc-ts";
import { Button } from "@intania-vote/shadcn";

import VoteChoicesInfo from "./vote-choices-info";
import VoteResultCard from "./vote-result-card";
import VoteVerifyCard from "./vote-verify-card";

interface VoteResultContainerProps {
  name: string;
  description: string;
  slug: string;
  startAt: Date;
  endAt: Date;
  choices: {
    number: number;
    name: string;
    description: string;
    information?: string;
    image?: string;
  }[];
  isEligible?: boolean;
  tally?: Tally;
}

const VoteResultContainer: React.FC<VoteResultContainerProps> = ({
  name,
  description,
  slug,
  startAt,
  endAt,
  choices,
  isEligible = true,
  tally,
}) => {
  const [tab, setTab] = useState<"result" | "info">("result");

  return (
    <>
      <div className="flex flex-grow flex-col gap-4 p-5">
        <div className="flex flex-shrink flex-col pt-2">
          <h1 className="pr-20 text-3xl font-bold text-neutral-800">{name}</h1>
          <p className="text-neutral-600">{description}</p>
        </div>
        <div className="flex flex-grow flex-col gap-5">
          <div className="grid grid-cols-2 gap-5">
            <Button
              variant={tab === "result" ? "default" : "outline"}
              onClick={() => setTab("result")}
              size="sm"
            >
              ผลการโหวต
            </Button>
            <Button
              variant={tab === "info" ? "default" : "outline"}
              onClick={() => setTab("info")}
              size="sm"
            >
              ข้อมูลการโหวต
            </Button>
          </div>
          {tab === "result" ? (
            <>
              <VoteResultCard
                startAt={startAt}
                endAt={endAt}
                choices={choices}
                tally={tally}
              />
              {isEligible ? <VoteVerifyCard /> : null}
            </>
          ) : (
            <>
              <VoteChoicesInfo
                choices={choices.sort((a, b) => a.number - b.number)}
              />
            </>
          )}
        </div>
      </div>
    </>
  );
};

export default VoteResultContainer;
