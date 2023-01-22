`use client`;
import Link from "next/link";
import { LinkButtonProps } from "../types/button";
import BaseButton from "./button/button";

export default function LinkButton({
  color,
  children,
  href,
}: LinkButtonProps) {
  return (
    <Link href={href} style={{ cursor: "pointer" }}>
      <BaseButton color={color}>{children}</BaseButton>
    </Link>
  );
}
