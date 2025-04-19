package models

type Data struct {
  Value     string `json:"value"`
  CreatedAt string `json:"created_at,omitempty"`
  TTL       int    `json:"ttl,omitempty"`
}