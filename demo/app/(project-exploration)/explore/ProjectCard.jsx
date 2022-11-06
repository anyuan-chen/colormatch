import React from "react";
import { Averia_Serif_Libre, Inter } from "@next/font/google";
import Link from "next/link";
const asl = Averia_Serif_Libre({
  weight: "400",
});
const inter = Inter();
const ProjectCard = ({
  link,
  children,
  title,
  subtitle,
  color,
  topLeft,
  topRight,
  bottomLeft,
  bottomRight,
}) => {
  return (
    <>
      <div
        className={`project-card`}
        style={{
          borderRadius: `${topLeft ? "16px" : "0px"} ${
            topRight ? "16px" : "0px"
          } ${bottomRight ? "16px" : "0px"} ${bottomLeft ? "16px" : "0px"}`,
          backgroundColor: color,
        }}
      >
        <Link href={`/${link}`}>
          <div className="padding-project-card">
            {children}
            <h2 className={`${asl.className} project-card-title`}>{title}</h2>
            <h3 className={`${inter.className} project-card-subtitle`}>
              {subtitle}
            </h3>
          </div>
        </Link>
      </div>
      <style jsx>
        {`
          .padding-project-card {
            padding: 16px;
          }
          .project-card {
            height: 500px;
            width: 50%;
            border-bottom: 1px solid black;
            border-right: 1px solid black;
          }
          .project-card-title {
            font-size: 32px;
          }
          .project-card-subtitle {
            font-size: 16px;
            font-weight: 400;
          }
        `}
      </style>
    </>
  );
};

export default ProjectCard;
