/* eslint-disable react-hooks/exhaustive-deps */
"use client";
import { useEffect, useState } from "react";
import {
  Ping,
  GetTrackAnalysis,
  GetTrackFeatures,
  GetRecommendations,
} from "../../../lib/dataFetching/SongInfo";
import { GetColors } from "../../../lib/dataFetching/SpotifyColors";
import { useSearchParams } from "next/navigation";

export default function SongPage() {
  const [ping, setPing] = useState(false);
  const [analysis, setAnalysis] = useState(null);
  const [features, setFeatures] = useState(null);
  const [recommendations, setRecommendations] = useState(null);
  const [loading, setLoading] = useState(true);
  const searchParams = useSearchParams();
  useEffect(() => {
    async function FetchData() {
      const id = searchParams.get("id");
      const album_id = searchParams.get("album_id")
      const [fetched_analysis, fetched_features, fetched_recommendations] =
        await Promise.all([
          GetTrackAnalysis(id),
          GetTrackFeatures(id),
          GetRecommendations([], [id], [], 5),
        ]);
      setAnalysis(fetched_analysis);
      setFeatures(fetched_features);
      setRecommendations(fetched_recommendations)
      console.log(GetColors(album_id))
      setLoading(false);
    }
    FetchData();
    console.log(analysis, features, recommendations);
  }, [searchParams]);
  if (loading) {
    <div>Loading...</div>;
  }

  return <div>{ping}</div>;
}

// async function getData() {
//   const nextCookies = cookies();
//   const token = nextCookies.get("spotify_session_id");
//   const req = await fetch("http://localhost:8080/spotify/ping");
//   const data = await req.json();
//   return data;
// }
