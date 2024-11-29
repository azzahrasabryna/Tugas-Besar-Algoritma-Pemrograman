package main

import (
	"fmt"
)

// Deklarasi variabel NMAX yang merupakan jumlah maksimum Array dalam program ini.
const NMAX = 100

// Struktur untuk variabel pengguna
type User struct {
	username     string
	password     string
	statuses     content
	friendNames  [NMAX]string
	friendsCount int
	statusCount  int
}

// Array untuk isi content pada setiap status
type content [NMAX]isiContent

// Struktur untuk isi content pada status
type isiContent struct {
	isi          string
	comments     [NMAX]string
	commentCount int
}

// Array untuk variabel pengguna
type users [NMAX]User

var userCount int = 0
var isValid bool
var ID int = -1

// Fungsi utama dari program berikut.
func main() {
	var choice int
	var pengguna users

	isValid = false
	for !isValid {
		fmt.Println("==========================================================")
		fmt.Println("HELLO WELCOME TO TELLYOU")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("  If you already have an account please select number 2,  ")
		fmt.Println("   but if u haven't please select number 1 to register.   ")
		fmt.Println("==========================================================")
		fmt.Printf("Select one (1/2/3): ")
		fmt.Scan(&choice)

		if choice == 1 {
			registrasi(&pengguna)
			fmt.Println()
		} else if choice == 2 {
			login(&pengguna)
			fmt.Println()
		} else if choice == 3 {
			isValid = true
		} else if choice < 1 || choice > 3 {
			fmt.Println("Invalid choice.")
			fmt.Println()
		}
	}
}

/*
	Prosedur menu yang ditampilkan ketika pengguna berhasil login

dengan parameter in/out, yaitu pengguna bertipe users
*/
func menuLogin(pengguna *users) {
	var choice int

	fmt.Println()
	for ID != -1 {
		fmt.Println("===============================")
		fmt.Println("             HOME              ")
		fmt.Println("===============================")
		fmt.Println("1. Add Status")
		fmt.Println("2. View Statuses")
		fmt.Println("3. Add Comment")
		fmt.Println("4. View Comments")
		fmt.Println("5. Add Friend")
		fmt.Println("6. View Friends")
		fmt.Println("7. Remove Friend")
		fmt.Println("8. Edit Profile")
		fmt.Println("9. View Profile")
		fmt.Println("10. Search user")
		fmt.Println("11. Logout")
		fmt.Println("===============================")
		fmt.Printf("Select one (1-11): ")
		fmt.Scan(&choice)

		if choice == 1 {
			addStatus(&*pengguna)
		} else if choice == 2 {
			viewStatuses(*pengguna)
		} else if choice == 3 {
			addComment(&*pengguna)
		} else if choice == 4 {
			viewComments(*pengguna)
		} else if choice == 5 {
			addFriend(&*pengguna)
		} else if choice == 6 {
			viewFriends(*pengguna)
		} else if choice == 7 {
			removeFriend(&*pengguna)
		} else if choice == 8 {
			editProfile(&*pengguna)
		} else if choice == 9 {
			viewProfile(*pengguna)
		} else if choice == 10 {
			searchUser(*pengguna)
		} else if choice == 11 {
			logout()
		} else {
			fmt.Println("Invalid choice")
		}
		fmt.Println()
	}
}

