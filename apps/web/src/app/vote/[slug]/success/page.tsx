import { verifyBallot } from "~/server/ballot";

interface SearchParamProps {
  ballot_key?: string;
}

interface PageProps {
  params: {
    slug: string;
  };
  searchParams: SearchParamProps;
}

const Page: React.FC<PageProps> = async ({ params, searchParams }) => {
  const ballotKey = searchParams.ballot_key;
  if (!ballotKey) {
    return <div>No ballot key provided</div>;
  }

  const res = await verifyBallot({ ballotKey });

  if (res?.data?.failure || !res?.data?.verifiedBallot) {
    return <div>Invalid ballot key</div>;
  }

  const verifiedData = res.data.verifiedBallot;

  console.log(verifiedData);

  return <div>{searchParams.ballot_key}</div>;
};

export default Page;
