package abstraction

import "strings"

const KeyDelimiter = ":"

func Combine(pathSegments ...string) string {
	return strings.Join(pathSegments, KeyDelimiter)
}

func GetSectionKey(path string) string {
	if path == "" {
		return path
	}
	lastIndex := strings.LastIndex(path, KeyDelimiter)
	if lastIndex == -1 {
		return path
	}
	return path[lastIndex+1:]
}

func GetParentPath(path string) string {
	if path == "" {
		return path
	}
	lastIndex := strings.LastIndex(path, KeyDelimiter)
	if lastIndex == -1 {
		return ""
	}
	return path[:lastIndex]
}
