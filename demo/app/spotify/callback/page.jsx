import React from "react";
import { Inter } from "@next/font/google";
const inter = Inter();
const Page = () => {
  return (
    <div className="spotify-callback-container">
      <h1 className={`${inter.className} spotify-callback-title`}>
        Thanks for logging in! You can now use any applet that requires Spotify
        login.
      </h1>

      <button className="spotify-callback-button">
        <a href={"/explore"} className={`${inter.className}`}>
          Back to Explore
        </a>
      </button>
      <style jsx>{`
        .spotify-callback-title {
          font-weight: 500;
          font-size: 48px;
        }
        .spotify-callback-container {
          display: flex;
          flex-direction: column;
          align-items: flex-start;
          row-gap: 64px;
          padding-top: 10%;
        }
        .spotify-callback-button {
          background-color: #739af0;
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
