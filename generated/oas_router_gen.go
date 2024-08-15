// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'd': // Prefix: "dummyLogin"
				origElem := elem
				if l := len("dummyLogin"); len(elem) >= l && elem[0:l] == "dummyLogin" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleDummyLoginGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 'f': // Prefix: "flat/"
				origElem := elem
				if l := len("flat/"); len(elem) >= l && elem[0:l] == "flat/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					origElem := elem
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleFlatCreatePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'u': // Prefix: "update"
					origElem := elem
					if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleFlatUpdatePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "house/"
				origElem := elem
				if l := len("house/"); len(elem) >= l && elem[0:l] == "house/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					origElem := elem
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleHouseCreatePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}
				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleHouseIDGetRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/subscribe"
					origElem := elem
					if l := len("/subscribe"); len(elem) >= l && elem[0:l] == "/subscribe" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleHouseIDSubscribePostRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'l': // Prefix: "login"
				origElem := elem
				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleLoginPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

				elem = origElem
			case 'r': // Prefix: "register"
				origElem := elem
				if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleRegisterPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'd': // Prefix: "dummyLogin"
				origElem := elem
				if l := len("dummyLogin"); len(elem) >= l && elem[0:l] == "dummyLogin" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = "DummyLoginGet"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/dummyLogin"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'f': // Prefix: "flat/"
				origElem := elem
				if l := len("flat/"); len(elem) >= l && elem[0:l] == "flat/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					origElem := elem
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "FlatCreatePost"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/flat/create"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'u': // Prefix: "update"
					origElem := elem
					if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "FlatUpdatePost"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/flat/update"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "house/"
				origElem := elem
				if l := len("house/"); len(elem) >= l && elem[0:l] == "house/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					origElem := elem
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "HouseCreatePost"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/house/create"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}
				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "HouseIDGet"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/house/{id}"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/subscribe"
					origElem := elem
					if l := len("/subscribe"); len(elem) >= l && elem[0:l] == "/subscribe" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "HouseIDSubscribePost"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/house/{id}/subscribe"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 'l': // Prefix: "login"
				origElem := elem
				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = "LoginPost"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/login"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'r': // Prefix: "register"
				origElem := elem
				if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = "RegisterPost"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/register"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}