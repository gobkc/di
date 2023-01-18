package main

import (
	"fmt"
	"reflect"
)

//type BaseI interface {
//}
//type Base struct {
//}
//
//type A struct {
//	Age int
//}
//
//type B struct {
//	Name string
//}
//
//type Container struct {
//	Deps interface{}
//}
//
//func (i *Container) Inject(deps ...interface{}) {
//	i.Deps = deps
//}
//
//func (i *Container) Invoke(deb interface{}) {
//	a := i.Deps
//	getType := reflect.TypeOf(deb)
//	if getType.Kind() == reflect.Ptr {
//		getType = getType.Elem()
//	}
//	getVal := reflect.ValueOf(deb)
//	switch getType.Kind() {
//	case reflect.Func:
//
//	}
//	realA := a.([]interface{})
//	//realParams := make([]reflect.Value, len(realA)) //参数
//	realParams := make([]reflect.Value, 1) //参数
//	for _, item := range realA {
//		t1 := reflect.TypeOf(item).Elem().Name()
//		// todo 需要获取 deb函数的参数，当前传入的 deb的名称
//		if t1 == "B" {
//			realParams[0] = reflect.ValueOf(item)
//		}
//	}
//
//	getVal.Call(realParams)
//}
//
//func ref(res interface{}) {
//	getType := reflect.TypeOf(res)
//	getValue := reflect.ValueOf(res)
//
//	// getElem := res.Type().Elem()
//	v := reflect.New(getType.Elem())
//	fmt.Println(getType, getType.Name(), getType.Kind())
//	fmt.Println(v)
//	fmt.Println("getValue", getValue)
//	a := getType.Kind()
//
//	fmt.Println(a)
//	switch getType.Kind() {
//	case reflect.Slice:
//		fmt.Println("Kind是：", getType.Kind())
//
//		sturct_type := getType.Elem()
//		fmt.Println("获取到slice内sturct类型:", sturct_type)
//		fmt.Println("判断是什么类型的指针:", getType.Elem().Kind())
//		var v reflect.Value
//		if getType.Elem().Kind() == reflect.Ptr {
//			subgetType := getType.Elem()
//			sturct_type = subgetType.Elem()
//			v = reflect.New(sturct_type)
//			fmt.Println("生在对象信息1:", v, v.Elem(), sturct_type)
//		} else {
//			//根据类型生成&{},通过.Elem()再转为{}
//			v = reflect.New(sturct_type).Elem()
//			fmt.Println("生在对象信息2:", v, sturct_type.NumField())
//		}
//		//根据类型生成slice
//		sl := reflect.MakeSlice(getType, 0, 0)
//		sl = reflect.Append(sl, v)
//		fmt.Println("反射获取slice :", sl.Interface())
//		// reflect.NewAt(ptr,sturct_type)
//		//根据获取到的v可以进行动态赋值
//		// reflect.NewAt(ptr,sturct_type)
//	default:
//		fmt.Println("err  ", getType)
//	}
//}
//
//func App(b *B) {
//	fmt.Println(b.Name)
//}

type Config struct {
	DbName string
	Pass   string
	Addr   string
	Port   int
}

func NewConfig() (*Config, error) {
	return &Config{
		DbName: "testmysql",
		Pass:   "123456",
		Addr:   "localhost",
		Port:   3306,
	}, nil
}

type MysqlIns struct {
	Conn string //test
	Conf *Config
}

func NewMysql(c *Config) *MysqlIns {
	return &MysqlIns{Conf: c}
}

type container struct {
}

func new() *container {
	return &container{}
}

func (c *container) provide(dest any) error {
	kind := reflect.TypeOf(dest).Kind()
	if kind != reflect.Func {
		return fmt.Errorf("provide kind:%v need:%v", kind, reflect.Func)
	}
	destValue := reflect.ValueOf(dest)
	destType := reflect.TypeOf(dest)
	numIn := destType.NumIn()
	numOut := destType.NumOut()
	in := fmt.Sprintf("%s", destType.In(0))
	out := fmt.Sprintf("%s", destType.Out(0))
	fmt.Println(destType, destValue, numIn, numOut, in, out)
	return nil
}

func (c *container) invoke(dest any) error {
	kind := reflect.TypeOf(dest).Kind()
	if kind != reflect.Func {
		return fmt.Errorf("invoke kind:%v need:%v", kind, reflect.Func)
	}

	return nil
}

func main() {
	c := new()
	if err := c.provide(NewMysql); err != nil {
		fmt.Println(err)
	}
	if err := c.provide(NewConfig); err != nil {
		fmt.Println(err)
	}
	c.invoke(func(i *MysqlIns) {

	})
}
