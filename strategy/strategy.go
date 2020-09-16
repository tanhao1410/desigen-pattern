package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 6, 4}

	//如数据排序策略有冒泡排序、选择排序、插入排序、二叉树排序等
	//策略上下文，客户端主要通过该环境类来实现不同的策略调整，而不需要改变代码，也不用加入过多的if else 判断
	sorter := new(QuickSort)
	//sorter := new(BubleSort)

	//设置不同的策略，则会用不同的策略方法，对客户端的调用来说，写法一样，因为有一个统一的接口，本例为ISort接口。
	context := new(StrategyContext)
	context.SetSort(sorter)
	//客户端通过统一的接口来进行引用
	sort := context.GetSort()
	sort.Sort(arr)

	//工厂策略模式，从工厂获取对应的策略方法接口实例
	factory := NewStrategyFactory()
	if sort2, ok := factory.GetSortStrategy("buble"); ok {
		sort2.Sort(arr)
	}

}

type ISort interface {
	Sort(arr []int)
}

type QuickSort struct {
}

type BubleSort struct {
}

func (this *BubleSort) Sort(arr []int) {
	//冒泡排序实现
	fmt.Println("通过冒泡来实现排序")
}
func (this *QuickSort) Sort(arr []int) {
	//快速排序实现
	fmt.Println("通过快排来实现排序")
}

//策略上下文
type StrategyContext struct {
	Sorter ISort
}

func (factory *StrategyContext) GetSort() ISort {
	return factory.Sorter
}

func (factory *StrategyContext) SetSort(sorter ISort) {
	factory.Sorter = sorter
}

//策略工厂
type StrategyFactory struct {
	strategys map[string]ISort
}

func NewStrategyFactory() *StrategyFactory {
	factory := new(StrategyFactory)
	//初始化 内部的策略
	var strategys map[string]ISort = make(map[string]ISort, 2)
	quickSort := new(QuickSort)
	bubbleSort := new(BubleSort)
	strategys["quick"] = quickSort
	strategys["buble"] = bubbleSort
	factory.strategys = strategys
	return factory
}

//策略工厂提供该方法，客户端通过此 来获取不同的策略。
func (factory *StrategyFactory) GetSortStrategy(name string) (ISort ISort, ok bool) {
	if v, ok := factory.strategys[name]; ok {
		return v, true
	}
	return nil, false
}
