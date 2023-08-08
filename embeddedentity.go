package main

import "fmt"

type Profile struct {
    Bio     string
    Website string
    // Other profile fields
}

type User struct {
    Username string
    Email    string
    Profile  Profile // Embedded profile
    // Other user fields
}

func main() {
    // Creating a profile
    profile := Profile{
        Bio:     "Software developer interested in Go programming.",
        Website: "https://example.com",
    }

    // Creating a user with the associated profile
    user := User{
        Username: "john_doe",
        Email:    "john@example.com",
        Profile:  profile,
    }

    // Accessing user's profile fields
    fmt.Println("User:", user.Username, user.Email)
    fmt.Println("Profile:", user.Profile.Bio, user.Profile.Website)
}
