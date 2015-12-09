package main

type Params struct {
	APIKey   string `json:"api_key"`
	Name     string `json:"name"`
	Gemspec  string `json:"gemspec"`
	Host     string `json:"host"`
}
