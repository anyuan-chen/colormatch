`use client`;
import Link from "next/link";

export default function PillButton({
  color,
  onClick,
  children,
  href,
}: {
  color: string;
  children: React.ReactNode;
  onClick: () => void;
  href: string;
}) {
  const BaseButton = (
    <button
      style={{ backgroundColor: color }}
      onClick={(e) => {
        onClick();
      }}
    >
      {children}
    </button>
  );
  if (href) {
    return <Link href={href}>{BaseButton}</Link>;
  } else {
    return BaseButton;
  }
}
