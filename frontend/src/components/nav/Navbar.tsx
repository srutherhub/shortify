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
      id: "navigation-linkto-signup",
      displayText: "Signup",
      onClick: () => navigate("/login"),
      btnStyle: EButtonStyles.primary,
    },
    {
      id: "navigation-linkto-dashboard",
      displayText: "Dashboard",
      route: "/dashboard",
      onClick: () => navigate("/dashboard"),
      btnStyle: EButtonStyles.tertiary,
    },
    {
      id: "navigation-linkto-profile",
      displayText: "Profile",
      route: "/dashboard/profile",
      onClick: () => navigate("/dashboard/profile"),
      btnStyle: EButtonStyles.secondary,
    },
  ];

  const hiddenNavOptionsWhenLoggedIn = [
    "navigation-pricing",
    "navigation-linkto-signup",
  ];

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
