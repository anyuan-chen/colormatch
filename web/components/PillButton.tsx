`use client`;
import Link from "next/link";
import { ButtonProps } from "../types/button";
import BaseButton from "./button/button";
export default function PillButton({
  color,
  onClick = () => {},
  children,
  type,
}: ButtonProps) {
  return (
    <BaseButton color={color} onClick={onClick} type={type}>
      {children}
    </BaseButton>
  );
}
