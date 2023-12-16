package main

import(
  // "time"
)

type WebsiteChecker func(string) bool

// struct of result type with no named value
// - value are anonymous
type result struct{
  string
  bool
}

// It returns a map of each URL checked to a bool value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool{
  results := make(map[string]bool)
  // a channel to associate the return value of the Websitechecker
  resultChannel := make(chan result)

  for _, url := range urls{
    // creating a go routine as an anonymous func
    // u -> is a copy of url so it cant be changed
    go func(u string){
      // Instead of writing to the map directly,
      // we're sending a result struct for each call to wc to the channel
      // with send statement '<-'
      resultChannel <- result{u, wc(u)}
    }(url)
  }

  // For each of the urls
  for i := 0; i < len(urls); i++{
    // assign channel value to a var r
    r := <-resultChannel
    // update the map
    results[r.string] = r.bool
  }

  // time.Sleep(2 * time.Second)

  return results
}
