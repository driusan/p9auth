package p9auth

import (
	"fmt"
)

// Implements auth_userpasswd from auth(2).
//
// Userpasswd validates a simple username/password pair.
func Userpasswd(user, password string) (*AuthInfo, error) {
	f, err := openRPC()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rpc := &AuthRpc{F: f}

	if r, _ := rpc.Rpc("start", "proto=dp9ik role=login"); r != ARok {
		if r, errresp := rpc.Rpc("start", "proto=p9sk1 role=login"); r != ARok {
			return nil, fmt.Errorf(errresp)
		}
	}
	if r, errresp := rpc.Rpc("write", user); r != ARok {
		return nil, fmt.Errorf(errresp)
	}
	if r, errresp := rpc.Rpc("write", password); r != ARok {
		return nil, fmt.Errorf(errresp)
	}

	if err := rpc.GetInfo(); err != nil {
		return nil, err
	}
	return rpc.Ai, nil
}
