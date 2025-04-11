"use client";

import { useEffect, useRef, useState } from "react";

import Image from "next/image";

import { Link, Vote } from "lucide-react";

import { cn } from "@intania-vote/shadcn";

interface StoryProps {
  className?: string;
  voteName: string;
}

const Story: React.FC<StoryProps> = ({ className, voteName }) => {
  const ref = useRef<HTMLDivElement>(null);

  const [fontFactor, setFontFactor] = useState(0);

  useEffect(() => {
    if (!ref.current) {
      return;
    }

    const handleResize = (): void => {
      setFontFactor((ref.current?.offsetWidth ?? 0) / 30);
    };

    handleResize();

    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  return (
    <div
      ref={ref}
      className={cn(
        "relative aspect-[9/16] select-none border bg-white",
        className,
      )}
    >
      <div className="absolute inset-0 z-20 mb-[30%]">
        <div className="flex h-full w-full flex-col items-center justify-center text-center">
          <div className="flex flex-col items-center">
            <Vote size={fontFactor * 5.4} className="text-primary" />
            <h1
              className="text-primary font-bold"
              style={{
                fontSize: `${fontFactor * 2.9}px`,
              }}
            >
              โหวตสำเร็จ
            </h1>
            <p
              className="font-semibold text-neutral-700"
              style={{
                fontSize: `${fontFactor * 1.4}px`,
              }}
            >
              {voteName}
            </p>
          </div>
          <div
            className="bg-primary flex items-center text-white"
            style={{
              borderRadius: `${fontFactor * 1}px`,
              fontSize: `${fontFactor * 1.2}px`,
              padding: `${fontFactor * 0.75}px ${fontFactor * 1.25}px`,
              marginTop: `${fontFactor * 1.5}px`,
            }}
          >
            <Link
              size={fontFactor * 1.5}
              style={{
                marginRight: `${fontFactor * 0.5}px`,
              }}
            />
            <p className="font-medium">vote.intania.org</p>
          </div>
        </div>
      </div>
      <div
        className="absolute left-1/2 aspect-square -translate-x-1/2"
        style={{
          top: `${fontFactor * 2}px`,
          width: `${fontFactor * 2}px`,
        }}
      >
        <Image
          src="/assets/esc-icon.svg"
          alt="esc"
          fill
          className="object-contain"
          sizes="100%"
        />
      </div>
      <div
        className="absolute left-1/2 aspect-square -translate-x-1/2"
        style={{
          bottom: `-${fontFactor * 0.6}px`,
          width: `${fontFactor * 8}px`,
        }}
      >
        <Image
          src="/assets/intania-vote.svg"
          alt="esc"
          fill
          className="object-contain"
          sizes="100%"
        />
      </div>
    </div>
  );
};

export default Story;
