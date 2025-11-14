import { useQuery } from "@tanstack/react-query";
import { useIsMobileView } from "../models/utils";
import { getQuickStats } from "../api/GetQuickStats";

export default function DashboardPage() {
  const { data, isLoading } = useQuery({
    queryKey: ["GetQuickStats"],
    queryFn: getQuickStats,
  });
  const isMobile = useIsMobileView();

  return (
    <div>
      <h2>Activity</h2>
      <h2>Engagement</h2>
      <div
        className={`gap-rem ${isMobile ? "verticalstack" : "horizontalstack"}`}
      >
        <BoxWithLabel
          title={"Number of Active Urls"}
          desc={data?.total_url_count}
          isLoading={isLoading}
        />
        <BoxWithLabel
          title={"Total Number of Clicks"}
          desc={data?.total_click_count}
          isLoading={isLoading}
        />
        <BoxWithLabel
          title="Average Number of Clicks per Url"
          desc={data?.avg_clicks_per_url}
          isLoading={isLoading}
        />
      </div>
    </div>
  );
}

interface IBoxwithLabelProps {
  title: string;
  desc: string | number | undefined;
  isLoading: boolean;
}

function BoxWithLabel(props: IBoxwithLabelProps) {
  if (props.isLoading) {
    return (
      <div className="b pad-rem">
        <p className="invis">{props.title}</p>
        <p className="invis font-large">Loading</p>
      </div>
    );
  }
  return (
    <div className="b pad-rem">
      <p>{props.title}</p>
      <p className="font-large">{props.desc}</p>
    </div>
  );
}
