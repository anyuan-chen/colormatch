import React from "react";
import ProjectCard from "./ProjectCard";

const Explore = () => {
  const projects = [
    {
      title: "Artist and Album Palettes",
      subtitle: "Colormatching Artists and Albums to Your Brand",
      color: "#EF8B7C",
      rounding: [true, false, false, false], //tl tr br bl
      link: "projects/brand-colors",
    },
    {
      title: "Sample Test 2",
      subtitle: "this is a correct subtitle",
      color: "#C3B9FA",
      rounding: [false, true, false, false], //tl tr br bl
    },
    {
      title: "Sample Test 3",
      subtitle: "this is a correct subtitle",
      color: "#739AF0",
      rounding: [false, false, false, true], //tl tr br bl
    },
    {
      title: "Sample Test 4",
      subtitle: "this is a correct subtfdsfitle",
      color: "#F6C944",
      rounding: [false, false, true, false], //tl tr br bl
    },
  ];
  //get url and query the projects based on that
  return (
    <>
      <div className="card-container">
        {projects.map((project) => {
          return (
            <ProjectCard
              title={project.title}
              subtitle={project.subtitle}
              color={project.color}
              topLeft={project.rounding[0]}
              topRight={project.rounding[1]}
              bottomLeft={project.rounding[3]}
              bottomRight={project.rounding[2]}
              key={project.title + project.subtitle + project.color}
              link={project.link}
            ></ProjectCard>
          );
        })}
      </div>
      <style jsx>
        {`
          .card-container {
            display: flex;
            flex-wrap: wrap;
            flex-direction: row;
            border-top: 1px solid black;
            border-left: 1px solid black;
            border-radius: 16px;
          }
        `}
      </style>
    </>
  );
};

export default Explore;
