package main

type WebsiteChecker func(string) bool

// It returns a map of each URL checked to a bool value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool{
  results := make(map[string]bool)

  for _, url := range urls{
    results[url] = wc(url)
  }

  return results
}
