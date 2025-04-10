"use client";

import { useRouter } from "next/navigation";

import { X } from "lucide-react";

import { Button } from "@intania-vote/shadcn";

interface XBackButtonProps {
  href?: string;
}

const XBackButton: React.FC<XBackButtonProps> = ({ href }) => {
  const router = useRouter();

  return (
    <Button
      variant="outline"
      size="icon"
      className="fixed right-6 top-5 h-14 w-14 rounded-full"
      onClick={() => {
        if (href) {
          router.push(href);
        } else {
          router.back();
        }
      }}
    >
      <X />
    </Button>
  );
};

export default XBackButton;
