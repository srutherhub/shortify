import type { ButtonStyle } from "./styles";

export interface IButton {
  id: string;
  displayText: string;
  route?: string;
  onClick?: () => void;
  btnStyle?: ButtonStyle;
  icon?: string;
}

export default function Button(props: IButton) {
  return (
    <button
      className={props.btnStyle}
      aria-label={props.id}
      onClick={props.onClick}
    >
      <span className={props.icon}></span> {props.displayText}
    </button>
  );
}
