package main

import (
	"fmt"

	admin "github.com/ManuelP84/go_interfaces/admin"
	editor "github.com/ManuelP84/go_interfaces/editor"
	user "github.com/ManuelP84/go_interfaces/user"
)

func auth(user user.User) string {
	permissions := user.Permissions()
	user_name := user.GetName()
	if permissions > 4 {
		return user_name + " has admin permissions. "
	} else if permissions == 3 {
		return user_name + " has editor permissions. "
	}
	return ""

}

func main() {
	editor := editor.Editor{Name: "manuel"}
	admin := admin.Admin{Name: "juan"}
	fmt.Println(auth(editor))
	fmt.Println(auth(admin))

	// Using interface array
	user_interface := []user.User{editor, admin}
	for _, person := range user_interface {
		fmt.Println(auth(person))
	}

}
