import { useState } from "react";
import type { IButton } from "../components/lib/Button";
import Button from "../components/lib/Button";
import Input, { type IInputProps } from "../components/lib/Input";
import { EButtonStyles } from "../components/lib/styles";
import { Utils } from "../models/utils";

export default function ManageLinksPage() {
  const [url, setUrl] = useState("");

  const newShortUrlButton: IButton = {
    id: "create-new-link-button",
    displayText: "Short Url",
    onClick: () => console.log("hi"),
    btnStyle: EButtonStyles.primary,
    icon: "material-symbols--add-link",
  };

  const newRouteButton: IButton = {
    id: "create-new-route-button",
    displayText: "Alias",
    onClick: () => console.log("hi"),
    btnStyle: EButtonStyles.tertiary,
    icon: "material-symbols--map-rounded",
  };
  const urlInput: IInputProps = {
    type: "text",
    onChange: setUrl,
    value: url,
    isRequired: true,
    labelText: "Url",
    placeholder: "https://mywebsite.com",
    validationFunc: Utils.Validators.isHttps,
    errorText: "Url must start with https://",
  };

  return (
    <div>
      <h2>Create New</h2>
      <div className="horizontalstack gap-half-rem">
        <Button {...newShortUrlButton} />
        <Button {...newRouteButton} />
      </div>
      <form onSubmit={(e) => e.preventDefault()}>
        <Input {...urlInput} />
      </form>
    </div>
  );
}
