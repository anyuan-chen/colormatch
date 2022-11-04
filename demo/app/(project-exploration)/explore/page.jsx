import React from "react";
import ProjectCard from "./ProjectCard";

const Explore = () => {
  const projects = [
    {
      title: "Sample Test 1",
      subtitle: "this is a correct subtitle",
      color: "#EF8B7C",
      rounding: [true, false, false, false], //tl tr br bl
    },
    {
      title: "Sample Test 2",
      subtitle: "this is a correct subtitle",
      color: "#EF8B7C",
      rounding: [false, true, false, false], //tl tr br bl
    },
    {
      title: "Sample Test 3",
      subtitle: "this is a correct subtitle",
      color: "#EF8B7C",
      rounding: [false, false, false, true], //tl tr br bl
    },
    {
      title: "Sample Test 4",
      subtitle: "this is a correct subtfdsfitle",
      color: "#EF8B7C",
      rounding: [false, false, true, false], //tl tr br bl
    },
  ];
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
