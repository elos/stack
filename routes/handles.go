package routes

import (
	"log"
	"net/http"

	"github.com/elos/data"
	"github.com/elos/stack/util"
	"github.com/elos/stack/util/auth"
	"github.com/elos/transfer"
	"github.com/julienschmidt/httprouter"
)

func Null() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	}
}

func Error(err error) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		util.ServerError(w, err)
	}
}

func BadMethod() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		util.InvalidMethod(w)
	}
}

func BadAuth(reason string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		util.Unauthorized(w)
	}
}

type AuthHandle func(http.ResponseWriter, *http.Request, httprouter.Params, data.Identifiable, data.Store)

func Auth(h AuthHandle, s data.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		agent, authenticated, err := auth.AuthenticateRequest(s, r)
		if err != nil {
			log.Printf("An error occurred during authentication, err: %s", err)
			// h.NewErrorHandler(err).ServeHTTP(w, r)
			return
		}

		if authenticated {
			h(w, r, ps, agent, s)
			log.Printf("Agent with id %s authenticated", agent.ID())
		} else {
			log.Printf("Agent with id %s authenticated", agent.ID())
			//	h.NewUnauthorizedHandler("Not authenticated").ServeHTTP(w, r)
		}

	}
}

func Post(k data.Kind, params []string) AuthHandle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, a data.Identifiable, s data.Store) {
		var attrs data.AttrMap

		for _, k := range params {
			attrs[k] = r.FormValue(k)
		}

		c := transfer.NewHTTPConnection(w, r, a)
		e := transfer.New(c, transfer.POST, k, attrs)
		go transfer.Route(e, s)
	}
}

func Values(p ...string) []string {
	return p
}
