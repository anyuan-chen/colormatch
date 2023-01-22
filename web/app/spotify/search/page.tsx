"use client";
import { useState, useEffect } from "react";
import BaseButton from "../../../components/PillButton";
import { RandomSong } from "../../../lib/RandomSong";
import { Search } from "../../../lib/dataFetching/SongInfo";
import { useRouter } from "next/navigation";

export default function Page() {
  const router = useRouter();
  const [songTitle, setSongTitle] = useState<string>("");
  const [currentSong, setCurrentSong] = useState<React.ReactNode>("");
  const [error, setError] = useState("");
  useEffect(() => {
    const songName = setInterval(() => {
      let song = RandomSong();
      setCurrentSong(song);
    }, 1000);
    return () => {
      clearInterval(songName);
    };
  }, []);

  const FindAndRedirect = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const types = ["track"];
    const res = await Search(songTitle, types);
    if (res.error === true) {
      setError("Try logging into Spotify again.");
    }
    if (res.tracks.items.length === 0) {
      setError("This song has no queries - try another");
    }
    const id = res.tracks.items[0].id;
    const album_id = res.tracks.items[0].album.id;
    const params = new URLSearchParams({
      id: id,
      album_id: album_id,
    });
    router.push(`/spotify/song/?` + params);
    console.log(res);
  };
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "start",
        rowGap: "32px",
        justifyContent: "center",
        height: "100%",
      }}
    >
      <h1 style={{ fontSize: "48px", fontFamily: "Value Serif Pro" }}>
        Search For {currentSong}
      </h1>
      <form
        style={{
          paddingBottom: "48px",
          display: "flex",
          flexDirection: "column",
          width: "100%",
        }}
        onSubmit={(e) => FindAndRedirect(e)}
      >
        <input
          type="text"
          className="text-input-large"
          value={songTitle}
          placeholder="Enter a song title here..."
          style={{
            fontSize: "24px",
            fontFamily: "Value Sans Pro",
            padding: "16px 32px",
            width: "100%",
          }}
          onChange={(e) => {
            setSongTitle(e.target.value);
          }}
        ></input>
        <BaseButton color="#1DB954" type="submit">
          <span className="caption">Find my song</span>
        </BaseButton>
      </form>
      {error != "" && <span style={{ color: "red" }}>{error}</span>}
    </div>
  );
}
