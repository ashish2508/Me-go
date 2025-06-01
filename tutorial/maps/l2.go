package main
import "fmt"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUsers, ok := users[name];
	if  !ok {
		return false, fmt.Errorf("user %s not found", name)
	}
	if existingUsers.scheduledForDeletion {
	delete(users, name)
	return true, nil
	}
	return false, nil

}
