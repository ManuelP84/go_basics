package admin

type Admin struct {
	Name string
}

func (admin Admin) Permissions() int {
	return 5
}

func (admin Admin) GetName() string {
	return admin.Name
}
