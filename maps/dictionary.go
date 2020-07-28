package maps

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

/**
map:
引用类型，可以不用传指针直接修改对应值。
声明后，需要初始化才可以使用，否则会nil指针异常。例如：var m1 map[string]string m1 := make(map[string]string, 10)
对一个已经存在的key进行赋值，会覆盖原有值。
删除一个不存在的key，没有影响。

错误处理：
可以通过定义全局的错误变量，对外暴露固定的错误类型。
可以自定义错误类型，实现Error()方法。
*/

// 对错误提取公共变量
//var ErrNotFound = errors.New("could not find the word you were looking for")
//var ErrWordExisted = errors.New("word had existed")

// 自定义错误类型 实现Error()方法
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExisted      = DictionaryErr("word had existed")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	//return d[word], nil

	// return ok 判断该key是否存在于map中
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// map是引用类型，无需传指针，也可以改变其中的值。map底层是*hmap，其实是一个指针类型
// 无论map有多大，都只会有一个副本
// map作为引用类型，应该初始化后才可以使用，否则会nil指针异常。例如：make()初始化。
func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}

func (d Dictionary) Add2(word string, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExisted
	}
	d.Add(word, definition)

	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		//return err
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// 删除一个不存在的键时是没有问题的
	delete(d, word)
}
