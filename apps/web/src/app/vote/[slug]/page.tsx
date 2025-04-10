import Link from "next/link";
import { notFound } from "next/navigation";

import { X } from "lucide-react";
import VoteContainer from "~/components/vote/vote-container";
import { getVoteBySlug } from "~/server/vote";

import { Button } from "@intania-vote/shadcn";

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
      <Link href="/">
        <Button
          variant="outline"
          size="icon"
          className="fixed right-6 top-5 h-14 w-14 rounded-full"
        >
          <X />
        </Button>
      </Link>
      <VoteContainer
        name={voteData.vote.name}
        description={voteData.vote.description}
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
