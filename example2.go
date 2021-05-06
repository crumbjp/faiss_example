package main

import (
	"fmt"
	"github.com/DataIntelligenceCrew/go-faiss"
)

func AddWithIDsDontWorkOnFlatIndex() {
	fmt.Println("*** AddWithIDsDontWorkOnFlatIndex ")
	index, _ := faiss.IndexFactory(2, "Flat", faiss.MetricL2)
	index.AddWithIDs([]float32{0,0,1,0,0,1,1,1}, []int64{0,1,2,3})
}
// *** AddWithIDsDontWorkOnFlatIndex
// Error in virtual void faiss::Index::add_with_ids(faiss::Index::idx_t, const float *, const faiss::Index::idx_t *) at /Users/crumbjp/git/faiss/faiss/Index.cpp:39: add_with_ids not implemented for this type of index

func RemoveIDsOnFlatIndex() {
	fmt.Println("*** RemoveIDsOnFlatIndex")
	index, _ := faiss.IndexFactory(2, "Flat", faiss.MetricL2)
	index.Add([]float32{0,0,1,0,0,1,1,1})
	fmt.Printf("Ntotal: %v\n", index.Ntotal())
	distances, labels, _ := index.Search([]float32{1, 0.1}, 4)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
	selector, _ := faiss.NewIDSelectorBatch([]int64{0})
	defer selector.Delete()
	index.RemoveIDs(selector)
	fmt.Printf("Ntotal: %v\n", index.Ntotal())
	distances, labels, _ = index.Search([]float32{1, 0.1}, 4)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
}
// *** RemoveIDsOnFlatIndex
// Ntotal: 4
// Distances: [0.010000001 0.80999994 1.01 1.81], Labels: [1 3 0 2]
// Ntotal: 3
// Distances: [0.010000001 0.80999994 1.81 3.4028235e+38], Labels: [0 2 1 -1]

func RemoveIDsOnHNSW() {
	fmt.Println("*** RemoveIDsOnHNSW ")
	index, _ := faiss.IndexFactory(2, "HNSW32", faiss.MetricL2)
	index.Add([]float32{0,0,1,0,0,1,1,1})
	fmt.Printf("Ntotal: %v\n", index.Ntotal())
	selector, _ := faiss.NewIDSelectorBatch([]int64{0})
	defer selector.Delete()
	index.RemoveIDs(selector)
}
// *** RemoveIDsOnHNSW
// Ntotal: 4
// Error in virtual size_t faiss::Index::remove_ids(const faiss::IDSelector &) at /Users/crumbjp/git/faiss/faiss/Index.cpp:43: remove_ids not implemented for this type of index

func IVFRequiresTrain() {
	fmt.Println("*** IVFRequiresTrain")
	index, _ := faiss.IndexFactory(2, "IVF4,Flat", faiss.MetricL2)
	index.AddWithIDs([]float32{0,0,1,0,0,1,1,1}, []int64{0,1,2,3})
}
// *** IVFRequiresTrain
// Error in virtual void faiss::IndexIVFFlat::add_core(faiss::Index::idx_t, const float *, const int64_t *, const int64_t *) at /Users/crumbjp/git/faiss/faiss/IndexIVFFlat.cpp:44: Error: 'is_trained' failed

func IVFWithNprobe() {
	fmt.Println("*** IVFWithNprobe")
	index, _ := faiss.IndexFactory(2, "IVF4,Flat", faiss.MetricL2)
	index.Train([]float32{0,0,1,0,0,1,1,1})
	index.AddWithIDs([]float32{0,0,1,0,0,1,1,1}, []int64{0,1,2,3})
	fmt.Printf("Ntotal: %v\n", index.Ntotal())
	distances, labels, _ := index.Search([]float32{1, 0.1}, 4)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
	parameterSpace, _ := faiss. NewParameterSpace()
	parameterSpace.SetIndexParameter(index, "nprobe", 2)
	distances, labels, _ = index.Search([]float32{1, 0.1}, 4)
	fmt.Printf("Distances: %v, Labels: %v\n", distances, labels)
}
// *** IVFWithNprobe
// WARNING clustering 4 points to 4 centroids: please provide at least 156 training points
// Ntotal: 4
// Distances: [0.010000001 3.4028235e+38 3.4028235e+38 3.4028235e+38], Labels: [1 -1 -1 -1]
// Distances: [0.010000001 0.80999994 3.4028235e+38 3.4028235e+38], Labels: [1 3 -1 -1]

func main() {
	AddWithIDsDontWorkOnFlatIndex()
	RemoveIDsOnFlatIndex()
	RemoveIDsOnHNSW()
	IVFRequiresTrain()
	IVFWithNprobe()
}
