package database

type UserDetails struct {
	Age    uint   `json:"age"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

var SampleDatabase = make(map[string]UserDetails)

// In a real life environment, this wouldn't be hardcoded like this
var PasswordDatabase = map[string]string{
	"kaz": "kaz@3412",
}
