# Plan 9 libauth for Go

This is a port of auth(2) from Plan 9 to Go. It's a fork of
bitbucket.org/mischief/libauth which implements missing
pieces as needed. In particular, it adds `auth_userpasswd`
to authenticate a user's password and adds some go doc documentation
based on the man page so that I don't constantly need to cross-reference.

