import React from "react";
import { Inter } from "@next/font/google";
const inter = Inter();
const Layout = ({ children }) => {
  return (
    <>
      <div className="content-container">
        <div className="content-box">
          <h1 className={`${inter.className} title`}>Colormatch</h1>
          {children}
        </div>
      </div>
      <style jsx>{styles}</style>
    </>
  );
};
const styles = `
    .content-container {
        display: flex;
        align-items: center;
        flex-direction: column;
        padding-top: 32px;
    }
    .content-box{
        width: 1200px;
        display: flex;
        flex-direction: column;
        row-gap: 16px;
    }
    .title {
        font-size: 48px;
        font-weight: 500;
    }
`;
export default Layout;
