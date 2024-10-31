package regexmatching

// IsMatch is a function that checks if a string matches a pattern.
func IsMatch(s string, p string) bool {
  var dfs func(i int, j int) bool
  var dfsCache = make(map[[2]int]bool)
  dfs = func(i int, j int) bool {
    if v, ok := dfsCache[[2]int{i, j}]; ok {
      return v
    }
    if i >= len(s) && j >= len(p) {
      return true
    }
    if j >= len(p) {
      return false
    }

    match := i < len(s) && (s[i] == p[j] || p[j] == '.')

    if (j + 1) < len(p) && p[j + 1] == '*' {
      return dfs(i, j + 2) || (match && dfs(i + 1, j))
    }

    if match {
      return dfs(i+1, j+1)
    }

    return false
  }

  return dfs(0, 0)
}
