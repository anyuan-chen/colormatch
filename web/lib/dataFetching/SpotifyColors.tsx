import resolveConfig from "tailwindcss/resolveConfig";
import tailwindConfig from "../../tailwind.config.js";

const fullConfig = resolveConfig(tailwindConfig);
async function GetColors(id) {
  const background = Object.entries(fullConfig.theme.colors)
    .map((color) => {
      return color[1][50];
    })
    .filter((value) => value !== undefined);
  const foreground = Object.entries(fullConfig.theme.colors)
    .map((color) => {
      return color[1][700];
    })
    .filter((value) => value !== undefined);
  const search_type = "album";
  const params = new URLSearchParams({
    search_type: search_type,
    id: id,
  });
  for (const color of background) {
    params.append("background_colors", color);
  }
  for (const color of foreground) {
    params.append("highlight_colors", color);
  }
  const req = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL + "/spotify/colors?" + params,
    {
      credentials: "include",
    }
  );
  const data = await req.json();
  return data;
}
function PrintColors() {
  console.log(fullConfig.theme.colors);
  console.log();
}
export { GetColors, PrintColors };
