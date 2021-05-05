package main

import (
	"fmt"
	"github.com/DataIntelligenceCrew/go-faiss"
)


func Sample1() {
	fmt.Println("*** Sample1 Flat L2")
	index, _ := faiss.IndexFactory(2, "Flat", faiss.MetricL2)
	index.Add([]float32{0,0,1,0,0,1,1,1})
	fmt.Printf("Ntotal: %v\n", index.Ntotal())
	distances, labels, _ := index.Search([]float32{1, 0.1}, 4)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
}
// *** Sample1 Flat L2
// Ntotal: 4
// Distances: [0.010000001 0.80999994 1.01 1.81], Labels: [1 3 0 2]

func Sample2() {
	fmt.Println("*** Sample2 Flat IP")
	index, _ := faiss.IndexFactory(2, "Flat", faiss.MetricInnerProduct)
	index.Add([]float32{0,0,1,0,0,1,1,1})
	distances, labels, _ := index.Search([]float32{1, 0.1}, 4)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
}
// *** Sample2 Flat IP
// Distances: [1.1 1 0.1 0], Labels: [3 1 2 0]

func Sample3() {
	fmt.Println("*** Sample3 L2norm Flat IP")
	index, _ := faiss.IndexFactory(2, "L2norm, Flat", faiss.MetricInnerProduct)
	index.Add([]float32{0,0,1,0,0,1,1,1})
	distances, labels, _ := index.Search([]float32{1, 0.1}, 4)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
}
// *** Sample3 L2norm Flat IP
// Distances: [0.99503714 0.77395725 0.09950372 0], Labels: [1 3 2 0]

func RFlatDontWorkWithTransform() {
	fmt.Println("*** RFlatDontWorkWithTransform")
	index, _ := faiss.IndexFactory(2, "L2norm,Flat,RFlat", faiss.MetricInnerProduct)
	index.Add([]float32{0,0,1,0,0,1,1,1})
	distances, labels, _ := index.Search([]float32{1, 0.1}, 4)
	// RFlat returns simple IP, (Ignoring transform layer)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
}
// *** RFlatDontWorkWithTransform
// Distances: [1.1 1 0.1 0], Labels: [3 1 2 0]

func main() {
	Sample1()
	Sample2()
	Sample3()
	RFlatDontWorkWithTransform()
}
