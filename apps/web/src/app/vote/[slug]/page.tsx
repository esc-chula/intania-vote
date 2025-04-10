import { notFound } from "next/navigation";

import XBackButton from "~/components/common/x-back-button";
import VoteContainer from "~/components/vote/vote-container";
import { getVoteBySlug } from "~/server/vote";

interface PageProps {
  params: {
    slug: string;
  };
}

const Page: React.FC<PageProps> = async ({ params }) => {
  const slug = params.slug;

  const res = await getVoteBySlug({ slug });

  if (res?.data?.failure || !res?.data?.vote) {
    return notFound();
  }

  const voteData = res.data.vote;

  if (!voteData.vote || !voteData.choices) {
    return notFound();
  }

  return (
    <>
      <XBackButton />
      <VoteContainer
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
    </>
  );
};

export default Page;
