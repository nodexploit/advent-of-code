package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	datastream := "mnlnvlljqqccznnjtjljbllrtllwwpmmhjjbbzppnndmmsppdqqwvvstvssgmsggmlmttnvvfbbdsssnzzbssjrsjjpmpvmmcjjwsssndsslwsswtwnwrrslshhvzzsppffmpfmmfvfpfpsssqpqzpqqcjcjnjcnnzbzjzpzbpbnbwbcctvvhgvgsvvpwwvjjvqjjjdqqrmrmqmsqszqsqpsqslsddhbhcbhbchcvvjvjcjnccdbcdcrddldblbffhvffpvpzpvvmvfmmwhwqhqvhqhmmpdmmlbmbgmbbrqqpmqqcvcmvvcncllptltvtdtbbqzzcggjgsjjvjsvvgmgffqhqqgpptspsffvdvbbhqhzzllvvjbvvbpppggfpgptgtvvzdvdzdgzgccmmphpmhppldlnlpnnhghhrrgwrwssnllmpllbvbvqvtvhtvvmnvvpgvgfvggtztpthhcfhfqfhhnhtnhhljhjppqjjffgjggrwrjjhphzhtztggwswnwzzvbzzmmtrtqtjqttwlwmmmmnddmnddwvvcllgrgfgzznwnsswjjhwwspsbbvzzqvzvbvcvmmtltnlnfnfnwwsvwwpswppjhjdhhmbbblfbfwffwvvjgghwhzhjzzrttwhwjhjchjhggdrgdgmmsjsfstftvtmtctggcwgcgzccgzgsgrgmmjhhqzzrmrttgtgfgcffvsfvvslsvvpqvvnjnrrwdwcwcnnhllwpwdpwdpdqqtwtftdftddppncpnccllqqrffpssgvsvtvmvssrbrhbbzggtssdsvddqfdfjfhjhdjhhncnddfpdfdmmrddncnvcnvccgvvhzvzwztwzwtzwtwqttrlrvrddztzrzcccgmmqgqjgqjgjqqspqqpjppbggchcqcpqqgbgdbbspbbrbhrhzhqhrqhhhtbhhvshspsvsggjdjwjwvjvdvwddjggmrrbnrrztthlllhlclbclbbhpbhphjhccdwccdbbjrbjjmrjrhhnlnjllltwltlmmlqlnqntqnntsnsqqvtvwvgwvwnncgcdctddnttfjfqfttrhrjhjqqcnnsmnmgmqgmgbblcctntrntnccnvvmpmjjvfjfrrbpbttsbttvnttmnnjdnndnzntnrnwrwcctllvhvqhhddmzztppphghphzzglzlnnfccrfffvvhllpspwssstwstsvttcrtccfssbccdjdqdfqddrbrqbbtllmmsfmfcmmzwzpznpnttjgtgbbdtdvdwwpmphpprsrjrbrqbqwwljlslrlrhhpchpcprcrtcrcfcfssndsspddcjjjfmfqfggmssnhsnhnpncnfnmffdrdjdhjjrgjrjgjqgjqqlmlljffbcfcrrrzwwftwtrrpgpprqqmrmhmwhwmmcrrhqrqwqppwjwggpdpgpvgvzvttqlljhhbvbhblhhcsssvmsmppcvpvrvzvbvtbtssplpgptgtnthhvwhvwwfvvfwwmtwmwfmfgmgnmnllgsgmsgswwhqhhhzqhhfwwnttmfmrrfnnpbpssvbsshqhqvhqqbmbpmbmqqjtjqjvjtjjhtttpzphpqqwqfqttqqhfhbfhfwwcpcpssdvvzhzwwqddjzdjdldlggvnvlnlbljlqjjmcczbccznnlnslswlwplpttvrvllfwflftllhclldhlhddbvvpvzpvpmmrccvgvdvqqjcqqwvvnjjlbbjwjrjhhlzhlhttljlcjjsnsgngrnntzzbsbmbsbrbdrdppjrjlrrjljqlqhqqnqsgdvhpgdhmnslqtjclmcfzrmgmlfnjbzznfgfprvwprwdbcgfcclmspgnzpbshwjbqvhzhrhswjzbfvnmcjtfvqbwmjpvfvctpmwsspdbtvfhfdfzjdpqnvslgmdvrnflzwzcnzmvzsvznwhpwtjwnqdgrrttmmdwzbbnwtllpbffrgtpjjjwltqrcbqcttdwnfjpmhdsbbpqmstjqchgjvfrmrbgqrlstnbdnzzzbzbsmsnnsssswmqhcbswtjhmcgnwmcclhzjqjzqcpbzgdzjgqzpqbmvvhtcznfrhdndswfvfhtfpdpszpjqrlwfdscvcngftwqmfttjtjrlbgcwvcjwsstqmcblmjzsgtgrqnqqvhzhvsphjmbcpfcznlcqldcvhlsvggbjngmhspwwqhlwstslvwmmbwqdmrgdvvnlstmjllhzscrhzjtmnsjfbndnlmzqbzgdgbcqchnbvwsftjtznnbsnvsgzpdzdqznjsslrlfnccdhwsljhczggvmgqswjltmrqqmwtbzmtdzhpjcvmwsscsdzpfnwlcrrdgzqqdmgwdlzvvvjcqsgpcwvrdnrstpcmgfjnjffbfmgzjthhllzrlsjtnqfppltbrlnqnjvqlvtpqvsbfgmmlcdzhgmzzqjwtqtzmpwwddbqrqnfzzpsjglsjddsslwwlrttzfzplmwsswlnvrvwwcgddjwcmvsjjbfgcfjmthfbpmcwjptchhnsmzttjqnwzdljffghhqdcwzwgbvfsmwqdbtblphdgcmbhprtbccjbzqrpvjdbnsmlwfntvjgptnshzmddwbhgwsnfrjbpqqwlsfdpnmmnnwhdmhzvjcmddbdnjzfzvffbgdqgwbggprcrbzwhvtzzgbhhcscrlmfgztfswjbsnwsmdfwlntwjzvlwhvlrfzszllmflmrsrcfnncvszvgdmmnvgrqnjhljcnrrhpdhffwmrsqfnbcpfdmmgppwjbjrwdfmpcrbznrjnbmssszhnlbpgmlczhhcdtgjqcbqrvzcbpgrftfhzdqthhspwnqqswntlpcmmqtcszngpggqvfjnmnprhdfjsngwrncjcmqdmjhpdlfnshpdlnlfpcnprwjgdvwwbvhvsbrfjtqsqjnvcpfdsrnfwmrrbtcvcqzflhdlbpcthzthdjzsrvwgbhjvhbtrngthfrszlvrbtnscsqlblcwlngslspcrhqzzdlzcdbhqdhthlpmdrntbhnqtwtzwpndbgphpsllbvgqjtmszdvjpgttzcmbwgrgdwmsbfgvgbbcmsnhvmnsbcsthsdwdqtghpdclfbglbdjgnnnwhmzzvnhbmgfbmvqwvwqhdswgtzslspmbmznnwdmjbzbddhzchtdzdzgwtlmlpmwrqvghpfwhvfjrtvmjwgjjnwdwnpdcqjdmcctjfcrdgpvczvnhlrbfmqgnrhmdwsrmmpqhvwgqgbqccpznpjfldwpntnvzgdfzljmtqwvfnrdsjsqgbvzjsczwwjggqtrpvwgqggwwhqggtgqfjmzsmjvdhdwqggbgnftpqqqlsfpflwrdpjwnhfdpchcgntjshgtwnwrnpsvwmplvqcltbgrcpflpgzbqfclghfnwjchnbgjnplgldmphdplvjrnrtzcmlftprsnmrjmnffpqjlvqlztbwprjwprrmmgzhjgdnhbfdrwjtvsvnbqhtfhbqgdrvcwlwfdbbcthgvttpvrwrqmpmrmvgpjzwlpvbqcvgccpgfddjbwhrvgqmjqzwgghllrtrblcpbttmcrgjsftlqhjfvqnbhmhbhngwnfqtgdttstzvmstrqcpfjrdgtsdbqqccqvbhpwhnpmpqgntfqszndjrmfhlqjqjbqvjtlfmrnnzzrtqlzzhjfqmmsmzvzrcplmfjpcmpfmpzbsbrmbnbnjqwjcfwnnwwrwzvvrsvhvrwnhlmwqjdztqthcwnwrlbjdflfsbplbwfmzqqnpwvzjbcfdgztpwttlrvhlfzzsfltpqwcpnzlsqgvwnqfvgclrfvssfcfmfvvjsndrhqdbrqfggfhjbvdvvmgpglqzwgjdmqtscjpfhgsbshghtmftrrhznttpzrzcsmrrvzdjmwtmbcbpqbsdmqzqdzrncwzmptltvdphsltfrhbbrdzbnbsqdhfvrgvmbgfvwblsjvfphlpzfvsllwnqjmbhngzzslcdmdzfgrgscbzggrzmbmwlzbnpzcvsbsfgdpnwzljsf"

	fmt.Println(GetMarker(datastream, 4))
	fmt.Println(GetMarker(datastream, 14))
}

func GetMarker(datastream string, size int) int {
	markerSet := make([]string, 0)
	for i, elem := range datastream {
		stringElem := string(elem)

		if slices.Contains(markerSet, stringElem) {
			idx := slices.Index(markerSet, stringElem)
			markerSet = append(markerSet, stringElem)
			markerSet = markerSet[idx+1:]
		} else {
			markerSet = append(markerSet, stringElem)
		}

		if len(markerSet) == size {
			return i + 1
		}
	}

	return 0
}
