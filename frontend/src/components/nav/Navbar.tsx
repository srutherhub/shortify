import type { ReactNode } from "react";
import Button from "../lib/Button";
import { EButtonStyles } from "../lib/styles";
import type { IButton } from "../lib/Button";

const NavOptions: IButton[] = [
  {
    id: "navigation-linkto-products",
    displayText: "Products",
    route: "/products",
    btnStyle: EButtonStyles.secondary,
  },
  {
    id: "navigation-linkto-docs",
    displayText: "Docs",
    route: "/docs",
    btnStyle: EButtonStyles.secondary,
  },
  {
    id: "navigation-pricing",
    displayText: "Pricing",
    btnStyle: EButtonStyles.secondary,
  },
  {
    id: "navigation-linkto-changelog",
    displayText: "Changelog",
    route: "/changelog",
    btnStyle: EButtonStyles.secondary,
  },
  {
    id: "navigation-linkto-signup",
    displayText: "Signup",
    btnStyle: EButtonStyles.primary,
  },
  {
    id: "navigation-linkto-dashboard",
    displayText: "Dashboard",
    route: "/dashboard",
    btnStyle: EButtonStyles.tertiary,
  },
];

export default function Navbar() {
  const RenderNavOptions: ReactNode = NavOptions.map((item) => {
    return <Button key={item.id} {...item} />;
  });

  return (
    <nav className="layout-flex gap-rem w-100 b-b pad-half-rem">
      {RenderNavOptions}
    </nav>
  );
}
