package proxy

type Server interface {
	HandleRequest(string, string) (int, string)
}

type application struct {
}

func (a *application) HandleRequest(m, p string) (int, string) {
	if m == "GET" && p == "1" {
		return 200, "ok"
	} else if m == "Post" && p == "2" {
		return 201, "pok"
	}
	return 404, "not ok"
}
