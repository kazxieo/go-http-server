package database

type UserDetails struct {
	Age    uint   `json:"age"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

var SampleDatabase = make(map[string]UserDetails)
