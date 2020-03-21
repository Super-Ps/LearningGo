package main
import "fmt"

func consts(){
	// 特殊的常量枚举 
	//显示的声明  关键字iota  默认值为0，const中每增加一行加1,实现自增值
	const(
		a=iota  //0
		b=iota  //1
	)
	// 没遇到一个const 关键字 iota机安检组就会重置，此时V == 0
	const v = iota    //0

	
	const (
		c = iota  //0
		_     //1 被弃掉
		python  //2
		java  //3
		node  //4
	)

	const(
		d,e,f = iota,iota,iota  // d=0 ,e=0,f=0   iota在同一行值相同 
	)

	const(

		g=iota   //0
		h="H"  // H
		i =iota  //2
		j,k,l =iota,iota,iota  //3
		m =iota  //4
	)


	fmt.Println(a,b,v,c,python,java,node,d,e,f,g,h,i,j,k,l,m)

}


func main(){
	fmt.Println("main")
	consts()
}