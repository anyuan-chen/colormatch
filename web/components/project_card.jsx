/* eslint-disable @next/next/no-img-element */
import Link from "next/link";
import React from "react";

const ProjectCard = ({ image, title, text_type, link = "", color }) => {
  return (
    <div
      style={{
        height: "100%",
        width: "100%",
      }}
    >
      <div style={{ height: "auto", width: "auto" }}>
        <Link href={link}>
          <img
            src={image}
            alt="project image"
            style={{ border: color + " 1px solid" }}
          ></img>
        </Link>
      </div>
      <Link href={link}>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            justifyContent: "space-between",
          }}
        >
          <h2 style={{ fontSize: "32px", fontWeight: 500 }}>{title}</h2>
          <h4
            style={{ fontWeight: "medium", fontSize: "12px", fontWeight: 400 }}
          >
            {text_type}
          </h4>
        </div>
      </Link>
    </div>
  );
};

export default ProjectCard;
