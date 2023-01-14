import React from "react";
import Header from "../../../components/header";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div
      style={{ maxWidth: "1200px", paddingLeft: "64px", minHeight: "100vh" }}
    >
      <main
        style={{ display: "flex", flexDirection: "column", rowGap: "48px" }}
      >
        <Header></Header>
        {children}
      </main>
    </div>
  );
}
