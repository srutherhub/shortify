export interface IQuickStats {
  total_url_count: number;
  total_click_count: number;
  avg_clicks_per_url: number;
}

export const getQuickStats = async () => {
  const baseUrl = import.meta.env.VITE_BACKEND_URL;
  const response = await fetch(baseUrl + "/url/getquickstats", {
    method: "GET",
    credentials: "include",
  });

  const result: IQuickStats = await response.json();

  return result;
};
