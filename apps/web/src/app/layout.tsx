import { notoSansThai } from "~/lib/fonts";
import "~/styles/global.css";

import { cn } from "@intania-vote/shadcn";

export const metadata = {
  title: "Intania Vote",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="th" className={cn(notoSansThai.variable)}>
      <body className="flex min-h-dvh flex-col">{children}</body>
    </html>
  );
}
