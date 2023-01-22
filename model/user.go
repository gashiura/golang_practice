package model

type User struct {
  ID        int    `json:"id"`
  UID       string `json:"uid"`
  Email     string `json:"email"`
  Password  string `json:"-"`
}
