package main

import (
	"fmt"
	"github.com/flof/cross_ref_repos_test/repo"
)

type User struct {
	Name string
}

func main() {
	var employeeRepo repo.UserRepo
	var roleRepo repo.RoleRepo

	employeeRepo = repo.NewUserRepo(func() repo.RoleRepo {
		return roleRepo
	})

	roleRepo = repo.NewRoleRepo(func() repo.UserRepo {
		return employeeRepo
	})

	fmt.Println(employeeRepo.Get(17))
	fmt.Println(roleRepo.Get(9))

	employeeRepo.AssignRole(3, 8)
	fmt.Println(roleRepo.Users(13))
}
