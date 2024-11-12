// TUGAS BESAR - CII1F4 ALGORITMA PEMROGRAMAN
// SEMESTER GENAP 2022/2023, FAKULTAS INFORMATIKA
// Dosen: DRI Kelas: IF4603

// Jihan Auliyah Muslimin (1301223064)
// Dinda Desfira (1301223236)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NMAX int = 2023

type Account struct {
	NewUsername string
	NewPassword string
	Next        *Account
}

type Tiles struct {
	Value1 int
	Value2 int
}

type Arr struct {
	ArrTiles [28]Tiles
	CountT   int
}

type User struct {
	Name    [NMAX]string
	WinRate [NMAX]float64
	CountG  [NMAX]int
	CountU  int
}

type Player struct {
	Name    string
	WinRate float64
	CountG  int
	Tiles   Arr
}

type arr [28]int

func main() {
	var head *Account
	var choice string
	for {
		intro()
		fmt.Println("Welcome!")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print(">> ")
		fmt.Scan(&choice)
		if choice == "1" {
			head = signin(head)
		} else if choice == "2" {
			login(head)
		} else if choice == "3" {
			fmt.Println()
			fmt.Println("Goodbye!")
			return
		} else {
			fmt.Println("Invalid input, try again")
			fmt.Println()
		}
	}
}

/* display */

func intro() {
	fmt.Println(" ---------------------------------------------------------------------------------")
	fmt.Println("|                              Domino Solitaire Ceme                              |")
	fmt.Println("|                   Created by Dinda Desfira and Jihan Auliyah M                  |")
	fmt.Println("|                                     IF-46-03                                    |")
	fmt.Println("|                             Algoritma Pemrograman 2023                          |")
	fmt.Println(" ---------------------------------------------------------------------------------")
	fmt.Println()
}

func welcome() {
	fmt.Println("======================= Welcome to the game of Ceme-4tile =======================")
}

func logedin() {
	fmt.Println(" ---------------------------------------------------------------------------------")
	fmt.Println("|                             Domino Solitaire Ceme                               |")
	fmt.Println(" ---------------------------------------------------------------------------------")
	fmt.Println()
}

func rules() {
	fmt.Println(".................................................................................")
	fmt.Println()
	fmt.Println("Here are what you'll need to know!")
	fmt.Println()
	fmt.Println("Rules:")
	fmt.Println("1. If only one of the players has a double, the player who has it wins.")
	fmt.Println("2. If both players have a double, the player with the highest double value wins.")
	fmt.Println("3. If no player has a double, the player with higher tile value wins.")
	fmt.Println("* double = two tiles with the same number of values.")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("1 = change tile 1")
	fmt.Println("2 = change tile 2")
	fmt.Println("3 = change tile 3")
	fmt.Println("4 = change tile 4")
	fmt.Println("0 = done")
	fmt.Println("9 = exit")
	fmt.Println()
	fmt.Println(".................................................................................")
	fmt.Println()
}

func rankings() {
	fmt.Println("=================================== Rankings ====================================")
	fmt.Println()
	fmt.Println("1. Show rank from the highest winning rate.")
	fmt.Println("2. Show rank from the lowest winning rate.")
	fmt.Println("3. Show the rounds the player has played from the highest.")
	fmt.Println("4. Show the rounds the player has played from the lowest.")
	fmt.Println("5. Find players winning rate.")
	fmt.Println("6. Back.")
	fmt.Print(">> ")
}

func menu_options() {
	fmt.Println("Menu : ")
	fmt.Println("1. Start game")
	fmt.Println("2. Show rankings")
	fmt.Println("3. Exit game")
	fmt.Print(">> ")
}

/* account */

func signin(head *Account) *Account {
	var user Account
	fmt.Println("\n-- Sign Up --")
	fmt.Print("Enter username: ")
	fmt.Scan(&user.NewUsername)
	if username_exists(head, user.NewUsername) {
		fmt.Println("Username already exists")
		fmt.Println()
		return head
	}
	fmt.Print("Enter password: ")
	fmt.Scan(&user.NewPassword)
	fmt.Println("Sign in successful")
	fmt.Println()
	if head == nil {
		head = &user
	} else {
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &user
	}
	return head
}

func username_exists(head *Account, username string) bool {
	current := head
	for current != nil {
		if current.NewUsername == username {
			return true
		}
		current = current.Next
	}
	return false
}

