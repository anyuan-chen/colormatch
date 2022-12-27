import Head from "next/head";
import Image from "next/image";
import Header from "../components/header";
import ProjectCard from "../components/project_card";
import useColor from "../hooks/useColor";

export default function Home() {
  const { color } = useColor();
  return (
    <div style={{ maxWidth: "1200px", paddingLeft: "64px" }}>
      <Head>
        <title>Colormatch</title>
        <meta name="description" content="Stories." />
      </Head>
      <main
        style={{ display: "flex", flexDirection: "column", rowGap: "48px" }}
      >
        <Header color={color}></Header>
        <h1 style={{ fontSize: "120px", fontWeight: 500, margin: 0 }}>
          A PLAYGROUND OF COLOR AND DATA
        </h1>
        <div
          style={{
            display: "grid",
            gridGap: "32px",
          }}
        >
          <ProjectCard
            image=""
            title="Last.fm - A Year In Review"
            text_type="applet"
            link="/lastfm"
            color={color}
          ></ProjectCard>
        </div>
      </main>
    </div>
  );
}
