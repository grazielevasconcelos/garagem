package model

type Skill struct {
	Id             int    `json:"id"`
	Description    string `json:"description"`
	ExpertiseLevel int    `json:"expertiseLevel"`
}

type Member struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Skill []Skill `json:"skill"`
}
type Video struct {
	Id          int      `json:"id"`
	Client      string   `json:"client"`
	Duration    string   `json:"duration"`
	BrandUnity  string   `json:"brandUnity"`
	Url         string   `json:"url"`
	Context     string   `json:"context"`
	Description string   `json:"description"`
	Background  string   `json:"background"`
	Member      []Member `json:"member"`
}
