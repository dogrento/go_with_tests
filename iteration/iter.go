package iteration

// Simple func that iters a characters
func Repeat(chars string) string{
  // declaring variable typed as string
  var rep string
  
  for i := 0; i < 5; i++{
    //  following line is the same as rep = rep + chars (Add AND)
    rep += chars
  }
  return rep 
}
