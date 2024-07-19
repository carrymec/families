# 创建节点
```shell
create(node-name:label-name {column1: value1, column2: value2})
# eg
create(users:person {id:123, name: "hello", age: 22})
```
# 查询并返回
```shell
match(node-name:label-name) return node-name #(node-name 这是我们要创建一个节点名称。)
# eg
match(u:person) return u
# 可以返回指定属性，多个属性用逗号隔开，可以使用 as 关键字重命名
match(u:person) return u.属性名,u.属性名1
```
# 条件查询并返回
```shell
match(u:person) where u.name = "hello" return u
```
# 现有节点创建有属性、无属性的关系
```shell
# 这里创建两个节点 用户和信用卡
create(cc:CreditCard{id:"5001",number:"123456789",ccv:"888",expiredate:"20/17"})
create(e:Customer{id:"1001",name:"abc",dob:"01/10/2001"})

# 为这个用户创建一条信用卡的关系语法
MATCH (<node1-label-name>:<node1-name>),(<node2-label-name>:<node2-name>)
CREATE  
	(<node1-label-name>)-[<relationship-label-name>:<relationship-name>
	{<define-properties-list>}]->(<node2-label-name>)
RETURN <relationship-label-name>
# eg 创建一条关系名为do_shopping_with，关系标签为r的关系，shopdate、price是这个关系的属性
match(cust:Customer),match(cc:CreditCard) create (cust) -[r:do_shopping_with{shopdate:"12/12/2020",price: 15000}] ->(cc) return r
# 创建一条没有属性的关系语法
MATCH (<node1-label-name>:<nodel-name>),(<node2-label-name>:<node2-name>)
CREATE  
	(<node1-label-name>)-[<relationship-label-name>:<relationship-name>{<define-properties-list>}]->(<node2-label-name>)
RETURN <relationship-label-name>
# eg
MATCH (e:Customer),(cc:CreditCard) 
CREATE (e)-[r:DO_SHOPPING_WITH ]->(cc) 
```
# 新节点无属性、有属性的关系
```shell
# 无属性语法
CREATE  
   (<node1-label-name>:<node1-name>)-
   [<relationship-label-name>:<relationship-name>]->
   (<node1-label-name>:<node1-name>)
RETURN <relationship-label-name>
# eg
CREATE (fb1:FaceBookProfile1)-[like:LIKES]->(fb2:FaceBookProfile2) 
# 查询
MATCH p=()-[r:LIKES]->() RETURN p 
# 有属性语法
CREATE  
	(<node1-label-name>:<node1-name>{<define-properties-list>})-
	[<relationship-label-name>:<relationship-name>{<define-properties-list>}]
	->(<node1-label-name>:<node1-name>{<define-properties-list>})
RETURN <relationship-label-name>
# eg
CREATE (video1:YoutubeVideo1{title:"Action Movie1",updated_by:"Abc",uploaded_date:"10/10/2010"})
-[movie:ACTION_MOVIES{rating:1}]->
(video2:YoutubeVideo2{title:"Action Movie2",updated_by:"Xyz",uploaded_date:"12/12/2012"}) 
```
# 查询关系
```shell
# 语法
MATCH 
(<node1-label-name>)-[<relationship-label-name>:<relationship-name>]->(<node2-label-name>)
RETURN <relationship-label-name>
# eg
MATCH (cust)-[r:do_shopping_with]->(cc) RETURN cust,cc
```
# 删除node、删除关系
```shell
match(c:FaceBookProfile1) -[r:ACTION_MOVIES] -> (cc:FaceBookProfile2) delete c,cc,r
MATCH (n:Employee) delete n 
```
# TODO https://www.w3cschool.cn/neo4j/neo4j_cql_remove.html 
```shell

```