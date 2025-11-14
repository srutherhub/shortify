export interface IShortifyBody {
  url: string;
  RouteName?: string;
  ExpiresAt?: number;
  utm_source: string;
  utm_medium: string;
  utm_campaign: string;
  utm_term: string;
  utm_content: string;
}

export interface IShortifyResponse {
  shorturl: string;
}

export const postShortify = async (
  data: IShortifyBody
): Promise<IShortifyResponse> => {
  const baseUrl = import.meta.env.VITE_BACKEND_URL;
  const response = await fetch(baseUrl + "/url/shortify", {
    method: "POST",
    credentials: "include",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  const result: IShortifyResponse = await response.json();

  return result;
};
