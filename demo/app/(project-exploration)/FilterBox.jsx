"use client";
import { Inter } from "@next/font/google";
import React from "react";
import { useState, useEffect } from "react";
import FilterToggle from "./FilterToggle";
const inter = Inter();
const FilterBox = () => {
  const [filterInfo, setFilterInfo] = useState([]);
  const [filters, setFilters] = useState([]);
  useEffect(() => {
    setFilterInfo([
      {
        name: "Color",
      },
      {
        name: "Music",
      },
    ]);
  }, []);

  useEffect(() => {
    setFilters(Array(filterInfo.length).fill(false));
  }, [filterInfo]);

  return (
    <>
      <div className="toggle-box">
        <h2 className={`label ${inter.className}`}>Filter:</h2>
        {filters.length != 0 &&
          filterInfo.map((filter, index) => {
            return (
              <FilterToggle
                key={filter.name + index}
                toggled={filters[index]}
                setToggled={() => {
                  const newData = [
                    ...filters.slice(0, index),
                    !filters[index],
                    ...filters.slice(index + 1),
                  ];
                  setFilters(newData);
                }}
              >
                {filter.name}
              </FilterToggle>
            );
          })}
      </div>
      <style jsx>{`
        .label {
          font-size: 24px;
          font-weight: 500;
          padding-right: 8px;
        }
        .toggle-box {
          display: flex;
          flex-direction: row;
          column-gap: 4px;
          align-items: center;
        }
      `}</style>
    </>
  );
};

export default FilterBox;
