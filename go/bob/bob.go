package bob

import "regexp"

const testVersion = 3

func Hey(in string) string {
  if len(in) == 0 {
    return "Fine. Be that way!"
  }

  nothing, _ := regexp.MatchString("^[.\\s-_0-9\\t]+$", in)
  if nothing {
    return "Fine. Be that way!"
  }

  exclamation, _ := regexp.MatchString("^([A-Z\\s]+|[A-Z0-9\\,\\s\\%\\^\\*\\@\\#\\$\\(\\!]{2,}[\\!\\?]{1})$", in)
  if exclamation {
    return "Whoa, chill out!"
  }

  question, _ := regexp.MatchString("\\?[\\s]*$", in)
  if question {
    return "Sure."
  }

  return "Whatever."
}