/*
	Prosedur untuk mencari pengguna menggunakan username

yang terdaftar dalam aplikasi dan menampilkan
beberapa detail informasi akun pengguna. Prosedur ini
memiliki parameter in, yaitu pengguna bertipe users
*/
func searchUser(pengguna users) {
	var choice, save int
	var username string

	save = -1

	fmt.Println("Menu Search")
	fmt.Println("1. Search Friend")
	fmt.Println("2. Search User")
	fmt.Println("3. Back")
	fmt.Printf("Select one (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Enter friend's username: ")
		fmt.Scan(&username)
		for i := 0; i < pengguna[ID].friendsCount; i++ {
			if pengguna[ID].friendNames[i] == username {
				for k := 0; k < userCount; k++ {
					if pengguna[k].username == username {
						save = k
					}
				}
			}
		}
		if save != -1 {
			fmt.Println("Profile Information")
			fmt.Println("Username:", pengguna[save].username)
			fmt.Println("Total Friends:", pengguna[save].friendsCount)
			fmt.Println("Total Statuses:", pengguna[save].statusCount)
		} else {
			fmt.Println("User not found.")
		}
	} else if choice == 2 {
		fmt.Printf("Enter username: ")
		fmt.Scan(&username)
		for i := 0; i < userCount; i++ {
			if pengguna[i].username == username {
				save = i
			}
		}
		if save != -1 {
			fmt.Println("Profile information:")
			fmt.Println("Username:", pengguna[save].username)
			fmt.Println("Total Friends:", pengguna[save].friendsCount)
			fmt.Println("Total Statuses:", pengguna[save].statusCount)
		} else {
			fmt.Println("User not found.")
		}
	} else if choice < 1 || choice > 3 {
		fmt.Println("Invalid choice.")
	}

}

/*
	Prosedur menu registrasi akun pengguna dengan parameter

in/out, yaitu pengguna bertipe users
*/
func registrasi(pengguna *users) {
	var username, password string
	var isAvailable bool = false

	fmt.Printf("Enter username: ")
	fmt.Scan(&username)
	for i := 0; i < userCount; i++ {
		if pengguna[i].username == username {
			fmt.Println("Username already taken.")
			isAvailable = isAvailable || true
		}
	}
	if !isAvailable && userCount < NMAX {
		pengguna[userCount].username = username
		fmt.Printf("Enter password: ")
		fmt.Scan(&password)
		pengguna[userCount].password = password
		userCount++
		fmt.Println("Registration successful.")
	} else if userCount >= NMAX {
		fmt.Println("Maximum number of accounts reached. No new accounts can be created at this time.")
	}
}

/*
	Prosedur login masuk ke aplikasi dengan parameter

in/out, yaitu pengguna bertipe users
*/
func login(pengguna *users) {
	var username, password string
	var isSuccess bool

	fmt.Printf("Username: ")
	fmt.Scan(&username)
	fmt.Printf("Password: ")
	fmt.Scan(&password)

	for i := 0; i < userCount; i++ {
		if pengguna[i].username == username && pengguna[i].password == password {
			isSuccess = isSuccess || true
			ID = i
		}
	}
	if isSuccess {
		fmt.Println("Login Berhasil")
		menuLogin(&*pengguna)
	} else {
		fmt.Println("Username atau password tidak ditemukan")
	}
}

// Prosedur untuk logout dari akun pengguna
func logout() {
	ID = -1
	isValid = false
	fmt.Println("Berhasil Keluar")
}

/*
	Prosedur yang digunakan ketika pengguna ingin menambah status.

Prosedur ini memiliki parameter in/out, yaitu pengguna bertipe users
*/
func addStatus(pengguna *users) {
	var content string

	if pengguna[ID].statusCount < NMAX {
		fmt.Printf("Enter status content: ")
		fmt.Scan(&content)
		pengguna[ID].statuses[pengguna[ID].statusCount].isi = content
		pengguna[ID].statusCount++
		fmt.Println("Status added.")
	} else {
		fmt.Println("Status limit reached.")
	}
}

