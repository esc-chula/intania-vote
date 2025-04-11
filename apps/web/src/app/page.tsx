import Link from "next/link";
import { redirect } from "next/navigation";

import Header from "~/components/common/header";
import VoteCard from "~/components/vote/vote-card";
import { getSession } from "~/lib/auth";
import { getVotesByUserEligibility } from "~/server/vote";

import { Button } from "@intania-vote/shadcn";

const Page: React.FC = async () => {
  const session = await getSession();

  if (!session) {
    return redirect("/api/auth/signin");
  }

  const res = await getVotesByUserEligibility();

  if (res?.data?.failure || !res?.data?.votes?.votes) {
    return redirect("/home");
  }

  const votesData = res.data.votes.votes;

  if (votesData.length === 0) {
    return redirect("/home");
  }

  return (
    <>
      <Header className="h-20">
        <div>
          <h1 className="text-lg font-medium text-neutral-700">โหวตของคุณ</h1>
        </div>
      </Header>
      <div className="flex min-h-dvh flex-col items-center justify-between gap-3.5 px-5 pb-5 pt-24">
        <div className="grid max-w-screen-md gap-5 sm:grid-cols-2">
          {votesData
            .sort((a, b) => {
              if (!a.vote || !b.vote) return 0;
              return (
                new Date(a.vote.startAt).getTime() -
                new Date(b.vote.startAt).getTime()
              );
            })
            .map((data) => {
              if (!data.vote || !data.choices) {
                return null;
              }

              let owner = "Unknown";

              switch (data.vote.owner.toString()) {
                case "0":
                  owner = "User";
                  break;
                case "1":
                  owner = "ESC";
                  break;
                case "2":
                  owner = "ISESC";
                  break;
                default:
                  owner = "Unknown";
                  break;
              }

              return (
                <VoteCard
                  key={data.vote.slug}
                  name={data.vote.name}
                  slug={data.vote.slug}
                  image={data.vote.image}
                  owner={owner}
                  isEligible={true}
                  startAt={new Date(data.vote.startAt)}
                  endAt={new Date(data.vote.endAt)}
                  choices={data.choices.map((choice) => ({
                    name: choice.name,
                  }))}
                />
              );
            })}
        </div>
        <Link href="/home">
          <Button variant="outline" size="lg">
            ไปหน้าหลัก
          </Button>
        </Link>
      </div>
    </>
  );
};

export default Page;
