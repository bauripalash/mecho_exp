package bengali

import "sort"

const (
	BN_URANGE_START       = rune('ঀ') //U+0980
	BN_URANGE_VISIBLE_END = rune('৽') //U+09FD
	BN_URANGE_END         = rune('\u09ff')
)

var (
	// Reserved Characters of Bengali Unicode Range
	BN_URANGE_RESERVED = []rune{
		'\u0984', '\u098D', '\u098E', '\u0991',
		'\u0992', '\u0992', '\u09A9',
		'\u09B1', '\u09B1', '\u09B3',
		'\u09B4', '\u09B5', '\u09BA',
		'\u09BB', '\u09C5', '\u09C6',
		'\u09C9', '\u09CA', '\u09CF',
		'\u09D0', '\u09D1', '\u09D2',
		'\u09D3', '\u09D4', '\u09D5',
		'\u09D6', '\u09D8', '\u09D9',
		'\u09DA', '\u09DB', '\u09DE',
		'\u09E4', '\u09E5', '\u09FF',
	}
	BN_VOWELS = []rune{
		'\u0985', '\u0986', '\u0987', '\u0988', '\u0989', '\u098A',
		'\u098B', '\u098C', '\u098F', '\u0990', '\u0993', '\u0994',
	}
	BN_VOWELS_MAP = map[string]rune{
		"a":  BN_VOWELS[0],
		"aa": BN_VOWELS[1],
		"i":  BN_VOWELS[2],
		"ii": BN_VOWELS[3],
		"u":  BN_VOWELS[4],
		"uu": BN_VOWELS[5],
		"vr": BN_VOWELS[6],
		"vl": BN_VOWELS[7],
		"e":  BN_VOWELS[8],
		"ai": BN_VOWELS[9],
		"o":  BN_VOWELS[10],
		"au": BN_VOWELS[11],
	}

	BN_CONSONANTS = []rune{
		'\u0995', '\u0996', '\u0997', '\u0998', '\u0999', '\u099A',
		'\u099B', '\u099C', '\u099D', '\u099E', '\u099F', '\u09A0',
		'\u09A1', '\u09A2', '\u09A3', '\u09A4', '\u09A5', '\u09A6',
		'\u09A7', '\u09A8', '\u09AA', '\u09AB', '\u09AC', '\u09AD',
		'\u09AE', '\u09AF', '\u09B0', '\u09B2', '\u09B6', '\u09B8',
		'\u09B8', '\u09B9', '\u09DC', '\u09DD', '\u09DF',
	}

	BN_CONSONANTS_MAP = map[string]rune{
		"ka":   BN_CONSONANTS[0],
		"kha":  BN_CONSONANTS[1],
		"ga":   BN_CONSONANTS[2],
		"gha":  BN_CONSONANTS[3],
		"nga":  BN_CONSONANTS[4],
		"ca":   BN_CONSONANTS[5],
		"cha":  BN_CONSONANTS[6],
		"ja":   BN_CONSONANTS[7],
		"jha":  BN_CONSONANTS[8],
		"nya":  BN_CONSONANTS[9],
		"tta":  BN_CONSONANTS[10],
		"ttha": BN_CONSONANTS[11],
		"dda":  BN_CONSONANTS[12],
		"ddha": BN_CONSONANTS[13],
		"nna":  BN_CONSONANTS[14],
		"ta":   BN_CONSONANTS[15],
		"tha":  BN_CONSONANTS[16],
		"da":   BN_CONSONANTS[17],
		"dha":  BN_CONSONANTS[18],
		"na":   BN_CONSONANTS[19],
		"pa":   BN_CONSONANTS[20],
		"pha":  BN_CONSONANTS[21],
		"ba":   BN_CONSONANTS[22],
		"bha":  BN_CONSONANTS[23],
		"ma":   BN_CONSONANTS[24],
		"ya":   BN_CONSONANTS[25],
		"ra":   BN_CONSONANTS[26],
		"la":   BN_CONSONANTS[27],
		"sha":  BN_CONSONANTS[28],
		"ssa":  BN_CONSONANTS[29],
		"sa":   BN_CONSONANTS[30],
		"ha":   BN_CONSONANTS[31],
		"rra":  BN_CONSONANTS[32],
		"rha":  BN_CONSONANTS[33],
		"yya":  BN_CONSONANTS[34],
		// Extra Signs
	}

	BN_SIGNS = map[string]rune{
		"chandrabindu": '\u0981',
		"anusvara":     '\u0982',
		"visarga":      '\u0983',
	}

	BN_VOWEL_SIGNS = []rune{
		'\u09BE', '\u09BF', '\u09C0', '\u09C1', '\u09C2', '\u09C3',
		'\u09C4', '\u09C7', '\u09C8', '\u09CB', '\u09CC',
	}
	BN_VOWEL_SIGNS_MAP = map[string]rune{
		"aa":  BN_VOWEL_SIGNS[0],
		"i":   BN_VOWEL_SIGNS[1],
		"ii":  BN_VOWEL_SIGNS[2],
		"u":   BN_VOWEL_SIGNS[3],
		"uu":  BN_VOWEL_SIGNS[4],
		"vr":  BN_VOWEL_SIGNS[5],
		"vrr": BN_VOWEL_SIGNS[6],
		"e":   BN_VOWEL_SIGNS[7],
		"ai":  BN_VOWEL_SIGNS[8],
		"o":   BN_VOWEL_SIGNS[9],
		"au":  BN_VOWEL_SIGNS[10],
	}

	BN_NUMS = []rune{
		'\u09E6', '\u09E7', '\u09E8', '\u09E9', '\u09EA',
		'\u09EB', '\u09EC', '\u09ED', '\u09EE', '\u09EF',
	}

	BN_NUMS_MAP = map[string]rune{
		"0": BN_NUMS[0],
		"1": BN_NUMS[1],
		"2": BN_NUMS[2],
		"3": BN_NUMS[3],
		"4": BN_NUMS[4],
		"5": BN_NUMS[5],
		"6": BN_NUMS[6],
		"7": BN_NUMS[7],
		"8": BN_NUMS[8],
		"9": BN_NUMS[9],
	}

)


