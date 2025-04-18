import { Noto_Sans_Thai as NotoSansThai } from "next/font/google";

export const notoSansThai = NotoSansThai({
  subsets: ["latin", "thai"],
  weight: ["100", "200", "300", "400", "500", "600", "700", "800", "900"],
  variable: "--font-noto-sans-thai",
});
