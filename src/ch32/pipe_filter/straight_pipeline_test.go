package pipe_filter

import "testing"

func TestStraightPipeline(t *testing.T) {
	split_er := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p1", split_er, converter, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("expect 6, but got %d", ret)
	}

}