func login(head *Account) {
	fmt.Println("\n-- Log In --")
	var username, password string
	var player User
	var found bool = false

	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	current := head
	for current != nil {
		if current.NewUsername == username && current.NewPassword == password {
			fmt.Println("Login success")
			fmt.Println()
			found = true
			menu(&player)
		}
		current = current.Next
	}

	if !found {
		fmt.Println("Login failed")
		fmt.Println()
	}
}

/* loged in */

func menu(player *User) {
	var choice string
	var name string
	logedin()
	menu_options()
	fmt.Scan(&choice)
	fmt.Println()
	if choice == "1" {
		start_game(player)
	} else if choice == "2" {
		rankings()
		fmt.Scan(&choice)
		for choice != "6" {
			if choice == "1" || choice == "2" {
				print_winrate(*player, choice)
			} else if choice == "3" || choice == "4" {
				print_gamecount(*player, choice)
			} else if choice == "5" {
				fmt.Println()
				fmt.Println("Who would like to search for?")
				fmt.Print("Enter players name: ")
				fmt.Scan(&name)
				fmt.Println()
				print_searchedP(*player, name)
				fmt.Println()
			}
			rankings()
			fmt.Scan(&choice)
		}
		fmt.Println()
		menu(player)
	} else if choice != "3" {
		fmt.Println("Invalid input, try again")
		fmt.Println()
		menu(player)
	}
}

/* sort */

// Selection sort

func rank_desc(player *User) {
	for i := 0; i < player.CountU-1; i++ {
		maxIndex := i
		for j := i + 1; j < player.CountU; j++ {
			if player.WinRate[j] > player.WinRate[maxIndex] {
				maxIndex = j
			}
		}
		tempName := player.Name[i]
		player.Name[i] = player.Name[maxIndex]
		player.Name[maxIndex] = tempName
		tempWinRate := player.WinRate[i]
		player.WinRate[i] = player.WinRate[maxIndex]
		player.WinRate[maxIndex] = tempWinRate
	}
}

func rank_asc(player *User) {
	for i := 0; i < player.CountU-1; i++ {
		minIndex := i
		for j := i + 1; j < player.CountU; j++ {
			if player.WinRate[j] < player.WinRate[minIndex] {
				minIndex = j
			}
		}
		tempName := player.Name[i]
		player.Name[i] = player.Name[minIndex]
		player.Name[minIndex] = tempName
		tempWinRate := player.WinRate[i]
		player.WinRate[i] = player.WinRate[minIndex]
		player.WinRate[minIndex] = tempWinRate
	}
}

// Insertion sort

func round_desc(player *User) {
	for i := 1; i < player.CountU; i++ {
		keyName := player.Name[i]
		keyTiles := player.CountG[i]
		j := i - 1

		for j >= 0 && player.CountG[j] < keyTiles {
			player.Name[j+1] = player.Name[j]
			player.CountG[j+1] = player.CountG[j]
			j = j - 1
		}

		player.Name[j+1] = keyName
		player.CountG[j+1] = keyTiles
	}
}

func round_asc(player *User) {
	for i := 1; i < player.CountU; i++ {
		keyName := player.Name[i]
		keyTiles := player.CountG[i]
		j := i - 1

		for j >= 0 && player.CountG[j] > keyTiles {
			player.Name[j+1] = player.Name[j]
			player.CountG[j+1] = player.CountG[j]
			j = j - 1
		}

		player.Name[j+1] = keyName
		player.CountG[j+1] = keyTiles
	}
}

/* tile */

func tiles(t *Arr) {
	for i := 0; i < 28; i++ {
		if i < 7 {
			t.ArrTiles[i].Value1 = 0
			t.ArrTiles[i].Value2 = i
		} else if i < 13 {
			t.ArrTiles[i].Value1 = 1
			t.ArrTiles[i].Value2 = i - 6
		} else if i < 18 {
			t.ArrTiles[i].Value1 = 2
			t.ArrTiles[i].Value2 = i - 11
		} else if i < 22 {
			t.ArrTiles[i].Value1 = 3
			t.ArrTiles[i].Value2 = i - 15
		} else if i < 25 {
			t.ArrTiles[i].Value1 = 4
			t.ArrTiles[i].Value2 = i - 18
		} else if i < 27 {
			t.ArrTiles[i].Value1 = 5
			t.ArrTiles[i].Value2 = i - 20
		} else {
			t.ArrTiles[i].Value1 = 6
			t.ArrTiles[i].Value2 = 6
		}
	}
	t.CountT = 28
}