/*
	Prosedur yang digunakan ketika pengguna ingin melihat

status apa saja yang telah ditambahkan oleh pengguna
aplikasi ini. Prosedur ini memiliki parameter in,
yaitu pengguna dengan tipe users
*/
func viewStatuses(pengguna users) {
	var choice int
	var check bool

	fmt.Println("Status Menu")
	fmt.Println("1. View your status")
	fmt.Println("2. View your friend's status")
	fmt.Println("3. View all status in this app")
	fmt.Println("4. Back")
	fmt.Printf("Select one (1/2/3/4): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("+---------------------------------------------+")
		fmt.Printf("| Your statuses\n")
		if pengguna[ID].statusCount <= 0 {
			fmt.Println("| There's no status that has been added.")
		} else {
			for k := 0; k < pengguna[ID].statusCount; k++ {
				fmt.Printf("| %d. %s\n", k+1, pengguna[ID].statuses[k].isi)
			}
		}
		fmt.Println("+---------------------------------------------+")

	} else if choice == 2 {
		if pengguna[ID].friendsCount <= 0 {
			fmt.Println("You haven't added anyone yet.")
		} else {
			for i := 0; i < pengguna[ID].friendsCount; i++ {
				for j := 0; j < userCount; j++ {
					if pengguna[ID].friendNames[i] == pengguna[j].username {
						fmt.Println("==================================================")
						fmt.Printf("%d. %s's statuses\n", i+1, pengguna[j].username)
						fmt.Println("--------------------------------------------------")
						if pengguna[j].statusCount <= 0 {
							fmt.Println("There's no status that has been added.")
						} else {
							for k := 0; k < pengguna[j].statusCount; k++ {
								fmt.Printf("%d. %s\n", k+1, pengguna[j].statuses[k].isi)
							}
						}
						fmt.Println("==================================================")
					}
				}
			}
		}
	} else if choice == 3 {
		for i := 0; i < userCount; i++ {
			if pengguna[i].statusCount > 0 {
				fmt.Printf("------------%s's statuses------------\n", pengguna[i].username)
				for j := 0; j < pengguna[i].statusCount; j++ {
					fmt.Printf("%d. %s\n", j+1, pengguna[i].statuses[j].isi)
				}
				check = check || true
			}
		}
		if !check {
			fmt.Println("There's no status yet.")
		}
	} else if choice < 1 || choice > 4 {
		fmt.Println("Invalid choice.")
	}
}

/*
	Prosedur yang digunakan ketika pengguna ingin menambah

komentar ke status yang dibuat oleh pengguna lain.
Prosedur ini memiliki parameter in/out, yaitu
pengguna bertipe users
*/
func addComment(pengguna *users) {
	var number, choice, save int
	var comment, username string

	fmt.Println("Comment Menu")
	fmt.Println("1. Comment on your friend's page")
	fmt.Println("2. Comment on all user page")
	fmt.Println("3. Back")
	fmt.Printf("Select one (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Enter friend's username to comment on: ")
		fmt.Scan(&username)
		for i := 0; i < pengguna[ID].friendsCount; i++ {
			if pengguna[ID].friendNames[i] == username {
				fmt.Printf("%s's statuses\n", username)
				for k := 0; k < userCount; k++ {
					if pengguna[k].username == username {
						save = k
					}
				}
				for j := 0; j < pengguna[save].statusCount; j++ {
					fmt.Printf("%d. %s\n", j+1, pengguna[save].statuses[j].isi)
				}

				fmt.Printf("Enter number: ")
				fmt.Scan(&number)

				if number > 0 && number < pengguna[ID].friendsCount {
					fmt.Printf("Enter comment: ")
					fmt.Scan(&comment)
					pengguna[save].statuses[number-1].comments[pengguna[save].statuses[number-1].commentCount] = comment

					pengguna[save].statuses[number-1].commentCount++
					fmt.Println("Comment added.")
				}
			}
		}
	} else if choice == 2 {
		fmt.Printf("Enter username to comment on: ")
		fmt.Scan(&username)
		for i := 0; i < userCount; i++ {
			if pengguna[i].username == username {
				fmt.Printf("%s's statuses\n", username)
				for j := 0; j < pengguna[i].statusCount; j++ {
					fmt.Printf("%d. %s\n", j+1, pengguna[i].statuses[j].isi)
				}

				fmt.Printf("Enter number: ")
				fmt.Scan(&number)

				if number > 0 && number < userCount {
					fmt.Printf("Enter comment: ")
					fmt.Scan(&comment)
					pengguna[i].statuses[number-1].comments[pengguna[i].statuses[number-1].commentCount] = comment

					pengguna[i].statuses[number-1].commentCount++
					fmt.Println("Comment added.")
				}
			}
		}
	}
}

