import { useEffect, useState } from "react";

export class Utils {
  public static Validators = {
    isHttps: (val: string): boolean => {
      if (val.slice(0, 8) == "https://") {
        return true;
      } else {
        return false;
      }
    },
  };

  public static isScreenMobile(breakpoint = 768): boolean {
    if (window.innerWidth <= breakpoint) {
      return true;
    } else {
      return false;
    }
  }
}

export function useIsMobileView(breakpoint = 768) {
  const [isMobile, setIsMobile] = useState(Utils.isScreenMobile(breakpoint));

  useEffect(() => {
    const handleResize = () => {
      setIsMobile(window.innerWidth <= breakpoint);
    };

    window.addEventListener("resize", handleResize);

    // Cleanup listener on unmount
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, [breakpoint]);

  return isMobile;
}