func shuffle_tiles(t *Arr) {
	var indeks [28]int
	var T Arr
	rand.Seed(time.Now().UnixNano())
	var random int = rand.Intn(28)
	for i := 0; i < t.CountT; i++ {
		for check_ada(indeks, i, random) {
			random = rand.Intn(28)
		}
		indeks[i] = random
	}
	for i := 0; i < t.CountT; i++ {
		T.ArrTiles[i].Value1 = t.ArrTiles[indeks[i]].Value1
		T.ArrTiles[i].Value2 = t.ArrTiles[indeks[i]].Value2
	}
	for i := 0; i < t.CountT; i++ {
		t.ArrTiles[i].Value1 = T.ArrTiles[i].Value1
		t.ArrTiles[i].Value2 = T.ArrTiles[i].Value2
	}
}

func take_tiles(tUser, t *Arr) {
	rand.Seed(time.Now().UnixNano())
	var indeks int = rand.Intn(t.CountT)
	if tUser.CountT < 4 {
		tUser.ArrTiles[tUser.CountT].Value1 = t.ArrTiles[indeks].Value1
		tUser.ArrTiles[tUser.CountT].Value2 = t.ArrTiles[indeks].Value2
		tUser.CountT++
		delete_tile(t.ArrTiles[indeks].Value1, t.ArrTiles[indeks].Value2, t)
	}
}

func delete_tile(value1, value2 int, t *Arr) {
	var i int = search_tile(value1, value2, *t)
	if i != -1 {
		for j := i; j < 27; j++ {
			t.ArrTiles[j].Value1 = t.ArrTiles[j+1].Value1
			t.ArrTiles[j].Value2 = t.ArrTiles[j+1].Value2
		}
		t.CountT--
	}
}

func change_tile(choice int, user *Player, t *Arr) {
	rand.Seed(time.Now().UnixNano())
	var indeks int = rand.Intn(t.CountT)
	user.Tiles.ArrTiles[choice-1].Value1 = t.ArrTiles[indeks].Value1
	user.Tiles.ArrTiles[choice-1].Value2 = t.ArrTiles[indeks].Value2
	delete_tile(t.ArrTiles[indeks].Value1, t.ArrTiles[indeks].Value2, t)
}

func count_tiles(t Arr) int {
	var max int = 0
	for i := 0; i < t.CountT; i++ {
		if max < t.ArrTiles[i].Value1+t.ArrTiles[i].Value2 {
			max = t.ArrTiles[i].Value1 + t.ArrTiles[i].Value2
		}
	}
	return max
}

/* game */

func start_game(A *User) {
	welcome()
	var user, opponent Player
	var win_count, choice, game_count, options int
	choice = 7
	game_count = 0
	options = 0
	fmt.Println()
	fmt.Print("How would you like us to call you? ")
	fmt.Scan(&user.Name)
	fmt.Println("Hello, ", user.Name, "! :)")
	fmt.Println()
	rules()
	fmt.Println("Your score is 0/0")
	for choice != 9 {
		game(&user, &opponent, &choice, &options)
		if choice == 0 || options == 2 {
			if check_win(user, opponent) {
				fmt.Println("You Win")
				win_count++
			} else {
				fmt.Println("You Lost")
			}
			fmt.Print("Dealers Tiles: ")
			print_tiles(opponent.Tiles)
			fmt.Println()
			game_count++
			fmt.Println("Your score is ", win_count, "/", game_count)
			fmt.Println()
		}
		user.Tiles.CountT = 0
		opponent.Tiles.CountT = 0
	}
	fmt.Println("Your last score is ", win_count, "/", game_count)
	fmt.Println("Thank you for playing with us.")
	winrate := (float64(win_count) / float64(game_count)) * 100
	fmt.Println("Your winning rate is ", winrate, "%")
	A.Name[A.CountU] = user.Name
	user.WinRate = winrate
	A.WinRate[A.CountU] = user.WinRate
	user.CountG = game_count
	A.CountG[A.CountU] = user.CountG
	A.CountU++
	check_name(A, user, A.CountU-1)
	menu(A)
}

