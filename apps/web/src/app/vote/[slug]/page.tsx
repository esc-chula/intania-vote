import Link from "next/link";

import { X } from "lucide-react";
import VoteContainer from "~/components/vote/vote-container";

import { Button } from "@intania-vote/shadcn";

const Page: React.FC = () => {
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
        name="Vote 1"
        description="Vote Description"
        choices={[
          {
            number: "1",
            name: "Choice 1",
            description: "Description 1",
            image: "/mock.jpg",
          },
          {
            number: "2",
            name: "Choice 2",
            description: "Description 2",
            image: "/mock.jpg",
          },
        ]}
      />
    </>
  );
};

export default Page;
