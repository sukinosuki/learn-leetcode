package main

import "fmt"

func countWords(words1 []string, words2 []string) int {

	cnt1 := make(map[string]int)
	cnt2 := make(map[string]int)

	for _, word := range words1 {
		cnt1[word]++
	}
	for _, word := range words2 {
		cnt2[word]++
	}

	if len(words2) < len(words1) {
		words1, words2 = words2, words1
	}

	ans := 0
	for _, word := range words1 {
		if cnt1[word] == 1 && cnt2[word] == 1 {
			ans++
		}
	}

	return ans
}

func main() {
	// 2
	// words1 := []string{"leetcode", "is", "amazing", "as", "is"}
	// words2 := []string{"amazing", "leetcode", "is"}
	// 1
	//words1 := []string{"a", "ab"}
	//words2 := []string{"a", "a", "a", "ab"}
	words1 := []string{"ibxyatvglhltxngewrcrqbbnew", "towokpjpkccmob", "kdmtwngzpebwpnvlazhdcmtwpjh", "muh", "fzzlmacbbvoqdueutjqoutwd", "ylluspdftxxqbwadivfdzulghq", "hioiacezaiptpsvcojzckhxz", "nzcjhjomaupevyekennyrfwyd", "tdwtuinstwsfyjnfkxkbnsptisuifo", "wrdwoxzsczzlnwjugopohxh", "p", "jkez", "drisymx", "fsva", "myqc", "aovjoxzpkylpecltwtottzidq", "wqspbhpberqjabockesc", "f", "qostobxgfliil", "gsekmhjpuedeivioudx", "tzelzowtgnvjsxgbw", "zgmpazgnioprk", "fucybddarjcve", "ldacfviysy", "yxyjairoxtvbkljaokca", "vxpiohhvjuwcpiceafcdzobalgpz", "wyflbpmkfwftndgtnftajgla", "xbxvvk", "bnrwyshimjamltmlugeiviu", "wsgqysmuakedrrmjk", "ppqmgibqljkwgmiwi", "fly", "uf", "tvvttzrsjbojve", "ztxtnmljdhyz", "vxonvloufeksfvg", "wql", "kotdenqjrdlgofubocb", "wlaqceczd", "mtmhtgvqwr", "aymzxpfvbqxydmilafyqvapuxtnqe", "ig", "atetjlhdcigunmmit", "enkdcxqnw", "gtlcmkxwvdhumgfurxkesmekmnhjo", "hurihasxncgtzleerslvwxkz", "zked", "xiaqvclhuhggcgoouzjgi", "mzejuubgyhvlfbecpmggddby", "boyotuukuiprtlvktypxboosw", "vwfceei", "gopsxsihawzhtlmdyiggljzggrhqr", "bckuuqszgncdhkeghudflczm", "e", "yvhwysrunkxsppbqjf", "lo", "bze", "kuzoqvgugnrpfkelktfg", "ntjtlwwmuevtsqujpxswgx", "zkdwtpdlvrfkbyktqsellmghaxj", "u", "rpmpq", "ajhlzwfrbysqloduofr", "gyfmhcskcrjepgeplbbj", "fe", "zyolvtetrdffy", "apbkyczsuvde", "fnkqf", "qwwxpwbr", "krkbnww", "zkvqkugfpziawiokdzlpjomfarkor", "jg", "l", "srbvxsnuhyqzmycvavmmakh", "dhkgzjrstir", "smaaptkzpwhukebwboysbnawgzgot"}
	words2 := []string{"p", "towokpjpkccmob", "vflbjyecpzxnuay", "fzas", "fzzlmacbbvoqdueutjqoutwd", "bwjjzw", "va", "manrvuldjzrdnwihzct", "tdwtuinstwsfyjnfkxkbnsptisuifo", "wrdwoxzsczzlnwjugopohxh", "p", "tylcyihdjruhaayzcwxrynnkch", "uojzddpgyvqslha", "fsva", "rucvbjzfewjlhddxefhf", "tfihr", "wqspbhpberqjabockesc", "f", "bmfo", "zsjbzjmbloaybdofsrqvzwoizz", "tzelzowtgnvjsxgbw", "tproznqma", "lmryjiyvkgsxsaylkdmmxeub", "ldacfviysy", "xpamoswlugwjxyny", "rmfvgm", "wm", "xbxvvk", "ubawz", "jbrabb", "rgegpb", "fly", "aofydpklgjqmxhvxuhq", "tvvttzrsjbojve", "wj", "vxonvloufeksfvg", "wql", "vu", "nhuxqdfyftrbbodztyydb", "mtmhtgvqwr", "aymzxpfvbqxydmilafyqvapuxtnqe", "fqksatpfo", "ylzkfvvzdsryl", "enkdcxqnw", "gtlcmkxwvdhumgfurxkesmekmnhjo", "nccwybkxuawwdqyhrhmbt", "zked", "eyzwtvsjt", "qy", "boyotuukuiprtlvktypxboosw"}
	res := countWords(words1, words2)

	fmt.Printf("res: %v\n", res)
}
