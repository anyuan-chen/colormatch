import Header from "../components/header";
import ProjectCard from "../components/project_card";

export default function Home() {
  return (
    <div style={{ maxWidth: "1200px", paddingLeft: "64px" }}>
      <main
        style={{ display: "flex", flexDirection: "column", rowGap: "48px" }}
      >
        <Header></Header>
        <h1 style={{ fontSize: "120px", fontWeight: 500, margin: 0 }}>
          A PLAYGROUND OF
        </h1>
        <h1 style={{ fontSize: "120px", fontWeight: 500, margin: 0, }}>
          COLOR AND DATA
        </h1>
        <div
          style={{
            display: "grid",
            gridGap: "32px",
          }}
        >
          <ProjectCard
            image=""
            title="Imagining Spotify Song Pages"
            text_type="applet"
            link="/spotify/search"
          ></ProjectCard>
        </div>
      </main>
    </div>
  );
}
