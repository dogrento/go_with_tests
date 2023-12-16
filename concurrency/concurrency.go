package main

import(
  // "time"
)

type WebsiteChecker func(string) bool
type result struct{
  string
  bool
}

// It returns a map of each URL checked to a bool value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool{
  results := make(map[string]bool)
  resultChannel := make(chan result)

  for _, url := range urls{
    // creating a go routine as an anonymous func
    // u -> is a copy of url so it cant be changed
    go func(u string){
      // results[u] = wc(u)
      resultChannel <- result{u, wc(u)}
    }(url)
  }

  for i := 0; i < len(urls); i++{
    r := <-resultChannel
    results[r.string] = r.bool
  }

  // time.Sleep(2 * time.Second)

  return results
}
