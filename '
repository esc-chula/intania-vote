import { cn } from "@intania-vote/shadcn";

interface TabBarWrapperProps {
  children: React.ReactNode;
  className?: string;
}

const TabBarWrapper: React.FC<TabBarWrapperProps> = ({
  children,
  className,
}) => {
  return (
    <div
      className={cn(
        "fixed bottom-0 left-0 right-0 h-20 border-t border-neutral-200 bg-neutral-50 p-3",
        className,
      )}
    >
      {children}
    </div>
  );
};

export default TabBarWrapper;
