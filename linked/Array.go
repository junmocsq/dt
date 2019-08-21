package linked

type array struct {
	Arr    []interface{}
	Size   uint32
	Length uint32
}

// 创建数组
func ACreate(length uint32) array {
	return array{Arr: make([]interface{}, length), Size: 0, Length: length}
}

func (arr *array) At(i uint32) interface{} {
	if i >= arr.Size {
		return nil
	}
	return arr.Arr[i]
}

func (arr *array) Empty() bool {
	if arr.Size == 0 {
		return true
	} else {
		return false
	}
}

func (arr *array) End() interface{} {
	if arr.Empty() {
		return nil
	}
	return arr.Arr[arr.Size-1]
}

func (arr *array) Fill(val interface{}, num uint32) bool {
	if num > arr.Length {
		return false
	}
	var i uint32
	for i = 0; i < num; i++ {
		arr.Arr[i] = val
	}
	return true
}

func (arr *array) MaxSize() uint32 {
	return arr.Length
}

func (arr *array) Len() uint32 {
	return arr.Size
}

func (arr *array) Put(val interface{}, i uint32) bool {
	if arr.Size <= i {
		return false
	}
	arr.Arr[i] = val
	return true
}

func (arr *array) IsFull() bool {
	if arr.Size == arr.Length {
		return true
	} else {
		return false
	}
}

func (arr *array) Add(val interface{}) uint32 {
	if arr.IsFull() {
		return -1
	}
	arr.Arr[arr.Size] = val
	arr.Size++
	return arr.Size - 1
}
