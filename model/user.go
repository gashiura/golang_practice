package model

import (
  "fmt"
)

type User struct {
  ID        int    `json:"id"`
  UID       string `json:"uid"`
  Email     string `json:"email"`
  Password  string `json:"-"`
}

var users = []User{
  { ID: 1, UID: "abcde", Email: "example1@test.com", Password: "Password1" },
  { ID: 2, UID: "fghij", Email: "example2@test.com", Password: "Password2" },
  { ID: 3, UID: "klmno", Email: "example3@test.com", Password: "Password3" },
}

func FindUser(userID int) (*User, error) {
  for _, u := range users {
    if u.ID == userID {
      return &u, nil
    }
  }
  return nil, &UserNotFoundError{id: userID}
}

type UserNotFoundError struct {
  id int
}

func (e *UserNotFoundError) Error() string {
  return fmt.Sprintf("UserID: %d not found", e.id)
}
