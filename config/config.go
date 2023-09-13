package config

type Config struct {
    Settings []Setting `json:"settings"`
}

type Setting struct {
    Title string `json:"title"`
    Url   string `json:"url"`
}
