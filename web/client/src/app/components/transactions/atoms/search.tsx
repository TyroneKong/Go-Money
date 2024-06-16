import React, { SetStateAction } from "react";

type SearchProps = {
  filter: string;
  setFilter: React.Dispatch<SetStateAction<string>>;
};

function Search({ filter, setFilter }: SearchProps) {
  return (
    <input
      className="w-40 h-10"
      value={filter}
      onChange={(e) => setFilter(e.target.value)}
      placeholder=" search..."
    />
  );
}

export default Search;
