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
              image: "/mock.jpg",
              owner: "ESC",
              eligibleStudentId: "*",
              eligibleDepartment: "*",
              eligibleYear: "*",
              isPrivate: false,
              isRealTime: false,
              isAllowMultiple: false,
              startAt: new Date(),
              endAt: new Date(new Date().getTime() + 1000 * 60 * 60),
              choices: [
                {
                  number: 1,
                  name: "ปราชญ์ธนบดี คนไว",
                  description: "พรรค เลิกกั๊กแล้วรักกวศ.",
                  information: "# Test Information",
                  image: "/mock.jpg",
                },
                {
                  number: 2,
                  name: "วรัญยชญ์ ข่ายม่าน",
                  description: "พรรค Avenginia",
                  information: "# Test Information",
                  image: "/mock.jpg",
                },
              ],
            }).catch((err) => {
              console.error(err);
            });

            createVote({
              name: "เลือกตั้งหัวหน้าชั้นปี 4 ปีการศึกษา 2568",
              description: "Test Description",
              slug: "year-4-68-president",
              image: "/mock.jpg",
              owner: "ESC",
              eligibleStudentId: "*",
              eligibleDepartment: "*",
              eligibleYear: "*",
              isPrivate: false,
              isRealTime: false,
              isAllowMultiple: false,
              startAt: new Date(),
              endAt: new Date(new Date().getTime() + 1000 * 60 * 60),
              choices: [
                {
                  number: 1,
                  name: "กฤติพงษ์ เหลืองนิยมกุล",
                  description: "Kneck",
                  information: "# Test Information",
                  image: "/mock.jpg",
                },
                {
                  number: 2,
                  name: "เนื้อบุญญ์ กระจ่างวงษ์",
                  description: "เนื้อบุญญ์",
                  information: "# Test Information",
                  image: "/mock.jpg",
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
              startAt: new Date(),
              endAt: new Date(new Date().getTime() + 1000 * 60 * 60),
              choices: [
                {
                  number: 1,
                  name: "คุณากร เชาวนวิรัตน์",
                  description: "ก้อง",
                  information: "# Test Information",
                  image: "/mock.jpg",
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
