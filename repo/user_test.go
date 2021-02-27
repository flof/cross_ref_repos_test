package repo

import (
	"fmt"
	"testing"
)

func TestAssignRole(t *testing.T) {
	repo := userRepoImpl{
		ID: 1,
		RoleRepo: func() RoleRepo {
			return roleRepoMock{
				MockGet: func(id int) string {
					return fmt.Sprintf("Mocked role %d", id)
				},
			}
		},
	}

	repo.AssignRole(1, 2)

	// TODO: verify test db
}