func GetMergedMap(incNums, incSigns bool) map[string]rune{
	result := make(map[string]rune)
	
	for idx , item := range BN_VOWELS_MAP{
		result[idx] = item
	}

	for idx , item := range BN_CONSONANTS_MAP{
		result[idx] = item
	}

	for idx , item := range BN_SIGNS{
		result[idx] = item
	}

	if incNums{
		
		for idx , item := range BN_NUMS_MAP{
			result[idx] = item
		}
		
	}

	if incSigns{
		for idx , item := range BN_VOWEL_SIGNS_MAP{
			result[idx+"_"] = item
		}
	}
    keys := make([]string , 0 , len(result))
	
	for k := range result{
		keys = append(keys, k)
	} 

	sort.SliceStable(keys , func(i, j int) bool {
        return result[keys[i]] < result[keys[j]]
	})
	return result
}

// Helper function
// Check If `C` (`rune`) is in `Array` `arr`
func doesContain(arr []rune, c rune) bool {
	for _, i := range arr {
		if i == c {
			return true
		}
	}

	return false
}

// Check if C (rune) is Bengali Unreserved
// and visible character of Bengali Unicode
// Range
func IsUnresBnChar(c rune) bool {
	if IsBnChar(c) {
		if !doesContain(BN_URANGE_RESERVED, c) {
			return true
		}
	}
	return false
}

// Check if C (rune) is in Bengali Unicode
// Range.
// It contains all characters including 
// reserved invisible characters.
// To check only Visible characters use:
// IsUnresBnChar(C)
func IsBnChar(c rune) bool {
	if c >= BN_URANGE_START && c <= BN_URANGE_END {
		return true
	}

	return false
}

// Check if C (rune) is a bengali number
func IsBnNum(c rune) bool {
	return c >= BN_NUMS[0] && c <= BN_NUMS[9]
}


// Check if C (rune) is a bengali vowel 
func IsBnVowel(c rune) bool {
	if IsUnresBnChar(c) {
		if c >= BN_VOWELS[0] && c <= BN_VOWELS[len(BN_VOWELS)-1] {
			return true
		}
	}

	return false
}


// Check if C (rune) is a bengali consonant.
// It also includes 'chandrabindu', 'anusvara'
// and 'visarga'
func IsBnCons(c rune) bool {
	if IsUnresBnChar(c) {
		if c >= BN_CONSONANTS[0] && c <= BN_CONSONANTS[len(BN_VOWELS)-1] {
			return true
		} else if c >= BN_SIGNS["chandrabindu"] && c <= BN_SIGNS["visarga"] {
			return true
		}
	}

	return false
}
