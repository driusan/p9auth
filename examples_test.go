package p9auth

func ExampleGetuserpasswd() {
	val, err := Getuserpasswd(
		"proto=pass service=git role=client server=%s",
		"https://example.com",
	)
	if err != nil {
		return
	}
	// val.User is now the password for the service git
	// on https://example.com as retrieved via factotum,
	// val.Password is the password
	println(val.User)
}
func ExampleUserpasswd() {
	_, err := Userpasswd("glenda", "badpassword")
	if err != nil {
		// Invalid password
		return
	}
	// Valid password
}
