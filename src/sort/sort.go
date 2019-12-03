package sort
import(
	"fmt"
	"math/rand"
	"time"
)

type Interface interface{
	Len() int
	Less(i,j int) bool
	Swap(i,j int)
	
	Get(i int) interface{}
	Set(i int,v interface{})
	Equal(i int,v interface{}) bool
	LessValue(i int,v interface{}) bool
	Copy() Interface
}

func IsSorted(origin Interface) bool{
	for i := 0; i < origin.Len() - 1; i++ {
		if(origin.Less(i+1,i)){
			return false
		}
	}
	return true
}
func ShowList(origin Interface) {
	fmt.Println(origin.Len(),origin.Get(0),origin.Get(origin.Len()-1));
}
func ShowFullList(origin Interface) {
	fmt.Println(origin);
}

//Selection Sort
func Selection(origin Interface){
	l := origin.Len()
	for i := 0; i < l-1; i++ {
		min := i
		for j := min + 1; j < l; j++ {
			if(origin.Less(j,min)){
				origin.Swap(j,min)
			}
		}
		if(min != i){
			origin.Swap(i,min)
		}
	}
}
//Insert Sort
func Insert(origin Interface){
	l := origin.Len()
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if(origin.Less(j,j-1)){
				origin.Swap(j,j-1);
			}else{
				break;
			}
		}
	}
}
//Puple Sort
func Pupple(origin Interface){
	l := origin.Len()
	for i := l-1; i > 0; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if(origin.Less(j+1,j)){
				origin.Swap(j,j+1);
				swap = true
			}
		}
		if(!swap){
			break
		}
	}
}
//shellSort, skip wide swaps
func Shell(origin Interface){
	l := origin.Len()
	h := 1
	for h < l {
		h = h * 3 + 1
	}

	for h > 0 {
		for i := h; i < l; i += h {
			for j := i; j > 0 && origin.Less(j,j-h); j -= h {
				//fmt.Println(i,j,h)
				origin.Swap(j,j-h)
			}
		}
		h = (h - 1) / 3
	}
}
//mergeSort
func bootUpMergeSort(origin Interface){
	l := origin.Len()
	target := origin.Copy()
	pt,po := &target,&origin

	for w := 2; w < 2*l; w *= 2 {
		lo,hi := 0,w
		mid := (hi - lo) / 2 + lo
		for  lo < l {
			if(hi > l) {
				hi = l
			}
			for k,i,j := lo,lo,mid; k < hi; k++ {
				if(j >= hi){
					(*pt).Set(k,(*po).Get(i))
					i++
				}else if(i >= mid || (*po).Less(j, i)){
					(*pt).Set(k,(*po).Get(j))
					j++
				}else{
					(*pt).Set(k,(*po).Get(i))
					i++
				}
			}
			lo,hi = lo + w, hi + w
			mid = (hi - lo) / 2 + lo
		}
		pt,po = po,pt
	}
	for i := 0; i < l; i++ {
		origin.Set(i,(*po).Get(i))
	}
}
func merge(origin Interface,target Interface,lo int,hi int,mid int){

	i,j,k := lo,mid,lo
	for k < hi {
		if(j >= hi){
			target.Set(k,origin.Get(i))
			i++
		}else if(i >= mid || origin.Less(j,i)){
			target.Set(k,origin.Get(j))
			j++
		}else{
			target.Set(k,origin.Get(i))
			i++
		}
		k++
	}

}
func subMSort(origin Interface,target Interface,lo int,hi int){
	if(hi - lo <= 1){
		return
	}
	mid := (hi - lo) / 2 + lo;
	subMSort(target,origin,lo,mid)
	subMSort(target,origin,mid,hi)
	merge(origin,target,lo,hi,mid)
}
func mSort(origin Interface){
	copyArray := origin.Copy()
	l := origin.Len()
	subMSort(copyArray,origin,0,l)
}
//heapSort
func swim(origin Interface,k int){
	if(k == 0) {
		return
	}
	pk := (k - 1) /2
	if(origin.Less(k,pk)){
		origin.Swap(k,pk)
		swim(origin,pk)
	}
}

func sink(origin Interface,k int, l int){

	choose,right := 2 * k + 1, 2* k +2
	if(l > right && origin.Less(right,choose)){
		choose = right
	}
	if(l > choose && origin.Less(choose,k)){
		origin.Swap(choose,k)
		sink(origin,choose,l)
	}
}

func heapSort(origin Interface){
	l := origin.Len()
	//build heap
	m := (l - 1 ) / 2
	for i := m; i >= 0; i-- {
		sink(origin,i,l)
	}

	//sort
	for t := l - 1; t >= 0; t-- {
		origin.Swap(0,t)
		sink(origin,0,t)
	}

	for i,j :=0, l-1; i < j; i,j = i+1,j-1 {
		origin.Swap(i,j)
	}
}

func partition(origin Interface,lo int, hi int) int{
	i,j := lo + 1,hi
	for true {
		//fmt.Println(lo,hi,i,j)
		for origin.Less(i,lo) {
			if(i == hi) {
				break
			}
			i++
		}

		for origin.Less(lo,j) {
			j--
		}
		if(i >= j){
			break
		}
		origin.Swap(i,j)
		i++
		j--
	}
	origin.Swap(lo,j)
	return j
}

func qSort(origin Interface, lo int, hi int){
	if(hi - lo < 1){
		return
	}
	j := partition(origin,lo,hi)
	qSort(origin,lo,j-1)
	qSort(origin,j+1,hi)

}

func quickSort(origin Interface){
	l := origin.Len()
	qSort(origin,0,l-1)
}


func shuffle(origin Interface){
	l := origin.Len()
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	for i := 0; i < l; i++ {
		j :=  i + r1.Intn(100 * l) / ((100 * l)/( l - i ) + 1)
		origin.Swap(i,j)
	}
}


func bSearch(origin Interface,lo int, hi int,v interface{}) int{
	if(hi < lo){
		return -1
	}
	mid := (hi - lo) / 2 + lo
	if(origin.LessValue(mid,v) || origin.Equal(mid,v)){
		f := bSearch(origin,mid + 1,hi,v)
		if(f != -1){
			return f
		}else if(origin.Equal(mid,v)){
			return mid
		}
	}else {
		return bSearch(origin,lo,mid-1,v)
	}
	return -1
}

func bottomUpSearch(origin Interface,lo int, hi int,v interface{}) int{
	if(hi < lo){
		return -1
	}
	mid := (hi - lo) / 2 + lo
	for lo <= hi {
		mid = (hi - lo + 1) / 2 + lo

		//right one
		if(origin.LessValue(mid,v) || origin.Equal(mid,v)){
			if(hi == mid){
				break
			}
			lo = mid
		}else{
			hi = mid-1
		}

		// //left one
		// if(origin.LessValue(mid,v)){
		// 	lo = mid + 1
		// }else{
		// 	if(hi == mid){
		// 		break
		// 	}
		// 	hi = mid
		// }
	}
	return mid
}


func Search(origin Interface,v interface{}) int{
	if(!IsSorted(origin)){
		panic("origin array is not sorted")
	}
	l := origin.Len()
	return bottomUpSearch(origin,0,l-1,v)
}


