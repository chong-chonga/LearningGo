package collections

import "fmt"

func TestMap() {
	// 声明一个未初始化的 map, 是不能直接使用的
	var imap map[int]int
	fmt.Println("imap = ", imap)
	fmt.Printf("imap type is %T\n", imap)
	if imap == nil {
		fmt.Println("current imap is nil!")
	}

	imap = make(map[int]int)
	imap[1] = 1
	imap[2] = 1
	for entry := range imap {
		fmt.Println(entry, ":", imap[entry])
	}

	// 使用 len 函数获取 map 的大小
	size := len(imap)
	fmt.Println("map size is ", size)

	// 可以通过下面的方式判断 map 的 key 是否存在
	key := 3
	val, exist := imap[key]
	if exist {
		fmt.Println("val is ", val)
	} else {
		// key 不存在时, 返回的是该类型的默认值, map 不会报错
		// 因此不能通过 val 的值判断 key 是否存在
		fmt.Println("key: ", key, " not exist, the corresponding val is ", val)
	}

	// 可以通过 delete 函数来删除 map 中的 key
	key = 1
	_, exist = imap[key]
	if exist {
		fmt.Println("key: ", key, " exists; now, delete the key")
		// delete 韩叔没有返回值
		delete(imap, key)
		_, exist = imap[key]
		if !exist {
			fmt.Println("key: ", key, " not exists, delete success!")
		} else {
			fmt.Println("delete failed")
		}
	} else {
		fmt.Println("key: ", key, " not exists")
	}

	// map 是引用类型, 赋给新变量是赋引用的
	imap2 := imap
	imap2[1] = 2
	for entry := range imap {
		fmt.Println(entry, ":", imap[entry])
	}
	fmt.Println("map is reference type, any modification on the map will reflect on the original map")
}
