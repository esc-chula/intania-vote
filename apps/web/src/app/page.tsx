import Link from "next/link";
import { redirect } from "next/navigation";

import RootContainer from "~/components/root/root-container";
import { getSession } from "~/lib/auth";
import { getVotesByUserEligibility } from "~/server/vote";

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

  return (
    <>
      {votesData.map((data) => {
        if (!data.vote || !data.choices) {
          return null;
        }
        return (
          <Link key={data.vote.slug} href={`/vote/${data.vote.slug}`}>
            {data.vote.name}
          </Link>
        );
      })}
      <RootContainer />
    </>
  );
};

export default Page;
