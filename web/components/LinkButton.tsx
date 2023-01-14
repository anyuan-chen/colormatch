`use client`;
import Link from "next/link";
import BaseButton from "./button/button";

export default function LinkButton({
  color,
  children,
  href,
}: {
  color: string;
  children: React.ReactNode;
  onClick: () => void;
  href: string;
}) {
  return (
    <Link href={href} style={{ cursor: "pointer" }}>
      <BaseButton color={color}>{children}</BaseButton>
    </Link>
  );
}