/*
	Prosedur yang digunakan ketika pengguna ingin melihat

atau menampilkan isi komentar dari setiap status yang
terdapat di dalam aplikasi. Prosedur ini memiliki parameter
in, yaitu pengguna bertipe users
*/
func viewComments(pengguna users) {
	var number, i int
	var username string
	fmt.Printf("Enter username to view comments: ")
	fmt.Scan(&username)

	i = 0
	for i < userCount {
		if pengguna[i].username == username {
			fmt.Printf("%s's statuses\n", username)
			for j := 0; j < pengguna[i].statusCount; j++ {
				fmt.Printf("%d. %s\n", j+1, pengguna[i].statuses[j].isi)
			}
			if pengguna[i].statusCount > 0 {
				fmt.Printf("Enter number: ")
				fmt.Scan(&number)

				for k := 0; k < pengguna[i].statuses[number-1].commentCount; k++ {
					fmt.Printf("%d. %s\n", k+1, pengguna[i].statuses[number-1].comments[k])
				}
				if pengguna[i].statuses[number-1].commentCount == 0 {
					fmt.Println("There's no comment here.")
				}
			} else {
				fmt.Println("There's no status yet.")
			}
		}
		i++
	}
}

/*
	Prosedur yang digunakan jika pengguna ingin menambah

teman di dalam aplikasi. Prosedur ini memiliki parameter
in/out, yaitu pengguna bertipe users
*/
func addFriend(pengguna *users) {
	var friendUsername string
	var notFound int = -1
	var isAvailable bool = false

	if pengguna[ID].friendsCount < NMAX {
		fmt.Printf("Enter friend's username: ")
		fmt.Scan(&friendUsername)
		if friendUsername == pengguna[ID].username {
			fmt.Println("You cannot add yourself as a friend.")
			notFound = 3
		} else {
			for i := 0; i < userCount; i++ {
				if i != ID {
					if pengguna[i].username == friendUsername {
						if pengguna[ID].friendsCount == 0 {
							notFound = 2
						}
						for j := 0; j < pengguna[ID].friendsCount; j++ {
							if pengguna[ID].friendNames[j] == friendUsername {
								isAvailable = isAvailable || true
								notFound = 1
							} else {
								if !isAvailable {
									notFound = 2
								}
							}
						}
					}
				}
			}
			if isAvailable && notFound == 1 {
				fmt.Println("Friend already added.")

			} else if notFound == 2 {
				pengguna[ID].friendNames[pengguna[ID].friendsCount] = friendUsername
				pengguna[ID].friendsCount++
				fmt.Println("Success. Friend added.")
			} else if notFound == -1 {
				fmt.Println("There's no account with this username.")
			}

		}
	} else {
		fmt.Println("Sorry, friend limit reached.")
	}
}

