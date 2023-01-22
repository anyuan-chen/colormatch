"use client";
import LinkButton from "../../../../components/LinkButton";
export default function SpotifyLoginPage() {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        rowGap: "64px",
        alignItems: "start",
        justifyContent: "space-between",
        height: "100%",
      }}
    >
      <h1 style={{ fontSize: "48px" }}>Thanks for logging in!</h1>
      <LinkButton color="#1e3a8a" href="/">
        <span className="caption">Back to Home</span>
      </LinkButton>
    </div>
  );
}
