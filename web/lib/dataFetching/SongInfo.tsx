async function GetTrackAnalysis(id: string) {
  const res = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL +
      "/spotify/trackanalysis?" +
      new URLSearchParams({
        id: id,
      }),
    {
      credentials: "include",
    }
  );
  return res.json();
}

async function GetRecommendations(
  artist_ids: string[],
  track_ids: string[],
  genres: string[],
  recommendation_count: number
) {
  const searchParams = new URLSearchParams({
    recommendation_count: recommendation_count.toString(),
  });
  for (let s of artist_ids) {
    searchParams.append("artist_ids", s);
  }
  for (let g of genres) {
    searchParams.append("genres", g);
  }
  for (let t of track_ids) {
    searchParams.append("track_ids", t);
  }
  const res = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL +
      "/spotify/recommendations?" +
      searchParams,
    {
      credentials: "include",
    }
  );
  return await res.json();
}

async function GetTrackFeatures(id: string) {
  const res = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL +
      "/spotify/trackfeatures?" +
      new URLSearchParams({
        id: id,
      }),
    {
      credentials: "include",
    }
  );
  const data =  await res.json();
  return data;
}

async function Ping() {
  const res = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL + "/spotify/ping",
    {
      credentials: "include",
    }
  );
  const data = await res.json();
  return data;
}
async function Search(query: string, types: string[]) {
  const params = new URLSearchParams({
    query: query,
  });
  for (const type of types) {
    params.append("type", type);
  }
  const res = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL + "/spotify/search?" + params,
    {
      credentials: "include",
    }
  );
  const data = await res.json();
  return data;
}

export { GetTrackAnalysis, GetRecommendations, GetTrackFeatures, Ping, Search };
