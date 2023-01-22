import { ButtonProps } from "../../types/button";
export default function BaseButton({
  color,
  children,
  onClick = () => {},
  type = "button",
}: ButtonProps) {
  return (
    <button
      style={{
        backgroundColor: color,
        padding: "12px 24px",
        borderRadius: "12px",
        border: "none",
        cursor: "pointer",
      }}
      onClick={() => onClick}
      type={type}
    >
      {children}
    </button>
  );
}
