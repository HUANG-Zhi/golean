package sort
import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)
type IntSlice []int
func (p IntSlice) Len() int {return len(p)}
func (p IntSlice) Less(i,j int) bool {return p[i] < p[j]}
func (p IntSlice) Swap(i,j int) {p[i],p[j] = p[j],p[i]}

func (p IntSlice) Copy() Interface{
	t := IntSlice{}
	for i := 0; i < len(p); i++ {
		t = append(t,p[i])
	}
	return t
}
func (p IntSlice) Get(i int) interface{} {return p[i]}
func (p IntSlice) Set(i int,v interface{}) {p[i] = v.(int)}
func (p IntSlice) Equal(i int,v interface{}) bool {return p[i] == v.(int)}
func (p IntSlice) LessValue(i int,v interface{}) bool {return p[i] < v.(int)}

// func TestExchange(t *testing.T){
// 	aIntOrigin := IntSlice {3,2,1}
// 	if(aIntOrigin[0] !=3){
// 		t.Error("should be 3 before echange")
// 	}
// 	aIntOrigin.Swap(0,2)
// 	if(aIntOrigin[0] !=1){
// 		t.Error("should be 1 after exchange")
// 	}
// }

// func TestIsAsc(t *testing.T){
// 	aIntOrigin := IntSlice {1,2,2}
// 	if(!IsSorted(aIntOrigin)){
// 		t.Error("not ordered")
// 	}
// }

func randIntArray(Max int) IntSlice{
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	size :=  r1.Intn(Max)
	if(size == 0){
		size = 1
	}
	a := make([] int,size)
	for i := 0; i < size; i++ {
		a[i] = r1.Intn(Max)
	}
	return a
}

func bench(f func(p Interface),o Interface,t *testing.T,name string){
	//fmt.Println(name)
	count := 20
	for i := 0; i < count; i++ {
		// ShowList(o)
		shuffle(o)
		f(o)
		if(!IsSorted(o)){
			t.Error("not ordered")
		}
	}
}	

var aIntOrigin IntSlice = randIntArray(100000)
var aSmallIntOrigin IntSlice = []int {17,4,9,4,18,2,18,9,9,16}

func KmpPartten(ts string) [][]int{
	ls := len(ts)
	ld := 256
	dic := make([][]int,ls)
	for i := range dic {
		dic[i] = make([]int,ld)
	}

	dic[0][ts[0]] = 1
	k := 0
	for i := 1; i < ls; i++ {
		for j := 0; j < 256; j++ {
			dic[i][j] = dic[k][j]
		}
		dic[i][ts[i]] = i + 1
		//"sssbasbbaasa" => "ssbasbbaasa"
		k = dic[k][ts[i]]
	}
	return dic
}

func KmpSearch(can string,  partten [][]int) int{
	l := len(partten)
	k := 0
	for i := 0; i < len(can); i++ {
		k = partten[k][can[i]]
		if(k == l){
			return i + 1 - l
		}
	}
	return -1
}
func TestKmp(t *testing.T){

	ts,can := "sa","sssbasabbaasatttttsv"
	i := KmpSearch(can,KmpPartten(ts))
	fmt.Println(i)
	if(i >= 0){
		fmt.Println(can[i:])
	}
}
// func TestSearch(t *testing.T){
// 	quickSort(aSmallIntOrigin)
// 	ShowFullList(aSmallIntOrigin)
// 	fmt.Println(Search(aSmallIntOrigin,18))
// }

// func TestShuffle(t *testing.T){
// 	ShowFullList(aSmallIntOrigin)
// 	shuffle(aSmallIntOrigin)
// 	ShowFullList(aSmallIntOrigin)
// }

// func TestShell(t *testing.T){
// 	bench(Shell,aIntOrigin,t,"Shell")
// }

// func TestQuickSort(t *testing.T){
	
// 	//aIntOrigin = [] int{6,0,5,0,7,15,16,19,8,0,19,6,2}
// 	// s := aIntOrigin.Copy()
// 	// s.Set(1,30)
// 	// ShowList(aIntOrigin)
// 	//ShowList(aIntOrigin)
// 	bench(quickSort,aIntOrigin,t,"quickSort")
// }

// func TestHeapSort(t *testing.T){
// 	bench(heapSort,aIntOrigin,t,"heapSort")
// }

// func TestMSort(t *testing.T){
// 	bench(mSort,aIntOrigin,t,"mSort")
// }

// func TestBootUpMergeSort(t *testing.T){
// 	bench(bootUpMergeSort,aIntOrigin,t,"bootUpMergeSort")
// }


// func TestPupple(t *testing.T){
// 	bench(Pupple,aIntOrigin,t,"Pupple")
// }

// func TestInsert(t *testing.T){
// 	bench(Insert,aIntOrigin,t,"Insert")
// }

// func TestSelection(t *testing.T){
// 	bench(Selection,aIntOrigin,t,"Selection")
// }


// === RUN   TestShell
// --- PASS: TestShell (27.05s)
// === RUN   TestQuickSort
// --- PASS: TestQuickSort (0.12s)
// === RUN   TestHeapSort
// --- PASS: TestHeapSort (0.16s)
// === RUN   TestMSort
// --- PASS: TestMSort (0.33s)
// === RUN   TestBootUpMergeSort
// --- PASS: TestBootUpMergeSort (0.36s)
// === RUN   TestPupple
// --- PASS: TestPupple (71.31s)
// === RUN   TestInsert
// --- PASS: TestInsert (29.29s)
// === RUN   TestSelection
// --- PASS: TestSelection (70.81s)



// === RUN   TestShell
// --- PASS: TestShell (45.37s)
// === RUN   TestQuickSort
// --- PASS: TestQuickSort (0.15s)
// === RUN   TestHeapSort
// --- PASS: TestHeapSort (0.20s)
// === RUN   TestMSort
// --- PASS: TestMSort (0.40s)
// === RUN   TestBootUpMergeSort
// --- PASS: TestBootUpMergeSort (0.44s)
// === RUN   TestPupple
// --- PASS: TestPupple (119.45s)
// === RUN   TestInsert
// --- PASS: TestInsert (53.17s)
// === RUN   TestSelection
// --- PASS: TestSelection (131.79s)
