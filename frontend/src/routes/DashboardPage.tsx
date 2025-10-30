import { Outlet, type Location } from "react-router-dom";
import { useNavigate, useLocation } from "react-router-dom";
import type { ReactNode } from "react";
import type { IButton } from "../components/lib/Button";
import { EButtonStyles } from "../components/lib/styles";
import Button from "../components/lib/Button";
import { useQuery } from "@tanstack/react-query";
import { ShortUrl } from "../models/urls";

export default function DashboardPage() {
  const navigate = useNavigate();
  const location = useLocation();

  const MenuOptions: IButton[] = [
    {
      id: "menu-linkto-dashboard",
      displayText: "Dashboard",
      route: "/dashboard",
      onClick: () => navigate("/app"),
      btnStyle: isCurrentRoute("/app", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
    {
      id: "menu-linkto-managelinks",
      displayText: "Manage Links",
      route: "Analytics",
      onClick: () => navigate("manage"),
      btnStyle: isCurrentRoute("/app/manage", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
    {
      id: "menu-linkto-analytics",
      displayText: "Analytics",
      route: "Analytics",
      onClick: () => navigate("analytics"),
      btnStyle: isCurrentRoute("/app/analytics", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
    {
      id: "menu-linkto-profile",
      displayText: "Profile",
      route: "profile",
      onClick: () => navigate("profile"),
      btnStyle: isCurrentRoute("/app/profile", location)
        ? EButtonStyles.tertiary
        : EButtonStyles.secondary,
    },
  ];

  const query = useQuery({ queryKey: ["GetUserUrls"], queryFn: getUserUrls });

  const RenderMenuOptions: ReactNode = MenuOptions.map((item) => {
    return <Button key={item.id} {...item} />;
  });

  return (
    <div className="grid gap-rem pad-t-4rem">
      <div className="grid-left pad-t-4rem">
        <div className="verticalstack gap-rem pad-half-rem b ">
          {RenderMenuOptions}
        </div>
      </div>
      <div className="grid-middle">
        <Outlet />
        <div> {JSON.stringify(query.data)}</div>
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

const getUserUrls = async (): Promise<ShortUrl[]> => {
  const baseUrl = import.meta.env.VITE_BACKEND_URL;
  const response = await fetch(baseUrl + "/url/getuserurls", {
    method: "GET",
    credentials: "include",
  });

  const result: ShortUrl[] = await response.json();

  const urls = result.map((item) => new ShortUrl(item));

  return urls;
};
