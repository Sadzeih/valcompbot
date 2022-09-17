package highlighter

type FlairedUser struct {
	Username string
	Role     string
}

type userMap map[string]FlairedUser

func (m userMap) Add(u FlairedUser) {
	m[u.Username] = u
}

func (m userMap) Delete(u string) {
	delete(m, u)
}

func (s userMap) Len() int {
	return len(s)
}

func (s userMap) Exists(u string) bool {
	_, ok := s[u]
	return ok
}

var (
	rioters = userMap{
		"Sadzeih": {Username: "Sadzeih", Role: "ValComp Mod"},
	}
)
