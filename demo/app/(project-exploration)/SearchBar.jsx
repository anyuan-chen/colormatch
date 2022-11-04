"use client";
import React from "react";
import { Inter } from "@next/font/google";
import Image from "next/image";
const inter = Inter();
const SearchBar = ({ text, setText }) => {
  return (
    <>
      <div className="search-bar">
        <Image
          width={24}
          height={24}
          src="/icons/search.svg"
          alt="search"
        ></Image>
        <input
          type="text"
          className={`input ${inter.className}`}
          placeholder="search..."
          value={text}
          onChange={(event) => {
            setText(event.target.value);
          }}
        ></input>
      </div>

      <style jsx>{styles}</style>
    </>
  );
};
const styles = `
    .search-bar {
        display: flex;
        border: 1px solid black;
        border-radius: 16px;
        padding: 4px 12px;
        column-gap: 8px;
    }
    .input {
        border: none;
    }
    .input::placeholder {
        color: black
    }
`;

export default SearchBar;
