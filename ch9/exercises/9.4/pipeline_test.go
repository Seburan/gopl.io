package pipeline

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {

		in, out, err := pipeline(i)
		if err != nil {
			b.Errorf("%v", err)
		}
		// b.ResetTimer()

		in <- 1
		val := <-out
		if val != 1 {
			b.Errorf("expected 1, got %v", val)
		}
		close(in)
		close(out)
	}
	// 	for i := 0; i < b.N; i++ {
	// 	echo2(args[:]);
	// }

}

func TestPipeline(t *testing.T) {
	start := time.Now()
	fmt.Println(start)

	in, out, err := pipeline(5000000)

	if err != nil {
		t.Errorf("%v", err)
	}
	in <- 1
	val := <-out
	if val != 1 {
		t.Errorf("expected 1, got %v", val)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("time since start : %v", secs)
}
