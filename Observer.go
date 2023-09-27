package main

import "fmt"

type Teacher interface{
	AddGroup(group Group )
	NotifyAll(mes string)
}

type Group interface{
	ReactToMes(mes string)
}

type teacher struct{
	groupList []Group
}

type group struct{
	groupId string
}

func GetNewTeacher() teacher{
	return teacher{groupList: make([]Group, 0)}
}

func GetNewGroup(Id string) group{
	return group{groupId: Id}
}

func(t *teacher) AddGroup(g group){
	t.groupList = append(t.groupList, g)
}

func(g group) ReactToMes(mes string){
	fmt.Printf("group %s got an message %v \n ", g.groupId,mes)
}

func(t teacher) NotifyAll(mes string){
	for _,g := range t.groupList{
		fmt.Println("Teacher add assignment for group ", g.(group).groupId)
		g.ReactToMes(mes)
	}
}

func main(){
	teach1 := GetNewTeacher()
	group1 := GetNewGroup("SE-2201")
	group2 := GetNewGroup("SE-2202")
	group3 := GetNewGroup("SE-2203")
    teach1.AddGroup(group1)
	teach1.AddGroup(group2)
	teach1.AddGroup(group3)
	teach1.NotifyAll("I added assignment 1 for your group")
}



