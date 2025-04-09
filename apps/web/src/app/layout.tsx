import { getSession } from "~/auth";
import { notoSansThai } from "~/lib/fonts";
import AuthProviders from "~/providers/auth";
import "~/styles/global.css";

import { cn } from "@intania-vote/shadcn";

export const metadata = {
  title: "Intania Vote",
};

const RootLayout = async ({ children }: { children: React.ReactNode }) => {
  const session = await getSession();

  return (
    <html lang="th" className={cn(notoSansThai.variable)}>
      <body className="flex min-h-dvh flex-col">
        <AuthProviders session={session}>{children}</AuthProviders>
      </body>
    </html>
  );
};

export default RootLayout;
