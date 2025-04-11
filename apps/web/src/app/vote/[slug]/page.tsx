import { notFound } from "next/navigation";

import XBackButton from "~/components/common/x-back-button";
import VoteBallotContainer from "~/components/vote/vote-ballot-container";
import VoteResultContainer from "~/components/vote/vote-result-container";
import { getVoteBySlug, hasUserVoted } from "~/server/vote";

export const dynamic = "force-dynamic";

interface PageProps {
  params: {
    slug: string;
  };
}

const Page: React.FC<PageProps> = async ({ params }) => {
  const slug = params.slug;

  const resGetVoteBySlug = await getVoteBySlug({ slug });

  if (resGetVoteBySlug?.data?.failure || !resGetVoteBySlug?.data?.vote) {
    return notFound();
  }

  const voteData = resGetVoteBySlug.data.vote;

  if (!voteData.vote || !voteData.choices) {
    return notFound();
  }

  const resHasUserVoted = await hasUserVoted({ slug });

  if (resHasUserVoted?.data?.failure || !resHasUserVoted?.data?.hasVoted) {
    return notFound();
  }

  const hasUserVotedData = resHasUserVoted.data.hasVoted;

  return (
    <>
      <XBackButton />
      {hasUserVotedData.hasVoted ? (
        <VoteResultContainer
          name={voteData.vote.name}
          description={voteData.vote.description}
          slug={voteData.vote.slug}
          startAt={new Date(voteData.vote.startAt)}
          endAt={new Date(voteData.vote.endAt)}
          choices={voteData.choices.map((choice) => ({
            number: choice.number,
            name: choice.name,
            description: choice.description,
            information: choice.information,
            image: choice.image,
          }))}
        />
      ) : (
        <VoteBallotContainer
          name={voteData.vote.name}
          description={voteData.vote.description}
          slug={voteData.vote.slug}
          choices={voteData.choices.map((choice) => ({
            number: choice.number,
            name: choice.name,
            description: choice.description,
            information: choice.information,
            image: choice.image,
          }))}
        />
      )}
    </>
  );
};

export default Page;
