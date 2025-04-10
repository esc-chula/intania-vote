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
  return <div>{searchParams.ballot_key}</div>;
};

export default Page;
