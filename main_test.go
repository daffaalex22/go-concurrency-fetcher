package main

import (
	"fmt"
	"testing"
)

var testCases = [][]int{{5000, 200}}

func BenchmarkConcFetch(b *testing.B) {
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("Jobs-%d-Workers-%d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConcurrentFetch(testCase[0], testCase[1])
			}
		})
	}
}

func BenchmarkConcFetchWrite(b *testing.B) {
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("Jobs-%d-Workers-%d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConcurrentFetchWrite(testCase[0], testCase[1])
			}
		})
	}
}

func BenchmarkSeqFetch(b *testing.B) {
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("Jobs-%d-Workers-%d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SequentialFetch(testCase[0])
			}
		})
	}
}
