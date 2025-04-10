"use client";

import { useState } from "react";

import { useRouter } from "next/navigation";

import { LoaderCircle } from "lucide-react";
import TabBarWrapper from "~/components/common/tab-bar-wrapper";
import VoteBallotCard from "~/components/vote/vote-ballot-card";
import { createBallot, createBallotProof } from "~/server/ballot";

import { Button, useToast } from "@intania-vote/shadcn";
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

interface VoteContainerProps {
  name: string;
  description: string;
  slug: string;
  choices: {
    number?: string;
    name: string;
    description: string;
    information?: string;
    image?: string;
  }[];
}

const VoteContainer: React.FC<VoteContainerProps> = ({
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
      choiceId: selectedChoice,
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
  };

  return (
    <>
      <div className="flex flex-grow flex-col gap-4 px-5 pb-20">
        <div className="flex flex-shrink flex-col pt-7">
          <h1 className="pr-20 text-3xl font-bold text-neutral-800">{name}</h1>
          <p className="text-neutral-600">{description}</p>
        </div>
        <div className="grid flex-grow grid-cols-1 grid-rows-2 gap-5 pb-5">
          {choices.map((choice, index) => (
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
      <TabBarWrapper>
        <div className="flex h-full w-full flex-col items-center justify-center px-2 py-2">
          <AlertDialog
            open={isOpenConfirmationAlert}
            onOpenChange={setIsOpenConfirmationAlert}
          >
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

export default VoteContainer;
