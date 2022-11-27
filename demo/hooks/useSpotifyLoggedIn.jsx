import React from "react";
import useSWR from "swr";
const SpotifyLoggedIn = () => {
  const fetcher = (...args) => fetch(...args).then((res) => res.json());
};

export default SpotifyLoggedIn;
