func simplifyPath(path string) string {

	splitPath := strings.Split(path, "/")
	answer := []string{}

	for i := 0; i < len(splitPath); i++ {
		switch splitPath[i] {
		case "..":
			if len(answer) == 0 {
				continue
			}
			answer = answer[:len(answer)-1]
		case ".", "":
			// do nothing
			continue
		default:
			answer = append(answer, splitPath[i])
		}
	}
	return "/" + strings.Join(answer, "/")
}