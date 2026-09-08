package rnatranscription

func ToRNA(dna string) string {
	var s string
    for _, v := range dna {
        switch v {
            case 'G':
            	s += "C"
            case 'C':
            	s += "G"
            case 'T':
            	s += "A"
            case 'A':
            	s += "U"
            default:
            	continue
        }
    }
    return s
}
