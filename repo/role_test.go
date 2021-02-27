package repo

var _ RoleRepo = roleRepoMock{}

type roleRepoMock struct {
	MockGet   func(id int) string
	MockUsers func(id int) []string
}

func (r roleRepoMock) Get(id int) string {
	return r.MockGet(id)
}

func (r roleRepoMock) Users(id int) []string {
	return r.MockUsers(id)
}
