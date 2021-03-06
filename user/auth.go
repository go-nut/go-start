package user

import "github.com/ungerik/go-start/view"

///////////////////////////////////////////////////////////////////////////////
// Auth

type Auth struct {
	LoginURL view.URL
}

func (self *Auth) Authenticate(ctx *view.Context) (ok bool, err error) {
	id, ok := ctx.Session.ID()
	if !ok {
		return false, nil
	}

	ok, err = IsConfirmedUserID(id)
	if !ok && err == nil && self.LoginURL != nil {
		err = view.Redirect(self.LoginURL.URL(ctx))
	}
	return ok, err
}
