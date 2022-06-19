/*
token : To ensure user has already login
*/

package commonctl

//TODO: 
func CreatToken(username, password string) string {

	password = MD5(password)
	return username + password
}
