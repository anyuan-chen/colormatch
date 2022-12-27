import React from "react";
import { Inter } from "@next/font/google";
const inter = Inter();
const Page = () => {
  return (
    <div className="spotify-login-container">
      <h1 className={`${inter.className} spotify-login-title`}>
        This applet uses the Spotify API, so we&apos;re going to need you to
        login to Spotify.
      </h1>
      <h4 className={`${inter.className} spotify-login-disclaimer`}>
        I don&apos;t store or sell your personal information.
      </h4>
      <button className="spotify-login-button">
        <a
          href={process.env.NEXT_PUBLIC_BACKEND_URL + "/auth/spotify/login"}
          className={`${inter.className}`}
        >
          Login To Spotify
        </a>
      </button>
      <style jsx>{`
        .spotify-login-title {
          font-weight: 500;
          font-size: 48px;
        }
        .spotify-login-disclaimer {
          font-weight: 400;
          font-size: 24px;
          margin-bottom: 28px;
        }
        .spotify-login-container {
          display: flex;
          flex-direction: column;
          align-items: flex-start;
          row-gap: 8px;
          padding-top: 10%;
        }
        .spotify-login-button {
          background-color: #1db954;
          margin: 0;
          border: none;
          font-size: 24px;
          font-weight: 500;
          padding: 8px 16px;
          border-radius: 16px;
        }
      `}</style>
    </div>
  );
};

export default Page;
