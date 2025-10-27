import type { ReactNode } from "react";
import Button from "../lib/Button";
import { EButtonStyles } from "../lib/styles";
import type { IButton } from "../lib/Button";
import { useNavigate } from "react-router-dom";
import useIsLoggedIn from "../auth/useIsLoggedIn";
import HankoLogout from "../auth/HankoLogout";

export default function Navbar() {
  const navigate = useNavigate();
  const isLoggedIn = useIsLoggedIn();
  const NavOptions: IButton[] = [
    {
      id: "navigation-linkto-products",
      displayText: "Products",
      route: "/products",
      onClick: () => navigate("/products"),
      btnStyle: EButtonStyles.secondary,
    },
    {
      id: "navigation-linkto-docs",
      displayText: "Docs",
      route: "/docs",
      onClick: () => navigate("/docs"),
      btnStyle: EButtonStyles.secondary,
    },
    {
      id: "navigation-pricing",
      displayText: "Pricing",
      onClick: () => navigate("/pricing"),
      btnStyle: EButtonStyles.secondary,
    },
    {
      id: "navigation-linkto-changelog",
      displayText: "Changelog",
      route: "/changelog",
      onClick: () => navigate("/changelog"),
      btnStyle: EButtonStyles.secondary,
    },
    {
      id: "navigation-linkto-app",
      displayText: "Dashboard",
      route: "/app",
      onClick: () => navigate("/app"),
      btnStyle: EButtonStyles.tertiary,
    },
    {
      id: "navigation-linkto-signin",
      displayText: "Sign in",
      onClick: () => navigate("/login"),
      btnStyle: EButtonStyles.primary,
    },
  ];

  const hiddenNavOptionsWhenLoggedIn = ["navigation-linkto-signin"];

  const showNavOptionWhenLoggedIn = ["navigation-linkto-profile"];

  const RenderNavOptions: ReactNode = NavOptions.map((item) => {
    if (isLoggedIn && hiddenNavOptionsWhenLoggedIn.includes(item.id)) {
      return null;
    } else {
      if (!isLoggedIn && showNavOptionWhenLoggedIn.includes(item.id)) {
        return null;
      } else {
        return <Button key={item.id} {...item} />;
      }
    }
  });

  return (
    <nav className="horizontalstack gap-rem width100 center b-b pad-half-rem">
      {RenderNavOptions}
      <HankoLogout />
    </nav>
  );
}
