import React from "react";
import Header from "../../../components/header";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div style={{ maxWidth: "1200px", paddingLeft: "64px" }}>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          rowGap: "48px",
          minHeight: "100vh",
        }}
      >
        <Header></Header>
        {children}
      </div>
    </div>
  );
}
