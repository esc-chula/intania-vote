const Header: React.FC = () => {
  return (
    <header className="fixed left-0 right-0 top-0 flex items-center justify-between border-b border-neutral-200 bg-neutral-50 p-4">
      <span className="text-2xl font-bold text-neutral-800 sm:text-3xl">
        Intania{" "}
        <span className="decoration-primary underline decoration-[2.5px]">
          Vote
        </span>
      </span>
    </header>
  );
};

export default Header;
