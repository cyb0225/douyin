/*
token : To ensure user has already login
I put token file (which store user login) in contorller floder, in order to judge quickly if user is logined,
instead of pass token to service even repository(database) to judege 
*/

package commonctl

type UserLoginComp struct {
	UserId uint64 
}


var (
	UserLoginMap = make(map[string]UserLoginComp)
)


func CreatToken(username, password string) string {

	return username + password
}