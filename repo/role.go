package repo

import (
	"fmt"
	"math/rand"
)

type RoleRepo interface {
	Get(id int) string
	Users(id int) []string
}

var _ RoleRepo = roleRepoImpl{}

type roleRepoImpl struct {
	ID       int
	UserRepo func() UserRepo
}

func NewRoleRepo(userRepo func() UserRepo) RoleRepo {
	return roleRepoImpl{ID: rand.Int(), UserRepo: userRepo}
}

func (r roleRepoImpl) Get(id int) string {
	return fmt.Sprintf("Role %d", id)
}

func (r roleRepoImpl) Users(id int) []string {
	return []string{r.UserRepo().Get(1), r.UserRepo().Get(2)}
}
