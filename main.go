package main

import "fmt"

type Friend struct {
	id int
	owes int
	inGroup bool
	friends []int
}

func (friend *Friend) addFriend(i int) {
	friend.friends = append(friend.friends, i)
}

func divideInGroups(friendId int, friends *map[int]Friend, group *[]Friend) {
	friend := (*friends)[friendId]
	if friend.inGroup {
		return
	}
	*group = append((*group), friend)
	friend.inGroup = true
	(*friends)[friendId] = friend
	for i := 0; i < len(friend.friends); i++ {
		divideInGroups(friend.friends[i], friends, group)
	}
}

func CalculateSumForFriends(groups *[][]Friend) bool {

	for i := 0; i < len(*groups); i++ {
		sum := 0
		for _, friend := range (*groups)[i] {
			sum = sum + friend.owes
		}
		if sum != 0 {
			return false
		}
	}
	return true
}

func isPossible (numberOfFriends int, everybody *map[int]Friend) bool {
	groups := [][]Friend{}

	for i := 0; i < numberOfFriends; i++ {
		if friend := (*everybody)[i]; friend.inGroup {
			continue
		}
		group := []Friend{}
		divideInGroups(i, everybody, &group)
		groups = append(groups, group)
	}
	return CalculateSumForFriends(&groups)
}

func main() {
	var numberOfFriends int
	var numberOfFriendships int
	fmt.Scan(&numberOfFriends)
	fmt.Scan(&numberOfFriendships)

	everybody := make(map[int]Friend)

	for i := 0; i < numberOfFriends; i++ {
		var debt int
		fmt.Scan(&debt)
		friend := Friend{id: i, owes: debt}
		everybody[i] = friend
	}

	for i := 0; i < numberOfFriendships; i++ {
		var first int
		var second int

		fmt.Scan(&first)
		fmt.Scan(&second)

		firstFr := everybody[first]
		secondFr := everybody[second]
		firstFr.addFriend(second)
		secondFr.addFriend(first)
		everybody[first] = firstFr
		everybody[second] = secondFr
	}

	if isPossible(numberOfFriends, &everybody) {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}
