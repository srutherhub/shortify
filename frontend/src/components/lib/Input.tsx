import { useState, type HTMLInputTypeAttribute } from "react";

export interface IInputProps {
  isRequired: boolean;
  value: string;
  onChange: (value: string) => void;
  type?: HTMLInputTypeAttribute;
  placeholder?: string;
  labelText?: string;
  errorText?: string;
  validationFunc?: (val: string) => boolean;
}

export default function Input(props: IInputProps) {
  const [focus, setFocus] = useState(false);

  const errorText =
    props.validationFunc && focus ? (
      props.validationFunc(props.value) ? (
        <p className="invis font-small marg-vert-small ">{props.errorText}</p>
      ) : (
        <p className="error font-small marg-vert-small">{props.errorText}</p>
      )
    ) : (
      <p className="invis font-small marg-vert-small ">{props.errorText}</p>
    );

  return (
    <div className="verticalstack">
      <div className="horizontalstack">
        {props.isRequired ? (
          <span className="error" style={{ paddingLeft: "4px" }}>
            *
          </span>
        ) : (
          ""
        )}
        <label className="marg-vert-small">{props.labelText}</label>
      </div>
      <div className="horizontalstack width100">
        <input
          className="inp-primary"
          value={props.value}
          type={props.type}
          onChange={(e) => props.onChange(e.target.value)}
          onBlur={() => setFocus(true)}
          placeholder={props.placeholder}
        />
      </div>
      {errorText}
    </div>
  );
}
