import React from "react";

const Header = () => {
  return (
    <div
      style={{
        display: "flex",
        columnGap: "24px",
        alignItems: "center",
      }}
    >
      <h3
        style={{
          fontWeight: 400,
          fontSize: "32px",
          fontFamily: "Value Serif Pro",
        }}
      >
        <span>Color</span>
        <span style={{ fontWeight: "bold" }}>match</span>
      </h3>
      <h4 style={{ fontWeight: 400, fontSize: "16px" }}>stories in color</h4>
    </div>
  );
};

export default Header;
