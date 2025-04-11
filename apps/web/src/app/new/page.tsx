import { Bell } from "lucide-react";
import Header from "~/components/common/header";
import Navigation from "~/components/common/navigation";

import { Button } from "@intania-vote/shadcn";

const Page: React.FC = () => {
  return (
    <>
      <Button
        variant="outline"
        size="icon"
        className="fixed right-6 top-5 z-50 h-14 w-14 rounded-full"
      >
        <Bell />
      </Button>
      <Header className="h-24" />
      <div className="mb-20 mt-24 flex flex-grow flex-col items-center justify-center p-5">
        <p className="font-medium text-neutral-300">Coming Soon!</p>
      </div>
      <Navigation />
    </>
  );
};

export default Page;
