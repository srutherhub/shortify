import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Hanko } from "@teamhanko/hanko-elements";
import type { IButton } from "../lib/Button";
import Button from "../lib/Button";
import { EButtonStyles } from "../lib/styles";
import useIsLoggedIn from "./useIsLoggedIn";

const hankoApi = import.meta.env.VITE_HANKO_API_URL;

function HankoLogout() {
  const navigate = useNavigate();
  const [hanko, setHanko] = useState<Hanko>();
  const isLoggedIn = useIsLoggedIn();

  useEffect(() => {
    import("@teamhanko/hanko-elements").then(({ Hanko }) =>
      setHanko(new Hanko(hankoApi ?? ""))
    );
  }, []);

  const logout = async () => {
    try {
      await hanko?.logout();

      navigate("/"); //Path to naviage to once the user logs out.
    } catch (error) {
      console.error("Error during logout:", error);
    }
  };

  const LogoutButton: IButton = {
    id: "navigation-linkto-logout",
    displayText: "Logout",
    onClick: () => logout(),
    btnStyle: EButtonStyles.quaternary,
  };

  return isLoggedIn ? <Button {...LogoutButton} /> : null;
}

export default HankoLogout;
