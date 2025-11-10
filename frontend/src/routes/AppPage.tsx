import { Outlet, type Location } from "react-router-dom";
import { useNavigate, useLocation } from "react-router-dom";
import type { ReactNode } from "react";
import type { IButton } from "../components/lib/Button";
import { EButtonStyles } from "../components/lib/styles";
import Button from "../components/lib/Button";

export default function AppPage() {
  const navigate = useNavigate();
  const location = useLocation();

  const MenuOptions: IButton[] = [
    {
      id: "menu-linkto-dashboard",
      displayText: "Dashboard",
      route: "/dashboard",
      onClick: () => navigate("/app"),
      icon: "material-symbols--space-dashboard-outline",
      btnStyle: isCurrentRoute("/app", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
    {
      id: "menu-linkto-managelinks",
      displayText: "Manage Links",
      route: "Analytics",
      onClick: () => navigate("manage"),
      icon: "material-symbols--manage-search-rounded",
      btnStyle: isCurrentRoute("/app/manage", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
    {
      id: "menu-linkto-analytics",
      displayText: "Analytics",
      route: "Analytics",
      onClick: () => navigate("analytics"),
      icon: "material-symbols--insert-chart-outline-rounded",
      btnStyle: isCurrentRoute("/app/analytics", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
    {
      id: "menu-linkto-profile",
      displayText: "Profile",
      route: "profile",
      onClick: () => navigate("profile"),
      icon: "material-symbols--account-circle",
      btnStyle: isCurrentRoute("/app/profile", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
  ];

  const RenderMenuOptions: ReactNode = MenuOptions.map((item) => {
    return <Button key={item.id} {...item} />;
  });

  return (
    <div className="grid gap-rem pad-rem">
      <div className="grid-middle flex center">
        <div className={`gap-rem pad-half-rem b horizontalstack`}>
          {RenderMenuOptions}
        </div>
      </div>
      <div className="grid-middle">
        <Outlet />
      </div>
      <div className="grid-right"></div>
    </div>
  );
}

function isCurrentRoute(route: string, location: Location): boolean {
  if (location.pathname == route) {
    return true;
  } else {
    return false;
  }
}
