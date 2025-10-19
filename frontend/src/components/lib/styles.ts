export const EButtonStyles = {
  secondary: "sec-btn",
  primary: "prim-btn",
  tertiary: "ter-btn",
  quaternary: "quad-btn",
};
export type ButtonStyle = (typeof EButtonStyles)[keyof typeof EButtonStyles];
