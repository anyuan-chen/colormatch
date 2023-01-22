import React from "react";
import Link from "next/link";
const Header = () => {
  return (
    <div
      style={{
        display: "flex",
        columnGap: "24px",
        alignItems: "flex-end",
        padding: "32px 0px",
      }}
    >
      <Link href="/">
        <h3
          style={{
            fontWeight: 400,
            fontSize: "32px",
            fontFamily: "Value Serif Pro",
            margin: 0,
          }}
        >
          <span>Color</span>
          <span style={{ fontWeight: "bold" }}>match</span>
        </h3>
      </Link>
      <h4 style={{ fontWeight: 400, fontSize: "16px", margin: "4px" }}>
        stories in color
      </h4>
    </div>
  );
};

export default Header;
