import Header from "~/components/common/header";
import Navigation from "~/components/common/navigation";
import RootContainer from "~/components/root/root-container";
import VoteCard from "~/components/vote/vote-card";
import { getVotes } from "~/server/vote";

const Page: React.FC = async () => {
  const res = await getVotes();

  if (res?.data?.failure || !res?.data?.votes?.votes) {
    return null;
  }

  const votesData = res.data.votes.votes;

  return (
    <>
      <Header className="h-20" />
      <div className="my-20 grid gap-5 p-5 sm:grid-cols-2">
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
      <Navigation />
    </>
  );
};

export default Page;
