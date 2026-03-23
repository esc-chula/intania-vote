import Image from "next/image";

const Page: React.FC = () => {
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
            <p className="text-neutral-600 font-medium">ระบบโหวตออนไลน์ Intania</p>
          </div>
        </div>
        <div className="flex w-full max-w-sm flex-col items-center justify-center text-center pb-24">
          <h1 className="text-xl font-semibold text-neutral-800 uppercase tracking-tight">Coming Soon ...</h1>
        </div>
      </div>
    </>
  );
};

export default Page;
