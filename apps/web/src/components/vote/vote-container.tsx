"use client";

import { useState } from "react";

import TabBarWrapper from "~/components/common/tab-bar-wrapper";
import VoteCard from "~/components/vote/vote-card";

import { Button } from "@intania-vote/shadcn";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@intania-vote/shadcn";

interface VoteContainerProps {
  name: string;
  description: string;
  choices: {
    number?: string;
    name: string;
    description: string;
    image: string;
  }[];
}

const VoteContainer: React.FC<VoteContainerProps> = ({
  name,
  description,
  choices,
}) => {
  const [selectedChoice, setSelectedChoice] = useState<number>(0);

  const selectedChoiceData = choices[selectedChoice - 1];
  const selectedChoiceDisplay = selectedChoiceData?.number
    ? `เบอร์ ${selectedChoiceData?.number}`
    : selectedChoiceData?.name;

  return (
    <>
      <div className="flex flex-grow flex-col gap-4 px-5 pb-20">
        <div className="flex flex-shrink flex-col pt-7">
          <h1 className="pr-20 text-3xl font-bold text-neutral-800">{name}</h1>
          <p className="text-neutral-600">{description}</p>
        </div>
        <div className="grid flex-grow grid-cols-1 grid-rows-2 gap-5 pb-5">
          {choices.map((choice, index) => (
            <VoteCard
              key={index}
              isActive={selectedChoice === index + 1}
              onClick={() => {
                if (selectedChoice === index + 1) {
                  setSelectedChoice(0);
                } else {
                  setSelectedChoice(index + 1);
                }
              }}
              onVote={() => {
                setSelectedChoice(index + 1);
              }}
              number={choice.number}
              name={choice.name}
              description={choice.description}
              info={`
## Test

- bullet test
- another bullet test

<br />

### Heading 3

> what do you mean?

<br />

lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec
                `}
              image={choice.image}
            />
          ))}
        </div>
      </div>
      <TabBarWrapper>
        <div className="flex h-full w-full flex-col items-center justify-center px-2 py-2">
          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button
                variant={selectedChoice ? "default" : "secondary"}
                size="lg"
                className="h-full w-full font-semibold"
              >
                {selectedChoice
                  ? `เลือก ${selectedChoiceDisplay}`
                  : "ไม่ออกเสียง"}
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>
                  ยืนยัน
                  {selectedChoice
                    ? `เลือก ${selectedChoiceDisplay}`
                    : "ไม่ออกเสียง"}
                </AlertDialogTitle>
                <AlertDialogDescription>
                  {selectedChoice
                    ? `คุณได้เลือก ${selectedChoiceData?.name} ${selectedChoiceData?.description}`
                    : "คุณต้องการไม่ออกเสียงใช่หรือไม่?"}
                  <br />
                  หากยืนยันการเลือกแล้ว จะไม่สามารถเปลี่ยนแปลงได้
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel>ยกเลิก</AlertDialogCancel>
                <AlertDialogAction>ยืนยัน</AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        </div>
      </TabBarWrapper>
    </>
  );
};

export default VoteContainer;
