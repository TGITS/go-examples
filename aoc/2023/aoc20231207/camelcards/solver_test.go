package camelcards

import "testing"

func TestComputeTypeOfHandPart1(t *testing.T) {
	cases := []struct {
		hand     string
		expected int
	}{
		{"AAAAA", FIVE_OF_A_KIND},
		{"KKKK9", FOUR_OF_A_KIND},
		{"QQQJJ", FULL_HOUSE},
		{"55532", THREE_OF_A_KIND},
		{"88664", TWO_PAIR},
		{"99753", ONE_PAIR},
		{"AKJ72", HIGH_CARD},
	}

	for _, testCase := range cases {
		actual := ComputeTypeOfHandPart1(testCase.hand)
		if actual != testCase.expected {
			t.Fatalf("hand %s expected %d, got %d", testCase.hand, testCase.expected, actual)
		}
	}
}

func TestComputeTypeOfHandPart2WithJokers(t *testing.T) {
	cases := []struct {
		hand     string
		expected int
	}{
		{"JJJJJ", FIVE_OF_A_KIND},
		{"QJJQ2", FOUR_OF_A_KIND},
		{"KTJJT", FOUR_OF_A_KIND},
		{"QQQJA", FOUR_OF_A_KIND},
		{"T55J5", FOUR_OF_A_KIND},
	}

	for _, testCase := range cases {
		actual := ComputeTypeOfHandPart2(testCase.hand)
		if actual != testCase.expected {
			t.Fatalf("hand %s expected %d, got %d", testCase.hand, testCase.expected, actual)
		}
	}
}

func TestSolvePart1(t *testing.T) {
	actual := SolvePart1("../data/input_test.txt")
	expected := 6440
	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	actual := SolvePart2("../data/input_test.txt")
	expected := 5905
	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
