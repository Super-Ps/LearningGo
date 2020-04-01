import(
	"fmt"
)
package models



type person struct{

	Name string   // 首字母大写 外包可以访问
	age int       // 首字母小写 外包不可以访问，
	slary float64
}


//  写一个工厂函数，类似于构造函数

func NewPerson (name string ) *person{
return &person{
	Name : name,
}

}

func (p *person) SetAge( age int) {

	if age > 0 && age < 150{
		p.age = age
	}else{
		fmt.println("年龄不再范围内")
	}
}

func (p *person) GetAge()int{
	return p.age
}