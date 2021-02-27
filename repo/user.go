package repo

import (
	"fmt"
	"log"
	"math/rand"
)

type UserRepo interface {
	Get(id int) string
	AssignRole(id, roleID int)
}

var _ UserRepo = userRepoImpl{}

type userRepoImpl struct {
	ID int
	RoleRepo func() RoleRepo
}

func NewUserRepo(roleRepo func() RoleRepo) UserRepo {
	return userRepoImpl{ID: rand.Int(), RoleRepo: roleRepo}
}

func (e userRepoImpl) Get(id int) string {
	return fmt.Sprintf("User %d", id)
}

func (e userRepoImpl) AssignRole(id, roleID int) {
	role := e.RoleRepo().Get(roleID)
	user := e.Get(id)
	log.Printf("Assign %s to %s", role, user)
}