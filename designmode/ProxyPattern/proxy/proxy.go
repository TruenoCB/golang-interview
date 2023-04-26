package proxy

type Myserver struct {
	application *application
	maxreq      int
	reqlimiter  map[string]int
}

func (a *Myserver) HandleRequest(m, p string) (int, string) {
	if !a.check(p) {
		return 403, "check fail"
	}
	return a.application.HandleRequest(m, p)
}

func (a *Myserver) check(p string) bool {
	if a.reqlimiter[p] >= a.maxreq {
		return false
	}
	a.reqlimiter[p]++
	return true
}

func GetMyServer(i int) *Myserver {
	return &Myserver{new(application), i, make(map[string]int, 0)}
}
