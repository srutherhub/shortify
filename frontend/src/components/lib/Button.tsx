import type { ButtonStyle } from "./styles";

export interface IButton {
  id: string;
  displayText: string;
  route?: string;
  onClick?: () => void;
  btnStyle?: ButtonStyle;
}

export default function Button(props: IButton) {
  return (
    <button
      className={props.btnStyle}
      aria-label={props.id}
      onClick={props.onClick}
    >
      {props.displayText}
    </button>
  );
}
