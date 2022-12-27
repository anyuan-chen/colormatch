import React from "react";

const Header = ({ color }) => {
  return (
    <div
      style={{
        display: "flex",
        columnGap: "24px",
        alignItems: "center",
      }}
    >
      <h3 style={{ fontWeight: 500, fontSize: "32px" }}>
        <span>Color</span>
        <span style={{ color: color, fontStyle: "italic" }}>match</span>
      </h3>
      <h4 style={{ fontWeight: 400, fontSize: "16px" }}>stories in color</h4>
    </div>
  );
};

export default Header;
