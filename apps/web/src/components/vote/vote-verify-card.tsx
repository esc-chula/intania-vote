"use client";

import { useState } from "react";

import { LoaderCircle } from "lucide-react";
import { verifyBallot } from "~/server/ballot";

import { Button, Input } from "@intania-vote/shadcn";

const VoteVerifyCard: React.FC = () => {
  const [ballotKey, setBallotKey] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const [ballotData, setBallotData] = useState<{
    isValid: boolean;
    choiceNumber: number;
    timestamp: string;
  } | null>(null);

  const verifyBallotHandler = async () => {
    setLoading(true);

    if (!ballotKey) {
      return;
    }

    const res = await verifyBallot({
      ballotKey,
    });

    if (res?.data?.failure || !res?.data?.verifiedBallot) {
      alert("ไม่พบข้อมูลการโหวต");
      return;
    }

    const ballotData = res.data.verifiedBallot;

    setBallotData({
      isValid: ballotData.isValid,
      choiceNumber: ballotData.choiceNumber,
      timestamp: ballotData.timestamp,
    });

    setLoading(false);
  };

  return (
    <div className="flex flex-col gap-3 rounded-3xl border border-neutral-200 bg-neutral-50 p-4">
      <h2 className="text-lg font-semibold text-neutral-700">ตรวจสอบการโหวต</h2>
      <div className="flex gap-2">
        <Input
          placeholder="กรอก Vote Key"
          value={ballotKey}
          onChange={(e) => setBallotKey(e.target.value)}
        />
        <Button
          variant="default"
          disabled={loading}
          onClick={verifyBallotHandler}
        >
          {loading ? <LoaderCircle className="animate-spin" /> : "ตรวจสอบ"}
        </Button>
      </div>
      {ballotData ? (
        <div className="grid grid-cols-2">
          <div>
            <p className="font-bold">ประวัติการโหวต</p>
            <p>
              {ballotData.choiceNumber === 0
                ? "งดออกเสียง"
                : ballotData.choiceNumber === -1
                  ? "ไม่รับรอง"
                  : `เบอร์ ${ballotData.choiceNumber}`}
            </p>
          </div>
          <div>
            <p className="font-bold">เวลา</p>
            <p>
              {new Date(ballotData.timestamp).toLocaleString("th-TH", {
                year: "numeric",
                month: "2-digit",
                day: "2-digit",
                hour: "2-digit",
                minute: "2-digit",
              })}
            </p>
          </div>
        </div>
      ) : (
        <p className="text-center text-xs text-neutral-600">
          กรอก Vote Key ที่ได้รับจากการโหวต เพื่อทำการตรวจสอบผลการโหวตของคุณ
          ระบบจะเข้ารหัสข้อมูลการโหวตไว้
          เพราะฉะนั้นมีเพียงคุณเท่านั้นที่สามารถตรวจการโหวตของคุณเองได้
        </p>
      )}
    </div>
  );
};

export default VoteVerifyCard;
