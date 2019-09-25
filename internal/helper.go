package doku

// cross populates a slice of indexes with
// the cross between A and B
func cross(A, B index) []index {
    result := make([]index, len(A)*len(B))
    i := 0
    for _, a := range A {
        for _, b := range B {
            result[i] = index(a) + index(b)
            i++
        }
    }
    return result
}
