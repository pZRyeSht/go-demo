package safemap

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafeMap(t *testing.T) {
	smp := NewSafeMap()
	wg := sync.WaitGroup{}

	// 并发新增
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		wg.Add(1)
		go func(smi SafeMapImp, k, v string) {
			smi.Insert(k, v)
			wg.Done()
		}(smp, key, key)
	}

	// 并发查找
	go func() {
		for i := 0; i < 20; i += 2 {
			key := fmt.Sprintf("key%d", i)
			wg.Add(1)
			go func(smi SafeMapImp, k string) {
				value, found := smi.Find(key)
				fmt.Printf("Find Key:%s,result:%s,%v\n", k, value, found)
				wg.Done()
			}(smp, key)
		}
	}()

	// 并发更新
	go func() {
		for i := 0; i < 30; i += 3 {
			key := fmt.Sprintf("key%d", i)
			wg.Add(1)
			go func(smi SafeMapImp, k string) {
				smi.Update(key, func(found bool) (newVal interface{}) {
					if found {
						newVal = "update"
						return
					}
					return nil
				})
				wg.Done()
			}(smp, key)
		}
	}()

	// 并发删除
	go func() {
		for i := 0; i < 40; i += 4 {
			key := fmt.Sprintf("key%d", i)
			wg.Add(1)
			go func(smi SafeMapImp, k string) {
				smi.Delete(k)
				wg.Done()
			}(smp, key)
		}
	}()

	// 并发计算长度
	wg.Add(1)
	go func(smi SafeMapImp) {
		l := smi.Len()
		fmt.Println("SafeMap Len:", l)
		wg.Done()
	}(smp)

	wg.Wait()

	// 关闭管到输出结果
	mp := smp.Close()
	fmt.Println("Print SafeMap:")
	for k, v := range mp {
		fmt.Printf("key:%s,value:%v\n", k, v)

	}
}
