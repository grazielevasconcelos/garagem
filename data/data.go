package data

import (
	"garagem/model"
)

var (
	foundedVideo model.Video
)

func Populate(paramJSON string) (foundedVideo model.Video) {
	SkillDev := model.Skill{
		Id:             2,
		Description:    "Node JS",
		ExpertiseLevel: 1,
	}
	MemberPopular := model.Member{
		Id:    10,
		Name:  "Max",
		Age:   34,
		Skill: []model.Skill{SkillDev},
	}
	VideoLearn := model.Video{
		Id:          1,
		Url:         "https://www.youtube.com/embed/klO9kWiq7rU",
		Context:     "learn",
		Description: "<h2>Learn</h2> <h4>Continuously deliver the right solution</h4> Continuously gain new insights from your customers interaction with your application and the metrics you collect to drive business decisions.",
		Background:  "#FF5050",
		Member:      []model.Member{MemberPopular},
	}
	SkillDevops := model.Skill{
		Id:          4,
		Description: "Golang",
	}
	MemberPerson := model.Member{
		Id:    11,
		Name:  "Anakyn Skywalker",
		Age:   23,
		Skill: []model.Skill{SkillDevops},
	}
	VideoThink := model.Video{
		Id:          2,
		Url:         "https://www.youtube.com/embed/9Jgmxhgg5i0",
		Context:     "think",
		Description: "<h2>Think</h2> <h4>Incrementally deliver awesome solutions</h4> Adopt IBM Design Thinking to conceptualize, design, refine and prioritize features that will delight your customers.",
		Background:  "#FBB731",
		Member:      []model.Member{MemberPerson},
	}
	switch paramJSON {
	case "1":
		foundedVideo = VideoLearn
	case "2":
		foundedVideo = VideoThink
	}
	return foundedVideo
}
