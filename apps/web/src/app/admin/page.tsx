import { notFound } from "next/navigation";

import RootContainer from "~/components/root/root-container";
import { getSession } from "~/lib/auth";

const Page: React.FC = async () => {
  const session = await getSession();

  if (session?.user.studentId !== "6538068821") {
    return notFound();
  }

  return (
    <div className="my-20 grid gap-5 p-5 sm:grid-cols-2">
      <RootContainer />
    </div>
  );
};

export default Page;
