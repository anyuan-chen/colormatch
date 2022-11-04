import React from "react";
import FilterBox from "./FilterBox";
import SearchBar from "./SearchBar";
import { Inter } from "@next/font/google";
import { Averia_Sans_Libre } from "@next/font/google";
const inter = Inter();
const Layout = ({ children }) => {
  return (
    <>
      <div className="container">
        <div className="content-box">
          <h1 className={`${inter.className} title`}>Colormatch</h1>
          <div className="filter-search">
            <FilterBox></FilterBox>
            <SearchBar></SearchBar>
          </div>
          {children}
        </div>
      </div>
      <style jsx>{styles}</style>
    </>
  );
};
const styles = `
    .container {
        display: flex;
        align-items: center;
        flex-direction: column;
        padding-top: 16px;
        row-gap: 16px;
    }
    .content-box{
        width: 1200px;
        display: flex;
        flex-direction: column;
        row-gap: 8px;
    }
    .title {
        font-size: 48px;
        font-weight: 500;
    }
    .filter-search {
        display: flex;
        justify-content: space-between;
        
    }
`;
export default Layout;
