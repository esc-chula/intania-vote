"use client";

import { signIn } from "next-auth/react";

import { createVote } from "~/server/vote";

import { Button } from "@intania-vote/shadcn";

const RootContainer: React.FC = () => {
  return (
    <div className="flex flex-col gap-4">
      <div>
        <Button
          onClick={() => {
            signIn("intania");
          }}
        >
          Login with Intania
        </Button>
      </div>
      <div>
        <Button
          onClick={() => {
            createVote({
              name: "เลือกตั้งหัวหน้านิสิต กวศ. 68",
              description: "Test Description",
              slug: "esc-68-president",
              owner: "ESC",
              eligibleStudentId: "*",
              eligibleDepartment: "*",
              eligibleYear: "*",
              isPrivate: false,
              isRealTime: false,
              isAllowMultiple: false,
              startAt: new Date("2025-04-11T10:00:00+07:00"),
              endAt: new Date("2025-04-11T18:00:00+07:00"),
              choices: [
                {
                  number: 1,
                  name: "ปราชญ์ธนบดี คนไว",
                  description: "พรรค เลิกกั๊กแล้วรักกวศ.",
                  information: "",
                  image: "/assets/images/esc-68-president-1.jpg",
                },
                {
                  number: 2,
                  name: "วรัญยชญ์ ข่ายม่าน",
                  description: "พรรค Avenginia",
                  information: "",
                  image: "/assets/images/esc-68-president-2.jpg",
                },
              ],
            }).catch((err) => {
              console.error(err);
            });

            createVote({
              name: "เลือกตั้งหัวหน้าชั้นปี 4 ปีการศึกษา 2568",
              description: "Test Description",
              slug: "year-4-68-president",
              owner: "ESC",
              eligibleStudentId: "*",
              eligibleDepartment: "*",
              eligibleYear: "*",
              isPrivate: false,
              isRealTime: false,
              isAllowMultiple: false,
              startAt: new Date("2025-04-11T10:00:00+07:00"),
              endAt: new Date("2025-04-11T18:00:00+07:00"),
              choices: [
                {
                  number: 1,
                  name: "กฤติพงษ์ เหลืองนิยมกุล",
                  description: "Kneck",
                  information: "",
                  image: "/assets/images/year-4-68-class-leader-1.jpg",
                },
                {
                  number: 2,
                  name: "เนื้อบุญญ์ กระจ่างวงษ์",
                  description: "เนื้อบุญญ์",
                  information: "",
                  image: "/assets/images/year-4-68-class-leader-2.jpg",
                },
              ],
            }).catch((err) => {
              console.error(err);
            });

            createVote({
              name: "เลือกตั้งหัวหน้าชั้นปี 2 ปีการศึกษา 2568",
              description: "Test Description",
              slug: "year-2-68-president",
              image: "/mock.jpg",
              owner: "ESC",
              eligibleStudentId: "*",
              eligibleDepartment: "*",
              eligibleYear: "*",
              isPrivate: false,
              isRealTime: false,
              isAllowMultiple: false,
              startAt: new Date("2025-04-11T10:00:00+07:00"),
              endAt: new Date("2025-04-11T18:00:00+07:00"),
              choices: [
                {
                  number: 1,
                  name: "คุณากร เชาวนวิรัตน์",
                  description: "ก้อง",
                  information: "",
                  image: "/assets/images/year-2-68-class-leader-1.jpg",
                },
              ],
            }).catch((err) => {
              console.error(err);
            });
          }}
        >
          Create Vote
        </Button>
      </div>
    </div>
  );
};

export default RootContainer;
