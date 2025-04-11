import Image from "next/image";
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
    <>
      <div className="flex min-h-svh flex-col items-center justify-between p-5">
        <div className="flex flex-grow flex-col items-center justify-center gap-6">
          <div className="relative aspect-square w-12">
            <Image
              src="/assets/esc-icon.svg"
              alt="esc"
              fill
              className="object-contain"
            />
          </div>
          <div className="flex flex-col items-center gap-2">
            <span className="text-5xl font-bold text-neutral-800 sm:text-3xl">
              Intania{" "}
              <span className="decoration-primary underline decoration-[4.5px]">
                Vote
              </span>
            </span>
            <p className="text-neutral-600">ระบบโหวตออนไลน์ Intania</p>
          </div>
        </div>
        <div className="flex w-full max-w-sm flex-col items-center justify-center gap-4 rounded-3xl border border-neutral-200 bg-neutral-50 p-6">
          <h1 className="text-lg font-bold text-neutral-800">เข้าสู่ระบบ</h1>
          <SignInButton />
          <p className="text-sm text-neutral-600">
            ใช้รหัสนิสิตและรหัสผ่าน CUNET
          </p>
        </div>
      </div>
    </>
  );
};

export default Page;
