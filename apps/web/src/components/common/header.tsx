import { cn } from "@intania-vote/shadcn";

interface HeaderProps {
  children?: React.ReactNode;
  className?: string;
}

const Header: React.FC<HeaderProps> = ({ children, className }) => {
  return (
    <header
      className={cn(
        "fixed left-0 right-0 top-0 z-40 flex items-center justify-between border-b border-neutral-200 bg-neutral-50 px-4 py-6",
        className,
      )}
    >
      <span className="text-2xl font-bold text-neutral-800 sm:text-3xl">
        Intania{" "}
        <span className="decoration-primary underline decoration-[2.5px]">
          Vote
        </span>
      </span>
      {children}
    </header>
  );
};

export default Header;
