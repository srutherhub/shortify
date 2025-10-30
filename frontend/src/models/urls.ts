export class ShortUrl {
  ID: string;
  Url: string;
  RouteName: string;
  CreatedAt: number;
  ExpiresAt: number | null;
  UtmSource: string | null;
  UtmMedium: string | null;
  UtmCampaign: string | null;
  UtmTerm: string | null;
  UtmContent: string | null;

  constructor(data: ShortUrl) {
    this.ID = data.ID;
    this.Url = data.Url;
    this.RouteName = data.RouteName;
    this.CreatedAt = data.CreatedAt;
    this.ExpiresAt = data.ExpiresAt;
    this.UtmSource = data.UtmSource;
    this.UtmMedium = data.UtmMedium;
    this.UtmCampaign = data.UtmCampaign;
    this.UtmTerm = data.UtmTerm;
    this.UtmContent = data.UtmContent;
  }
  getShortUrl(): string {
    return this.Url + this.RouteName + this.ID;
  }
}
