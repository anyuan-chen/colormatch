export default function BaseButton({ color, children }) {
  return (
    <button
      style={{
        backgroundColor: color,
        padding: "12px 24px",
        borderRadius: "12px",
        border: "none",
        cursor: "pointer",
      }}
    >
      {children}
    </button>
  );
}
