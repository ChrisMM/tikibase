package helpers

import "net/url"

// IsURL indicates whether the given string is a URL.
func IsURL(link string) bool {
	_, err := url.ParseRequestURI(link)
	return err == nil
}
