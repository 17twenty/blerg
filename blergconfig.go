package main

// Generated from https://xuri.me/toml-to-go/

type BlergConfig struct {
	General struct {
		Linkedin string `toml:"linkedin"`
		Github   string `toml:"github"`
		Email    string `toml:"email"`
	} `toml:"general"`
	Post []struct {
		PubDate  string `toml:"pub_date"`
		Article  string `toml:"article"`
		Image    string `toml:"image"`
		Title    string `toml:"title"`
		Subtitle string `toml:"subtitle"`
	} `toml:"post"`
}
