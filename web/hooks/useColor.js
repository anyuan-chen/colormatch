import React from "react";

const useColor = () => {
  const [color, setColor] = React.useState("red");
  return { color };
};

export default useColor;
