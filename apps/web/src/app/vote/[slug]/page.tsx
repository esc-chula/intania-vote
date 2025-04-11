import { notFound } from "next/navigation";

import XBackButton from "~/components/common/x-back-button";
import VoteBallotContainer from "~/components/vote/vote-ballot-container";
import VoteResultContainer from "~/components/vote/vote-result-container";
import { getSession } from "~/lib/auth";
import { getVoteBySlug, hasUserVoted } from "~/server/vote";

export const dynamic = "force-dynamic";

interface PageProps {
  params: {
    slug: string;
  };
}

const Page: React.FC<PageProps> = async ({ params }) => {
  const session = await getSession();
  const studentId = session?.user?.studentId;
  if (!studentId) {
    return notFound();
  }

  const slug = params.slug;

  const resGetVoteBySlug = await getVoteBySlug({ slug });

  if (resGetVoteBySlug?.data?.failure || !resGetVoteBySlug?.data?.vote) {
    return notFound();
  }

  const voteData = resGetVoteBySlug.data.vote;

  if (!voteData.vote || !voteData.choices) {
    return notFound();
  }

  let isEligible = true;

  if (voteData.vote.eligibleStudentId !== "*") {
    const regex = new RegExp(voteData.vote.eligibleStudentId);
    if (!regex.test(studentId)) {
      isEligible = false;
    }
  }

  const resHasUserVoted = await hasUserVoted({ slug });

  if (resHasUserVoted?.data?.failure || !resHasUserVoted?.data?.hasVoted) {
    return notFound();
  }

  const hasUserVotedData = resHasUserVoted.data.hasVoted;

  return (
    <>
      <XBackButton />
      {hasUserVotedData.hasVoted || !isEligible ? (
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
          isEligible={isEligible}
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
