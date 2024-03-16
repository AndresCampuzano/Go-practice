package main

import "fmt"

// Add a helper function to check if a string is in a slice
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// Deletes an occurrence of a string from a string slice.
func deleteFromStringSlice(slice []string, s string) []string {
	for i := range slice {
		if slice[i] == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	// return the original slice if s is not found
	return slice
}

func findMutualFriends(username string, friendships map[string][]string) []string {
	// Search all friends of given username's friends: friendsOfFriends
	// Delete all friends from given username
	// Delete all duplicates
	usernameFriends := friendships[username]
	var friendsOfFriends []string
	for _, friend := range usernameFriends {
		friendsOfFriends = append(friendsOfFriends, friendships[friend]...)
	}

	// Deleting usernameFriends from friendsOfFriends
	for _, friend := range usernameFriends {
		deleteFromStringSlice(friendsOfFriends, friend)
	}

	// Delete duplicate friends
	var cleanedFriends []string
	for _, friend := range friendsOfFriends {
		if !stringInSlice(friend, cleanedFriends) {
			cleanedFriends = append(cleanedFriends, friend)
		}
	}

	// Delete given username from slice
	return deleteFromStringSlice(cleanedFriends, username)
}

func main() {
	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}

	fmt.Println(findMutualFriends("Alice", friendships))
	// [David Eve]
}
