import { useState } from "react";
import type { IButton } from "../components/lib/Button";
import Button from "../components/lib/Button";
import Input, { type IInputProps } from "../components/lib/Input";
import { EButtonStyles } from "../components/lib/styles";
import { Utils } from "../models/utils";
import { useQuery } from "@tanstack/react-query";
import { postShortify } from "../api/PostShortify";

export default function ManageLinksPage() {
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

  return (
    <div>
      <h2>Create New</h2>
      <div className="horizontalstack gap-half-rem">
        <Button {...newShortUrlButton} />
        <Button {...newRouteButton} />
      </div>
      <CreateShortLinkForm />
    </div>
  );
}

function CreateShortLinkForm() {
  const [formError, setFormError] = useState("");
  const [url, setUrl] = useState("");
  const [route, setRoute] = useState("");
  const [source, setSource] = useState("");
  const [medium, setMedium] = useState("");
  const [campaign, setCampaign] = useState("");
  const [term, setTerm] = useState("");
  const [content, setContent] = useState("");

  const { refetch } = useQuery({
    queryKey: ["PostShortify"],
    queryFn: () =>
      postShortify({
        url: url,
        utm_source: source,
        utm_medium: medium,
        utm_campaign: campaign,
        utm_term: term,
        utm_content: content,
      }),
    enabled: false,
  });

  const inputs: IInputProps[] = [
    {
      type: "text",
      onChange: setUrl,
      value: url,
      isRequired: true,
      labelText: "Url",
      placeholder: "https://mywebsite.com",
      validationFunc: Utils.Validators.isHttps,
      errorText: "Url must start with https://",
    },
    {
      type: "text",
      onChange: setRoute,
      value: route,
      isRequired: true,
      labelText: "Alias",
      placeholder: "MYBRAND",
      validationFunc: Utils.Validators.isValidRoute,
      errorText:
        "Alias must be 10 character or less and must contain url safe characters",
    },
    {
      type: "text",
      onChange: setSource,
      value: source,
      isRequired: false,
      labelText: "Source",
      placeholder: "Google, Facebook, Youtube, etc.",
      validationFunc: Utils.Validators.isChar64,
      errorText: "Must be 64 characters or less",
    },
    {
      type: "text",
      onChange: setMedium,
      value: medium,
      isRequired: false,
      labelText: "Medium",
      placeholder: "Email, SMS, Social, etc.",
      validationFunc: Utils.Validators.isChar64,
      errorText: "Must be 64 characters or less",
    },
    {
      type: "text",
      onChange: setCampaign,
      value: campaign,
      isRequired: false,
      labelText: "Campaign",
      placeholder: "2025_01_03_Daily_Newsletter",
      validationFunc: Utils.Validators.isChar64,
      errorText: "Must be 64 characters or less",
    },
    {
      type: "text",
      onChange: setTerm,
      value: term,
      isRequired: false,
      labelText: "Term",
      validationFunc: Utils.Validators.isChar64,
      errorText: "Must be 64 characters or less",
    },
    {
      type: "text",
      onChange: setContent,
      value: content,
      isRequired: false,
      labelText: "Content",
      placeholder: "",
      validationFunc: Utils.Validators.isChar64,
      errorText: "Must be 64 characters or less",
    },
  ];

  const submitButton: IButton = {
    id: "submit-create-new-short-url",
    displayText: "Submit",
    btnStyle: EButtonStyles.primary,
    onClick: () => handleSubmitButton(),
  };

  const renderInputs = inputs.map((item, ind) => {
    return <Input {...item} key={ind} />;
  });

  const handleSubmitButton = async () => {
    let err: string = "";
    for (const i in inputs) {
      const input = inputs[i];
      if (input.isRequired && input.value == "") {
        err = "Missing required fields";
      } else if (input.validationFunc) {
        if (!input.validationFunc(input.value)) {
          err = "Resolve all errors";
        }
      } else {
        err = "";
      }
    }
    if (err) {
      setFormError(err);
      return;
    }
    setFormError("");
    const result = await refetch();
    console.log(result.data);
  };

  return (
    <form onSubmit={(e) => e.preventDefault()}>
      {renderInputs}
      <span>{formError}</span>
      <Button {...submitButton} />
    </form>
  );
}
