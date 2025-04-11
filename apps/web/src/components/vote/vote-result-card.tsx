"use client";

import { useEffect, useState } from "react";

import { calculateTime } from "~/lib/vote";

import { Tally } from "@intania-vote/grpc-ts";

interface VoteResultCardProps {
  startAt: Date;
  endAt: Date;
  choices: {
    number: number;
    name: string;
    description: string;
    information?: string;
    image?: string;
  }[];
  tally?: Tally;
}

const VoteResultCard: React.FC<VoteResultCardProps> = ({
  startAt,
  endAt,
  choices,
  tally,
}) => {
  const [time, setTime] = useState<string | null>(null);

  useEffect(() => {
    setTime(calculateTime(startAt, endAt));

    const interval = setInterval(() => {
      const calculatedTime = calculateTime(startAt, endAt);
      setTime(calculatedTime);
    }, 1000);

    return () => clearInterval(interval);
  }, [endAt, startAt]);

  const now = new Date();

  return (
    <div className="relative flex flex-grow flex-col gap-3 overflow-hidden rounded-3xl border border-neutral-200 bg-neutral-50 p-4">
      <h2 className="text-lg font-semibold text-neutral-700">ผลการโหวต</h2>
      {now > endAt && tally ? (
        <>
          {choices.map((choice) => (
            <div className="flex h-8 w-full justify-between rounded-full border border-neutral-200 bg-white">
              <div
                className="bg-primary flex h-full min-w-max items-center justify-start rounded-full px-4 text-sm font-semibold text-white"
                style={{
                  width: `${
                    ((tally.choices.find((c) => c.number === choice.number)
                      ?.count ?? 0) *
                      100) /
                    (tally.choices.reduce((acc, cur) => acc + cur.count, 0) ||
                      1)
                  }%`,
                }}
              >
                เบอร์ {choice.number}
              </div>
              <div className="flex h-full min-w-max items-center justify-center px-4 text-xs">
                {tally.choices.find((c) => c.number === choice.number)?.count ??
                  0}{" "}
                Votes
              </div>
            </div>
          ))}
          <div className="grid grid-cols-2 rounded-2xl border border-neutral-200 bg-white p-3 text-sm">
            <p className="font-bold text-neutral-800">ข้อมูล</p>
            <p className="font-bold text-neutral-800">จำนวน</p>
            <p className="text-neutral-600">งดออกเสียง</p>
            <p className="text-neutral-600">
              {tally.choices.find((c) => c.number === 0)?.count ?? 0}
            </p>
            <p className="text-neutral-600">ไม่รับรอง</p>
            <p className="text-neutral-600">
              {tally.choices.find((c) => c.number === -1)?.count ?? 0}
            </p>
            <p className="text-neutral-600">เลือกโหวต</p>
            <p className="text-neutral-600">
              {tally.choices
                .filter((c) => c.number > 0)
                .reduce((acc, cur) => acc + cur.count, 0)}
            </p>
            <p className="text-neutral-600">รวม</p>
            <p className="text-neutral-600">
              {tally.choices.reduce((acc, cur) => acc + cur.count, 0)}
            </p>
          </div>
        </>
      ) : (
        <MockVoteResult />
      )}

      {now > endAt ? null : (
        <div className="absolute bottom-0 left-0 right-0 top-9 flex flex-col items-center justify-center gap-4 text-center backdrop-blur-sm">
          {time ? (
            <div className="flex items-center gap-1.5 rounded-full border border-neutral-200 bg-white py-1 pl-2.5 pr-3 text-xs font-bold text-neutral-700 shadow-xl">
              <span className="aspect-square w-2 animate-pulse rounded-full bg-red-500" />
              <span>{time}</span>
            </div>
          ) : null}
          <p className="text-xs text-neutral-600">
            การโหวตของคุณได้ถูกบันทึกแล้ว
            <br />
            ในอีกไม่นาน หน้านี้ก็จะแสดงผลลัพธ์ของการโหวต
          </p>
        </div>
      )}
    </div>
  );
};

export default VoteResultCard;

const MockVoteResult: React.FC = () => {
  return (
    <>
      <div className="h-8 w-full rounded-full border border-neutral-200 bg-white">
        <div
          className="bg-primary h-full rounded-full px-4"
          style={{
            width: "40%",
          }}
        ></div>
      </div>
      <div className="h-8 w-full rounded-full border border-neutral-200 bg-white">
        <div
          className="bg-primary h-full rounded-full px-4"
          style={{
            width: "60%",
          }}
        ></div>
      </div>
      <div className="grid grid-cols-2 rounded-2xl border border-neutral-200 bg-white p-3 text-sm">
        <p className="font-bold text-neutral-800">ข้อมูล</p>
        <p className="font-bold text-neutral-800">จำนวน</p>
        <p className="text-neutral-600">งดออกเสียง</p>
        <p className="text-neutral-600">XX</p>
        <p className="text-neutral-600">ไม่รับรอง</p>
        <p className="text-neutral-600">XX</p>
        <p className="text-neutral-600">เลือกโหวต</p>
        <p className="text-neutral-600">XXX</p>
        <p className="text-neutral-600">รวม</p>
        <p className="text-neutral-600">XXX</p>
      </div>
    </>
  );
};
