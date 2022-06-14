/*
token : To ensure user has already login
*/

package commonctl

type UserLoginComp struct {
	Id uint64
}

var (
	UserLoginMap = make(map[string]UserLoginComp)
)

//TODO: 
func CreatToken(username, password string) string {

	password = MD5(password)
	return username + password
}