/*
	Prosedur yang digunakan jika pengguna ingin melihat

teman mana yang telah ditambahkan oleh pengguna.
Prosedur ini memiliki parameter in, yaitu pengguna
bertipe users
*/
func viewFriends(pengguna users) {
	var choice int

	if pengguna[ID].friendsCount == 0 {
		fmt.Println("No friend to display.")
	} else {
		fmt.Println("View method:")
		fmt.Println("1. View friends from A-Z")
		fmt.Println("2. View friends from Z-A")
		fmt.Printf("Select one (1/2): ")
		fmt.Scan(&choice)
		if choice == 1 {
			selectionSort(&pengguna)
			cetakTeman(pengguna)
		} else if choice == 2 {
			insertionSort(&pengguna)
			cetakTeman(pengguna)
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

/*
	Prosedur yang digunakan jika pengguna ingin menampilkan

daftar teman yang telah dia tambahkan. Prosedur ini memiliki
parameter in yang dialiaskan sebagai A, bertipe users
*/
func cetakTeman(A users) {
	if A[ID].friendsCount > 1 {
		fmt.Println("You have", A[ID].friendsCount, "friends.")
	} else {
		fmt.Println("You have", A[ID].friendsCount, "friend.")
	}
	for i := 0; i < A[ID].friendsCount; i++ {
		fmt.Printf("%d. %s\n", i+1, A[ID].friendNames[i])
	}
}

/*
	Prosedur yang digunakan jika pengguna ingin mengurutkan

daftar teman yang telah ditambahkan berdasarkan abjad
secara menaik dengan menggunakan selection sort. Prosedur ini
memiliki parameter in/out yang dialiaskan sebagai A, bertipe users
*/
func selectionSort(A *users) {
	var pass, idxMin int
	var temp string

	for pass = 0; pass < A[ID].friendsCount; pass++ {
		idxMin = pass
		for i := pass + 1; i < A[ID].friendsCount; i++ {
			if A[ID].friendNames[idxMin] > A[ID].friendNames[i] {
				idxMin = i
			}
		}
		temp = A[ID].friendNames[idxMin]
		A[ID].friendNames[idxMin] = A[ID].friendNames[pass]
		A[ID].friendNames[pass] = temp
	}
}

/*
	Prosedur yang digunakan jika pengguna ingin mengurutkan

daftar teman yang telah ditambahkan berdasarkan abjad
secara menurun dengan menggunakan insertion sort. Prosedur ini
memiliki parameter in/out yang dialiaskan sebagai A, bertipe users
*/
func insertionSort(A *users) {
	var i, pass int
	var temp string

	for pass = 0; pass < A[ID].friendsCount; pass++ {
		temp = A[ID].friendNames[pass]
		i = pass

		for i > 0 && temp > A[ID].friendNames[i-1] {
			A[ID].friendNames[i] = A[ID].friendNames[i-1]
			i--
		}
		A[ID].friendNames[i] = temp
	}
}

/*
	Prosedur yang digunakan untuk menghapus teman yang sebelumnya

telah ditambahkan oleh pengguna. Prosedur ini memiliki
parameter in/out, yaitu pengguna bertipe users
*/
func removeFriend(pengguna *users) {
	var friendUsername string
	var friendsID int = -1
	if pengguna[ID].friendsCount == 0 {
		fmt.Println("There's no friend, please add someone first.")
	} else {
		fmt.Printf("Enter friend's username to remove: ")
		fmt.Scan(&friendUsername)
		for i := 0; i < pengguna[ID].friendsCount; i++ {
			if pengguna[ID].friendNames[i] == friendUsername {
				friendsID = i
			} else {
				fmt.Println("Friend not found.")
			}
		}
		if friendsID != -1 {
			for j := friendsID; j < pengguna[ID].friendsCount; j++ {
				pengguna[ID].friendNames[j] = pengguna[ID].friendNames[j+1]
			}
			pengguna[ID].friendsCount--
			fmt.Println("Friend removed.")
		}
	}
}

/*
	Prosedur yang digunakan untuk mengedit data pengguna,

yaitu username ataupun password. Prosedur ini memiliki parameter
in/out, yaitu pengguna bertipe users
*/
func editProfile(pengguna *users) {
	var newProfile string
	var choice int
	fmt.Println("Edit Menu:")
	fmt.Println("1. Edit Username")
	fmt.Println("2. Edit Password")
	fmt.Printf("Select one (1/2): ")
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Printf("New Username: ")
		fmt.Scan(&newProfile)
		pengguna[ID].username = newProfile
		fmt.Println("Success. Profile updated.")
	} else if choice == 2 {
		fmt.Printf("New Password: ")
		fmt.Scan(&newProfile)
		pengguna[ID].password = newProfile
		fmt.Println("Success. Profile updated.")
	}

}

/*
	Prosedur yang digunakan untuk menampilkan informasi dari

pengguna yang sedang login, yaitu berupa username, jumlah
teman, dan jumlah status yang telah ditambahkan.Prosedur
ini memiliki parameter in, yaitu pengguna bertipe users
*/
func viewProfile(pengguna users) {
	fmt.Println("Profile Information")
	fmt.Println("Username:", pengguna[ID].username)
	fmt.Println("Total Friends:", pengguna[ID].friendsCount)
	fmt.Println("Total Statuses:", pengguna[ID].statusCount)
}
