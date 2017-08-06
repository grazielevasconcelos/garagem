package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"fmt"
)

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

func handler(w http.ResponseWriter, r *http.Request) {
	skill := Skill{
		Id:             2,
		Description:    "Node JS",
		ExpertiseLevel: 1,
	}

	member := Member{
		Id:    10,
		Name:  "Max",
		Age:   34,
		Skill: []Skill{skill},
	}

	video := Video{
		Id:          1,
		Url:         "https://www.youtube.com/embed/klO9kWiq7rU",
		Context:     "learn",
		Description: "<h2>Learn</h2> <h4>Continuously deliver the right solution</h4> Continuously gain new insights from your customers interaction with your application and the metrics you collect to drive business decisions.",
		Background:  "#FF5050",
		Member:      []Member{member},
	}
	skill2 := Skill{
		Id:          4,
		Description: "Golang",
	}

	member2 := Member{
		Id:    11,
		Name:  "Anakyn Skywalker",
		Age:   23,
		Skill: []Skill{skill2},
	}

	video2 := Video{
		Id:          2,
		Url:         "https://www.youtube.com/embed/9Jgmxhgg5i0",
		Context:     "think",
		Description: "<h2>Think</h2> <h4>Incrementally deliver awesome solutions</h4> Adopt IBM Design Thinking to conceptualize, design, refine and prioritize features that will delight your customers.",
		Background:  "#FBB731",
		Member:      []Member{member2},
	}

	/*
	skill3 := Skill{
		Id:          5,
		Description: "Java",
	}

	member3 := Member{
		Id:    12,
		Name:  "Brian Johson",
		Age:   19,
		Skill: []Skill{skill3},
	}

	video3 := Video{
		Id:          3,
		Url:         "https://www.youtube.com/embed/PvQDAYrQe6w",
		Context:     "code",
		Description: "<h2>Code</h2> <h4>Create innovative solutions fast</h4> Adopt development practices to build cloud-native applications, release incremental function, gather feedback, and measure results.",
		Background:  "#19A69F",
		Member:      []Member{member3},
	}
	skill4 := Skill{
		Id:          5,
		Description: "Python",
	}

	member4 := Member{
		Id:    17,
		Name:  "John",
		Age:   29,
		Skill: []Skill{skill4},
	}

	video4 := Video{
		Id:          5,
		Url:         "https://www.youtube.com/watch?v=BWyDA5y6TIg",
		Context:     "code",
		Description: "<h2>Code</h2> <h4>Create innovative solutions fast</h4> Adopt development practices to build cloud-native applications, release incremental function, gather feedback, and measure results.",
		Background:  "#349643",
		Member:      []Member{member4},
	}
	*/
	w.Header().Set("Context-Type", "aplication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	paramJson := r.URL.Query().Get("Id")	

	switch paramJson{
		case "1":
				j, err := json.Marshal(video)
					if err != nil {
					panic(err)
				}
				//fmt.Println(reflect.TypeOf(j).Name())
				w.Write(j)
		case "2":
				k, err := json.Marshal(video2)
				if err != nil {
					panic(err)
				}
				w.Write(k)	
				//fmt.Println("Primerio Json: ", string(k))
	}
	/*
	l, err := json.Marshal(video3)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Context-Type", "aplication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(l)
	//fmt.Println("Primerio Json: ", string(l))
	m, err := json.Marshal(video4)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Context-Type", "aplication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(m)		
	*/
	//fmt.Println("Primerio Json: ", string(m))
}

func main() {

/* 	var videos []Video
	videos[0] = video
	videos[1] = video2
	videos[2] = video3
	videos[3] = video4

	for vid := range videos {
		b, err := json.Marshal(vid)
		if err != nil {
			panic(err)
		}
		fmt.Println("Primerio Json: ", string(b))
	}
 */
	http.HandleFunc("/", handler)

	http.ListenAndServe(":9000", nil)
}
