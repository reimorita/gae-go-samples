package pkg

import (
	"net/http"
	"fmt"
	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/auth", auth)
}

func auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	// Context
	ctx := appengine.NewContext(r)
	// Get Login User
	u := user.Current(ctx)
	// Not Loggined
	if u == nil {
		// Create Login Url
		loginUrl, _ := user.LoginURL(ctx, "/auth")
		// Output Login Url
		fmt.Fprintf(w, `<a href="%s">Sign in</a>`, loginUrl)
		return
	}
	// Create Logout Url
	logoutUrl, _ := user.LogoutURL(ctx, "/auth")
	// Output Logout Url
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">Sign out</a>)`, u.Email, logoutUrl)
}

