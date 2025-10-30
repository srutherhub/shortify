import { useEffect, useState } from "react";
import type { ReactNode } from "react";
import { Navigate, useLocation } from "react-router-dom";

const backendUrl = import.meta.env.VITE_BACKEND_URL;

export default function HankoPrivateRoute({
  children,
}: {
  children: ReactNode;
}) {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
  const location = useLocation();

  useEffect(() => {
    fetch(backendUrl + "/auth/validate", {
      credentials: "include",
    })
      .then((res) => {
        setIsAuthenticated(res.ok);
      })
      .catch(() => {
        setIsAuthenticated(false);
      });
  }, []);

  if (isAuthenticated === null) {
    return null; // Or a loading spinner
  }

  if (isAuthenticated) {
    return <>{children}</>;
  }

  //Url to naviage user to if they arent authenticated
  return <Navigate to="/login" replace state={{ from: location }} />;
}
