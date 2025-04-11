import React from "react";
import Markdown from "react-markdown";

import Image from "next/image";

import { Info } from "lucide-react";
import rehypeRaw from "rehype-raw";
import { choicesGridClassName } from "~/lib/vote";

import { cn } from "@intania-vote/shadcn";
import {
  Drawer,
  DrawerContent,
  DrawerDescription,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from "@intania-vote/shadcn";

interface VoteChoicesInfoProps {
  choices: {
    number: number;
    name: string;
    description: string;
    information?: string;
    image?: string;
  }[];
}

const VoteChoicesInfo: React.FC<VoteChoicesInfoProps> = ({ choices }) => {
  return (
    <div
      className={cn(
        "grid flex-grow gap-5 pb-5",
        choicesGridClassName(choices.length),
      )}
    >
      {choices.map(
        ({ number, name, description, information, image }, index) => (
          <div
            key={index}
            className="relative flex flex-col items-center justify-center gap-4 overflow-hidden rounded-3xl border-2 border-neutral-300 bg-neutral-50 duration-100"
          >
            {image ? (
              <div className="relative aspect-square w-28 select-none overflow-hidden rounded-full">
                <Image src={image} alt={name} fill className="object-cover" />
              </div>
            ) : null}

            <div className="pb-6 text-center text-neutral-900">
              <h2 className="text-xl font-bold">{name}</h2>
              <p className="text-sm opacity-75">{description}</p>
            </div>

            <div className="absolute inset-0 z-10 cursor-pointer" />
            {/* {number ? (
              <span className="absolute -bottom-24 left-0 select-none text-[12rem] font-bold text-black/10">
                {number}
              </span>
            ) : null} */}
            {number ? (
              <span className="absolute bottom-4 left-4 select-none font-semibold text-neutral-500">
                เบอร์ {number}
              </span>
            ) : null}
            <Drawer>
              <DrawerTrigger asChild>
                <div className="absolute bottom-1 right-1 z-20 cursor-pointer p-3">
                  <Info className="text-neutral-500 duration-100" />
                </div>
              </DrawerTrigger>
              <DrawerContent>
                <div className="mx-auto flex h-[80vh] w-full flex-col justify-between overflow-y-auto">
                  <div className="px-4 pt-4">
                    <DrawerHeader className="flex flex-col items-center gap-4">
                      {image ? (
                        <div className="relative aspect-square w-28 select-none overflow-hidden rounded-full">
                          <Image
                            src={image}
                            alt={name}
                            fill
                            className="object-cover"
                          />
                        </div>
                      ) : null}
                      <div className="flex flex-col items-center gap-1 text-center">
                        <DrawerTitle>{name}</DrawerTitle>
                        <DrawerDescription>{description}</DrawerDescription>
                      </div>
                    </DrawerHeader>
                    <div className="rounded-lg border border-neutral-200 bg-neutral-50 p-4">
                      <Markdown
                        rehypePlugins={[rehypeRaw]}
                        components={{
                          h1: ({ node, ...props }) => (
                            <h1 className="text-2xl font-bold" {...props} />
                          ),
                          h2: ({ node, ...props }) => (
                            <h2 className="text-xl font-semibold" {...props} />
                          ),
                          h3: ({ node, ...props }) => (
                            <h3 className="text-lg font-semibold" {...props} />
                          ),
                          p: ({ node, ...props }) => (
                            <p
                              className="text-sm text-neutral-700"
                              {...props}
                            />
                          ),
                          ul: ({ node, ...props }) => (
                            <ul className="list-disc pl-5" {...props} />
                          ),
                          li: ({ node, ...props }) => (
                            <li
                              className="text-sm text-neutral-700"
                              {...props}
                            />
                          ),
                          a: ({ node, ...props }) => (
                            <a className="text-blue-500" {...props} />
                          ),
                          strong: ({ node, ...props }) => (
                            <strong className="font-semibold" {...props} />
                          ),
                          em: ({ node, ...props }) => (
                            <em className="italic" {...props} />
                          ),
                          blockquote: ({ node, ...props }) => (
                            <blockquote
                              className="border-l-4 border-neutral-300 pl-4 italic text-neutral-600"
                              {...props}
                            />
                          ),
                          code: ({ node, ...props }) => (
                            <code
                              className="rounded-md bg-neutral-100 px-1 py-0.5 font-mono text-sm text-neutral-700"
                              {...props}
                            />
                          ),
                          pre: ({ node, ...props }) => (
                            <pre
                              className="rounded-md bg-neutral-100 p-4 font-mono text-sm text-neutral-700"
                              {...props}
                            />
                          ),
                        }}
                      >
                        {information}
                      </Markdown>
                    </div>
                  </div>
                </div>
              </DrawerContent>
            </Drawer>
          </div>
        ),
      )}
    </div>
  );
};

export default VoteChoicesInfo;
