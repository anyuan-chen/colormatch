interface ButtonProps {
  color: string;
  children: React.ReactNode;
  onClick?: any;
  type?: "submit" | "button" | "reset";
}
interface LinkButtonProps {
  color: string;
  children: React.ReactNode;
  href: string;
  onClick?: any;
  type?: "submit" | "button" | "reset";
}

export type { ButtonProps, LinkButtonProps };
