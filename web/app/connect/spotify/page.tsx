"use client";
import LinkButton from "../../../components/LinkButton";
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
      <h1 style={{ fontSize: "48px" }}>
        This applet uses the Spotify API, so we’re going to need you to log in
        to Spotify. I don’t store or sell your personal information.
      </h1>
      <LinkButton
        color="#1DB954"
        href="http://localhost:8080/auth/spotify/login"
      >
        <span className="caption">Login to Spotify</span>
      </LinkButton>
    </div>
  );
}
