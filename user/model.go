package user

type Profile struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var users = []Profile{
	{Name: "Harry Potter", Phone: "0"},
	{Name: "The Lord of the Rings", Phone: "1"},
	{Name: "The Wizard of Oz", Phone: "2"},
}
