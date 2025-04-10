import { redirect } from "next/navigation";

import Header from "~/components/common/header";
import RootContainer from "~/components/root/root-container";
import VoteCard from "~/components/vote/vote-card";
import { getSession } from "~/lib/auth";
import { getVotesByUserEligibility } from "~/server/vote";

const Page: React.FC = async () => {
  const session = await getSession();

  if (!session) {
    return redirect("/api/auth/signin");
  }

  const res = await getVotesByUserEligibility();

  if (res?.data?.failure || !res?.data?.votes?.votes.length) {
    return redirect("/home");
  }

  const votesData = res.data.votes.votes;

  // if (votesData.length === 1 && votesData[0].vote) {
  //   return redirect(`/vote/${votesData[0].vote.slug}`);
  // }

  return (
    <>
      <Header className="h-24" />
      <div className="mt-24 grid gap-5 p-5 sm:grid-cols-2">
        {votesData.map((data) => {
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
              startAt={new Date(data.vote.startAt)}
              endAt={new Date(data.vote.endAt)}
              choices={data.choices.map((choice) => ({
                name: choice.name,
              }))}
            />
          );
        })}
        <RootContainer />
      </div>
    </>
  );
};

export default Page;
