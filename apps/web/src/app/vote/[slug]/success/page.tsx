import Link from "next/link";
import { notFound } from "next/navigation";

import { Share, Vote } from "lucide-react";
import CopyButton from "~/components/common/copy-button";
import Header from "~/components/common/header";
import XBackButton from "~/components/common/x-back-button";
import { verifyBallot } from "~/server/ballot";
import { getVoteBySlug } from "~/server/vote";

import { Button, Input } from "@intania-vote/shadcn";

export const dynamic = "force-dynamic";

interface SearchParamProps {
  ballot_key?: string;
}

interface PageProps {
  params: {
    slug: string;
  };
  searchParams: SearchParamProps;
}

const Page: React.FC<PageProps> = async ({ params, searchParams }) => {
  const slug = params.slug;

  const resGetVoteBySlug = await getVoteBySlug({ slug });

  if (resGetVoteBySlug?.data?.failure || !resGetVoteBySlug?.data?.vote) {
    return notFound();
  }

  const voteData = resGetVoteBySlug.data.vote;

  if (!voteData.vote || !voteData.choices) {
    return notFound();
  }

  const ballotKey = searchParams.ballot_key;
  if (!ballotKey) {
    return notFound();
  }

  const resVerifyBallot = await verifyBallot({ ballotKey });

  if (
    resVerifyBallot?.data?.failure ||
    !resVerifyBallot?.data?.verifiedBallot
  ) {
    return notFound();
  }

  const verifiedData = resVerifyBallot.data.verifiedBallot;

  return (
    <>
      <XBackButton href="/" />
      <Header className="h-24" />
      <div className="flex flex-grow flex-col items-center justify-between gap-10 p-5">
        <div className="mt-24 flex flex-grow flex-col items-center justify-center gap-10 text-center">
          <div className="flex flex-col items-center">
            <Vote size={80} className="text-primary" />
            <h1 className="text-primary text-4xl font-bold">โหวตสำเร็จ</h1>
            <p className="font-semibold text-neutral-700">
              {voteData.vote.name}
            </p>
            <p className="mt-2 font-medium text-neutral-500">
              “
              {verifiedData.choiceNumber === 0
                ? "งดออกเสียง"
                : verifiedData.choiceNumber === -1
                  ? "ไม่รับรอง"
                  : `หมายเลข ${verifiedData.choiceNumber}`}
              ”
            </p>
          </div>
          <div className="fle-col flex flex-col items-center gap-4">
            <Link href="/">
              <Button variant="default" size="lg">
                กลับหน้าแรก
              </Button>
            </Link>
            <Button variant="outline" size="lg">
              <Share />
              แชร์
            </Button>
          </div>
        </div>
        <div className="flex w-full max-w-sm flex-col items-center gap-3 rounded-2xl border border-neutral-200 bg-neutral-50 px-2 py-4 text-center">
          <h3 className="font-semibold text-neutral-800">Vote Key</h3>
          <div className="flex items-center gap-2.5">
            <Input value={ballotKey} className="text-sm" />
            <CopyButton value={ballotKey} />
          </div>
          <p className="text-xs text-neutral-600">
            หากต้องการตรวจสอบการโหวตของตัวเอง
            <br />
            สามารถเก็บ Vote Key นี้ไว้ตรวจสอบในหน้าแสดงผลการโหวตได้
          </p>
        </div>
      </div>
    </>
  );
};

export default Page;
