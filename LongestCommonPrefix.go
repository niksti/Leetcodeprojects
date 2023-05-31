func longestCommonPrefix(strs []string) string {
   
   prefix := ""
    for index, a := range strs[0]{

        for _, b := range strs{
            if index == len(b) {
                return prefix
            }
            if a != []rune(b)[index]{
                return prefix
            }
        }
        prefix += string(a)
    }
    return prefix
}
