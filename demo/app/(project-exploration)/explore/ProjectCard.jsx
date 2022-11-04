import React from "react";
import { Averia_Serif_Libre, Inter } from "@next/font/google";
const asl = Averia_Serif_Libre({
  weight: "400",
});
const inter = Inter();
const ProjectCard = ({
  children,
  title,
  subtitle,
  color,
  topLeft,
  topRight,
  bottomLeft,
  bottomRight,
}) => {
  console.log(topLeft, topRight, bottomLeft, bottomRight);
  return (
    <>
      <div
        className={`project-card`}
        style={{
          borderRadius: `${topLeft ? "16px" : "0px"} ${topRight ? "16px" : "0px"} ${
            bottomRight ? "16px" : "0px"
          } ${bottomLeft ? "16px" : "0px"}`,
        }}
      >
        {children}
        <h2 className={`${asl.className} title`}>{title}</h2>
        <h3 className={`${inter.className} subtitle`}>{subtitle}</h3>
      </div>
      <style jsx>
        {`
          .project-card {
            height: 500px;
            width: 50%;
            background-color: ${color};
            border-bottom: 1px solid black;
            border-right: 1px solid black;
          }
          .title {
            font-size: 24px;
            margin-top: auto;
          }
          .subtitle {
            font-size: 16px;
            font-weight: 400;
          }
        `}
      </style>
    </>
  );
};

export default ProjectCard;