func game(user, opponent *Player, choice, options *int) {
	var t Arr
	var decision, commands, tile int
	tiles(&t)
	shuffle_tiles(&t)
	fmt.Println("Dealing... ")
	take_tiles(&user.Tiles, &t)
	take_tiles(&user.Tiles, &t)
	take_tiles(&user.Tiles, &t)
	take_tiles(&user.Tiles, &t)
	tile = 4
	take_tiles(&opponent.Tiles, &t)
	take_tiles(&opponent.Tiles, &t)
	take_tiles(&opponent.Tiles, &t)
	take_tiles(&opponent.Tiles, &t)
	take_tiles(&user.Tiles, &t)
	take_tiles(&user.Tiles, &t)
	fmt.Print("Your Tiles: ")
	print_tiles(user.Tiles)
	fmt.Println()
	fmt.Print("Decision? ")
	fmt.Scan(&decision)
	for decision != 0 && commands < 2 && decision != 9 {
		if decision > tile {
			fmt.Println("Tile has not been taken!")
			commands--
		} else {
			change_tile(decision, user, &t)
			fmt.Print("Your Tiles: ")
			print_tiles(user.Tiles)
			fmt.Println()
		}
		fmt.Print("Decision? ")
		fmt.Scan(&decision)
		*choice = decision
		commands++
		*options = commands
	}
	*choice = decision
	*options = commands
}

/* search */

// Sequential Search
func search_tile(value1, value2 int, t Arr) int {
	for i := 0; i < t.CountT; i++ {
		if t.ArrTiles[i].Value1 == value1 && t.ArrTiles[i].Value2 == value2 {
			return i
		}
	}
	return -1
}

// Binary Search
func search_player(t User, X string) int {
	rank_asc(&t)
	var right, left, mid int
	left = 0
	right = t.CountU - 1
	for left <= right {
		mid = (right + left) / 2
		if t.Name[mid] > X {
			right = mid - 1
		} else if t.Name[mid] < X {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

/* check */

func check_ada(t arr, n, x int) bool {
	for i := 0; i < n; i++ {
		if t[i] == x {
			return true
		}
	}
	return false
}

func check_same(t Arr) bool {
	for i := 0; i < t.CountT; i++ {
		if t.ArrTiles[i].Value1 != t.ArrTiles[i].Value2 {
			return false
		}
	}
	return true
}

func check_win(user, opponent Player) bool {
	if check_same(user.Tiles) && !check_same(opponent.Tiles) {
		return true
	} else if check_same(user.Tiles) && check_same(opponent.Tiles) {
		if count_tiles(user.Tiles) > count_tiles(opponent.Tiles) {
			return true
		} else {
			return false
		}
	} else {
		if count_tiles(user.Tiles) > count_tiles(opponent.Tiles) {
			return true
		} else {
			return false
		}
	}
}

func check_name(z *User, x Player, n int) {
	var cek bool = false
	for i := 0; i < z.CountU; i++ {
		if z.Name[i] == x.Name && i != n {
			cek = true
			if z.WinRate[i] < x.WinRate {
				z.WinRate[i] = x.WinRate
			}
		}
	}
	if cek {
		z.CountU--
	}
}

/* print */

func print_tiles(t Arr) {
	for i := 0; i < t.CountT; i++ {
		fmt.Print(t.ArrTiles[i], " ")
	}
}

func print_winrate(t User, choice string) {
	if choice == "1" {
		rank_desc(&t)
	} else if choice == "2" {
		rank_asc(&t)
	}

	fmt.Println()
	fmt.Println(".................................................................................")
	fmt.Println()
	for i := 0; i < t.CountU; i++ {
		fmt.Println("Name: ", t.Name[i], "     Win rate: ", t.WinRate[i])
	}
	fmt.Println()
	fmt.Println(".................................................................................")
	fmt.Println()
}

func print_searchedP(A User, name string) {
	var found bool = false
	fmt.Println("Searching for player...")
	fmt.Println()
	for i := 0; i < A.CountU; i++ {
		if A.Name[i] == name {
			fmt.Println(".................................................................................")
			fmt.Println()
			fmt.Println("Player: ", A.Name[i])
			fmt.Println("Winning Rate: ", A.WinRate[i], "%")
			fmt.Println("Rounds Played: ", A.CountG[i])
			fmt.Println()
			fmt.Println(".................................................................................")
			found = true
			return
		}
	}
	if !found {
		fmt.Println("Player not found")
	}
}

func print_gamecount(t User, choice string) {
	if choice == "3" {
		round_desc(&t)
	} else if choice == "4" {
		round_asc(&t)
	}
	fmt.Println()
	fmt.Println(".................................................................................")
	fmt.Println()
	for i := 0; i < t.CountU; i++ {
		fmt.Println("Name: ", t.Name[i], "     Rounds: ", t.CountG[i])
	}
	fmt.Println()
	fmt.Println(".................................................................................")
	fmt.Println()
}
