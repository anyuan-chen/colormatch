import React from "react";

function RandomSong(): React.ReactNode {
  const songs = [
    <span key="1" style={{ color: "#4c1d95" }}>
      Lucy In the Sky With Diamonds
    </span>,
    <span key="2" style={{ color: "#14532d" }}>
      Green Light
    </span>,
    <span key="3" style={{ color: "#312e81" }}>
      Happiness is a Butterfly
    </span>,
    <span key="4" style={{ color: "#064e3b" }}>
      This Is What It Feels Like
    </span>,
    <span key="5" style={{ color: "#881337" }}>
      Lose Yourself
    </span>,
    <span key="6" style={{ color: "#1e40af" }}>
      Boulevard of Broken Dreams
    </span>,
    <span key="7" style={{ color: "black" }}>
      Jocelyn Flores
    </span>,
    <span key="8" style={{ color: "#dc2626" }}>
      Revenge
    </span>,
    <span key="8" style={{ color: "#1e3a8a" }}>
      Kill Bill
    </span>,
  ];
  let randomIndex = ~~(Math.random() * songs.length);
  return songs[randomIndex];
}

export { RandomSong };
