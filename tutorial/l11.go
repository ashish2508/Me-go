//error

package main
import "fmt"

func main(){
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	profile, err := getUserProfile(user.ID)
		if err != nil {
		fmt.Println(err)
		return
	}
}

func getUser() (User, error){
return User{}, nil
}

func getUserProfile() (UserID int) (Profile, error){
	return Profile{}, nil
}
