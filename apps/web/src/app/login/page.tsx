import { redirect } from "next/navigation";

import SignInButton from "~/components/common/sign-in-button";
import { getSession } from "~/lib/auth";

interface SearchParamProps {
  callbackUrl?: string;
}

interface PageProps {
  searchParams: SearchParamProps;
}

const Page: React.FC<PageProps> = async ({ searchParams }) => {
  const callbackUrl = searchParams.callbackUrl || "/";

  const session = await getSession();

  if (session?.user) {
    return redirect(callbackUrl);
  }

  return (
    <div>
      <SignInButton />
    </div>
  );
};

export default Page;
