"use client";

import { useState } from "react";

import { useRouter } from "next/navigation";

import { LoaderCircle } from "lucide-react";
import TabBarWrapper from "~/components/common/tab-bar-wrapper";
import VoteBallotCard from "~/components/vote/vote-ballot-card";
import { choicesGridClassName } from "~/lib/vote";
import { createBallot, createBallotProof } from "~/server/ballot";

import { Button, cn, useToast } from "@intania-vote/shadcn";
import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@intania-vote/shadcn";

interface VoteBallotContainerProps {
  name: string;
  description: string;
  slug: string;
  choices: {
    number: number;
    name: string;
    description: string;
    information?: string;
    image?: string;
  }[];
}

const VoteBallotContainer: React.FC<VoteBallotContainerProps> = ({
  name,
  description,
  slug,
  choices,
}) => {
  const router = useRouter();
  const { toast } = useToast();

  const [selectedChoice, setSelectedChoice] = useState<number>(0);

  const selectedChoiceData = choices[selectedChoice - 1];
  const selectedChoiceDisplay = selectedChoiceData?.number
    ? `เบอร์ ${selectedChoiceData?.number}`
    : selectedChoiceData?.name;

  const [isOpenConfirmationAlert, setIsOpenConfirmationAlert] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleVote = async () => {
    setLoading(true);

    const resCreateBallotProof = await createBallotProof({
      choiceNumber: selectedChoice,
      voteSlug: slug,
    });

    if (
      resCreateBallotProof?.data?.failure ||
      !resCreateBallotProof?.data?.proof
    ) {
      toast({
        title: "เกิดข้อผิดพลาด",
        description: "ไม่สามารถสร้างบัตรเลือกตั้งได้",
        variant: "destructive",
      });
      setLoading(false);

      console.error(
        "Error creating ballot proof",
        resCreateBallotProof?.data?.failure || "Unknown error",
      );
      return;
    }

    const proofData = resCreateBallotProof.data.proof;

    if (!proofData.proof) {
      toast({
        title: "เกิดข้อผิดพลาด",
        description: "ไม่สามารถสร้างบัตรเลือกตั้งได้",
        variant: "destructive",
      });
      setLoading(false);

      console.error("Proof is undefined", proofData);
      return;
    }

    const resCreateBallot = await createBallot({
      voteSlug: slug,
      proof: proofData.proof,
    });

    if (resCreateBallot?.data?.failure || !resCreateBallot?.data?.ballot) {
      toast({
        title: "เกิดข้อผิดพลาด",
        description: "ไม่สามารถสร้างบัตรเลือกตั้งได้",
        variant: "destructive",
      });
      setLoading(false);

      console.error(
        "Error creating ballot",
        resCreateBallot?.data?.failure || "Unknown error",
      );
      return;
    }

    const ballotKey = resCreateBallot.data.ballot.ballotKey;

    router.push(`/vote/${slug}/success?ballot_key=${ballotKey}`);
    router.refresh();
  };

  return (
    <>
      <div
        className={cn(
          "flex flex-grow flex-col gap-4 px-5",
          selectedChoice <= 0 && choices.length === 1 ? "pb-32" : "pb-20",
        )}
      >
        <div className="flex flex-shrink flex-col pt-7">
          <h1 className="pr-20 text-3xl font-bold text-neutral-800">{name}</h1>
          <p className="text-neutral-600">{description}</p>
        </div>
        <div
          className={cn(
            "grid flex-grow gap-5 pb-5",
            choicesGridClassName(choices.length),
          )}
        >
          {choices
            .sort((a, b) => a.number - b.number)
            .map((choice, index) => (
              <VoteBallotCard
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
                information={choice.information}
                image={choice.image}
              />
            ))}
        </div>
      </div>
      <TabBarWrapper
        className={
          selectedChoice <= 0 && choices.length === 1 ? "h-32" : "h-20"
        }
      >
        <div className="flex h-full w-full flex-col items-center justify-center gap-3 px-2 py-2">
          <AlertDialog
            open={isOpenConfirmationAlert}
            onOpenChange={setIsOpenConfirmationAlert}
          >
            <AlertDialogTrigger asChild>
              {selectedChoice <= 0 && choices.length === 1 ? (
                <Button
                  variant="outline"
                  size="lg"
                  className="h-full w-full font-semibold"
                  onClick={() => {
                    setSelectedChoice(-1);
                  }}
                >
                  ไม่รับรอง
                </Button>
              ) : null}
            </AlertDialogTrigger>
            <AlertDialogTrigger asChild>
              <Button
                variant={selectedChoice > 0 ? "default" : "secondary"}
                size="lg"
                className="h-full w-full font-semibold"
                onClick={() => {
                  if (selectedChoice <= 0) {
                    setSelectedChoice(0);
                  }
                }}
              >
                {selectedChoice <= 0
                  ? "งดออกเสียง"
                  : `เลือก ${selectedChoiceDisplay}`}
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>
                  ยืนยัน
                  {selectedChoice === 0
                    ? "งดออกเสียง"
                    : selectedChoice === -1
                      ? "ไม่รับรอง"
                      : `เลือก ${selectedChoiceDisplay}`}
                </AlertDialogTitle>
                <AlertDialogDescription>
                  {selectedChoice === 0
                    ? "คุณต้องการงดออกเสียงใช่หรือไม่?"
                    : selectedChoice === -1
                      ? "คุณต้องการไม่รับรองใช่หรือไม?"
                      : `คุณต้องการเลือก ${selectedChoiceData?.name} (${selectedChoiceData?.description})`}
                  <br />
                  หากยืนยันการเลือกแล้ว จะไม่สามารถเปลี่ยนแปลงได้
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel>ยกเลิก</AlertDialogCancel>
                <Button onClick={handleVote} disabled={loading}>
                  {loading ? (
                    <LoaderCircle className="animate-spin" />
                  ) : (
                    "ยืนยัน"
                  )}
                </Button>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        </div>
      </TabBarWrapper>
    </>
  );
};

export default VoteBallotContainer;
