package regular_expression

func IsMatch(s string, p string) bool {
	if len(p) < 1 {
		return len(s) == 0
	}
	// x* matches empty string or at least one character: x* -> xx*
	// *s is to ensure s is non-empty
	if len(p) > 1 && '*' == p[1] {
		return IsMatch(s, p[2:]) || len(s) > 0 && (p[0] == s[0] || '.' == p[0]) && IsMatch(s[1:], p)
	} else {
		return len(s) > 0 && (p[0] == s[0] || '.' == p[0]) && IsMatch(s[1:], p[1:])
	}
}
