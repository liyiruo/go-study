# go-study
学习go语言了

## go语言中的关键字

var和const ：变量和常量的声明
var varName type  或者 varName : = value
package and import: 导入
func： 用于定义函数和方法
return ：用于从函数返回
defer someCode ：在函数退出之前执行
go : 用于并行
select 用于选择不同类型的通讯
interface 用于定义接口
struct 用于定义抽象数据类型
break、case、continue、for、fallthrough、else、if、switch、goto、default 流程控制
chan用于channel通讯
type用于声明自定义类型
map用于声明map类型数据
range用于读取slice、map、channel数据


## Go 语言范围(Range)
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
## Go 语言Map(集合)
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。