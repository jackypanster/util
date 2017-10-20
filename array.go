package util

func Uniq(src []string) []string {
  m := make(map[string]int)
  var dst []string

  for _, item := range src {
    m[item] = 0
  }

  for key := range m {
    dst = append(dst, key)
  }
  return dst
}
