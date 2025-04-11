import { Bell } from "lucide-react";
import Header from "~/components/common/header";
import Navigation from "~/components/common/navigation";
import VoteCard from "~/components/vote/vote-card";
import { getSession } from "~/lib/auth";
import { getVotes } from "~/server/vote";

import { Button } from "@intania-vote/shadcn";

const Page: React.FC = async () => {
  const session = await getSession();
  const studentId = session?.user?.studentId;
  if (!studentId) {
    return null;
  }

  const res = await getVotes();

  if (res?.data?.failure || !res?.data?.votes?.votes) {
    return null;
  }

  const votesData = res.data.votes.votes;

  return (
    <>
      <Button
        variant="outline"
        size="icon"
        className="fixed right-6 top-5 z-50 h-14 w-14 rounded-full"
      >
        <Bell />
      </Button>
      <Header className="h-24" />
      <div className="flex min-h-dvh flex-col items-center justify-between gap-3.5 px-5 pb-24 pt-28">
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

              let isEligible = false;

              // test regex with data.vote.eligibleStudentId with session.user.studentId, eligibleStudentId is a regex string and can be * to match any string
              if (data.vote.eligibleStudentId === "*") {
                isEligible = true;
              } else {
                const regex = new RegExp(data.vote.eligibleStudentId);
                isEligible = regex.test(studentId);
              }

              return (
                <VoteCard
                  key={data.vote.slug}
                  name={data.vote.name}
                  slug={data.vote.slug}
                  image={data.vote.image}
                  owner={owner}
                  isEligible={isEligible}
                  startAt={new Date(data.vote.startAt)}
                  endAt={new Date(data.vote.endAt)}
                  choices={data.choices.map((choice) => ({
                    name: choice.name,
                  }))}
                />
              );
            })}
        </div>
      </div>
      <Navigation />
    </>
  );
};

export default Page;
