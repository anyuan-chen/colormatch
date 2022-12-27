"use client";
import React, { useEffect, useState } from "react";
import useSWR from "swr";
const fetcher = (...args) => fetch(...args).then((res) => res.json());

const Page = () => {
  const [loggedInToSpotify, setLoggedInToSpotify] = useState(false);
  useEffect(() => {
    const GetSpotifyData = async () => {
      const res = await fetch(
        process.env.NEXT_PUBLIC_BACKEND_URL + "/spotify/ping",
        {
          credentials: "include",
        }
      );
      const data = await res.json();
      console.log(data);
      setLoggedInToSpotify(data.valid);
    };
    GetSpotifyData();
  }, []);
  return <div>{JSON.stringify(loggedInToSpotify)}</div>;
};

export default Page;
