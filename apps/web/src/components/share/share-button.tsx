"use client";

import { useRef } from "react";

import html2canvas from "html2canvas-pro";
import { Share } from "lucide-react";

import { Button, useToast } from "@intania-vote/shadcn";

import Story from "./story";

interface ShareButtonProps {
  isMobile: boolean;
  voteName: string;
}

const ShareButton: React.FC<ShareButtonProps> = ({ isMobile, voteName }) => {
  const { toast } = useToast();

  const storyRef = useRef<HTMLDivElement>(null);
  const handleDownload = (): void => {
    if (!storyRef.current) {
      return;
    }

    html2canvas(storyRef.current)
      .then((canvas) => {
        const link = document.createElement("a");
        link.href = canvas.toDataURL();
        link.download = "image.png";
        link.click();
      })
      .catch((error: unknown) => {
        toast({
          title: "Error sharing story",
          description:
            error instanceof Error ? error.message : "An error occurred",
          variant: "destructive",
        });
      });
  };

  const handleShare = async (): Promise<void> => {
    try {
      if (!storyRef.current) {
        return;
      }

      if (!isMobile) {
        handleDownload();
        return;
      }

      const dataUrl = await html2canvas(storyRef.current).then((canvas) => {
        const dataUrl = canvas.toDataURL();
        return dataUrl;
      });

      const dataBlob = await (await fetch(dataUrl)).blob();

      const image = new File([dataBlob], "intania_vote_story.png", {
        type: dataBlob.type,
      });

      const shareData: ShareData = {
        title: "Intania Vote",
        text: "เชิญชวนทุกคนร่วมกันใช้เสียงเลือกตั้ง",
        files: [image],
      };

      await navigator.share(shareData);
    } catch (error) {
      console.error(error);
      handleDownload();
    }
  };
  return (
    <>
      <div ref={storyRef} className="absolute bottom-[100vh]">
        <Story className="w-[1080px]" voteName={voteName} />
      </div>

      <Button variant="outline" size="lg" onClick={handleDownload}>
        <Share />
        แชร์
      </Button>
    </>
  );
};

export default ShareButton;
