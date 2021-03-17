package person

type Person struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
	Animal   string `json:"animal"`
}
