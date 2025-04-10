import { getSession } from "~/lib/auth";
import { notoSansThai } from "~/lib/fonts";
import AuthProviders from "~/providers/auth";
import "~/styles/global.css";

import { cn, Toaster } from "@intania-vote/shadcn";

export const metadata = {
  title: "Intania Vote",
};

const RootLayout = async ({ children }: { children: React.ReactNode }) => {
  const session = await getSession();

  return (
    <html lang="th" className={cn(notoSansThai.variable)}>
      <body className="flex min-h-dvh flex-col">
        <AuthProviders session={session}>
          {children}
          <Toaster />
        </AuthProviders>
      </body>
    </html>
  );
};

export default RootLayout;
