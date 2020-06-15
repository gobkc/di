package main

import (
	"fmt"
	"reflect"
)

type BaseI interface {

}
type Base struct {

}


type A struct {
	Age int
}

type B struct {
	Name string
}

type Container struct {
	Deps interface{}
}

func (i *Container)Inject(deps ...interface{})  {
	i.Deps = deps
}

func (i *Container)Invoke(deb interface{})  {
	a:=i.Deps
	getType := reflect.TypeOf(deb)
	if getType.Kind() == reflect.Ptr {
		getType = getType.Elem()
	}
	getVal := reflect.ValueOf(deb)
	switch getType.Kind() {
	case reflect.Func:

	}
	realA := a.([]interface{})
	//realParams := make([]reflect.Value, len(realA)) //参数
	realParams := make([]reflect.Value,1) //参数
	for _, item := range realA {
		t1 := reflect.TypeOf(item).Elem().Name()
		// todo 需要获取 deb函数的参数，当前传入的 deb的名称
		if t1=="B"{
			realParams[0] = reflect.ValueOf(item)
		}
	}

	getVal.Call(realParams)
}

func ref(res interface{}){
	getType := reflect.TypeOf(res)
	getValue := reflect.ValueOf(res)

	// getElem := res.Type().Elem()
	v := reflect.New(getType.Elem())
	fmt.Println(getType,getType.Name(),getType.Kind())
	fmt.Println(v)
	fmt.Println("getValue",getValue)
	a:=getType.Kind()

	fmt.Println(a)
	switch getType.Kind() {
	case  reflect.Slice:
		fmt.Println("Kind是：", getType.Kind())

		sturct_type := getType.Elem()
		fmt.Println("获取到slice内sturct类型:",sturct_type)
		fmt.Println("判断是什么类型的指针:",getType.Elem().Kind())
		var v reflect.Value
		if getType.Elem().Kind() == reflect.Ptr {
			subgetType := getType.Elem()
			sturct_type = subgetType.Elem()
			v = reflect.New(sturct_type)
			fmt.Println("生在对象信息1:",v,v.Elem(),sturct_type)
		}else{
			//根据类型生成&{},通过.Elem()再转为{}
			v = reflect.New(sturct_type).Elem()
			fmt.Println("生在对象信息2:",v,sturct_type.NumField())
		}
		//根据类型生成slice
		sl := reflect.MakeSlice(getType,0,0)
		sl = reflect.Append(sl,v)
		fmt.Println("反射获取slice :" , sl.Interface())
		// reflect.NewAt(ptr,sturct_type)
		//根据获取到的v可以进行动态赋值
		// reflect.NewAt(ptr,sturct_type)
	default:
		fmt.Println("err  ",getType)
	}

}
func App(b *B) {
	fmt.Println(b.Name)
}

func main() {
	var container = Container{}
	//1.注入依赖
	container.Inject(&A{
		Age:18,
	},&B{
		Name:"张三",
	})

	//2.调用依赖
	container.Invoke(App)

}