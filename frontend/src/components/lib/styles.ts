export const EButtonStyles = {
  secondary: "sec-btn",
  primary: "prim-btn",
  tertiary: "ter-btn",
};
export type ButtonStyle = (typeof EButtonStyles)[keyof typeof EButtonStyles];
