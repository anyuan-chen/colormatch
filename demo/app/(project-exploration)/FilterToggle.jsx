import React from "react";
import { Inter } from "@next/font/google";

const inter = Inter();
const FilterToggle = ({ children, toggled, setToggled }) => {
  return (
    <>
      <button
        className={`toggle-switch ${toggled ? "toggled" : "untoggled"} ${
          inter.className
        }`}
        onClick={() => {
          setToggled();
        }}
      >
        {children}
      </button>
      <style jsx>{`
        .toggle-switch {
          border: 1px solid black;
          border-radius: 16px;
          padding: 4px 16px;
          cursor: pointer;
          font-size: 16px;
          height: 30px;
          display: flex;
          align-items: center;
          justify-content: center;
        }
        .untoggled {
          background-color: transparent;
        }
        .toggled {
          color: white;
          background-color: black;
        }
      `}</style>
    </>
  );
};

export default FilterToggle;
