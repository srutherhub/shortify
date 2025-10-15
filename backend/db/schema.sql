CREATE TABLE urls
(
  id TEXT NOT NULL,
  base_route INTEGER NOT NULL DEFAULT 1,
  url TEXT NOT NULL,
  created_at BIGINT NOT NULL,
  expires_at BIGINT,
  click_count BIGINT DEFAULT 0,
  utm_source TEXT,
  utm_medium TEXT,
  utm_campaign TEXT,
  utm_term TEXT,
  utm_content TEXT,
  PRIMARY KEY (id,base_route),
  FOREIGN KEY (base_route) REFERENCES routes(route_id) ON DELETE CASCADE
);

CREATE TABLE routes
(
  route_id SERIAL PRIMARY KEY,
  route_name TEXT NOT NULL
);