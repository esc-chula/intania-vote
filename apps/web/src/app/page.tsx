import Header from "~/components/common/header";
import Navigation from "~/components/common/navigation";

import { Button } from "@intania-vote/shadcn";

export default function Home() {
  return (
    <>
      <Header />
      <div className="flex h-screen w-full flex-col items-center justify-center">
        ทดสอบ
        <Button>Click me</Button>
      </div>
      <Navigation />
    </>
  );
}
