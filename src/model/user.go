package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

var passwordSalt = ""

// SetPasswordSalt sets the password salt
func SetPasswordSalt(salt string) {
	passwordSalt = salt
}

// User structure
type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

// Login a user
func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	row := db.QueryRow(`
	SELECT id,email,firstname,lastname
	FROM public.user
	WHERE email = $1 AND password = $2`, email, pwd)

	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}

	t := time.Now()
	_, err = db.Exec(`
	UPDATE public.user
	SET lastlogin = $1
	WHERE id = $2`, t, result.ID)
	if err != nil {
		log.Printf("Failed to update loging time for user %v to %v: %v", result.Email, t, err)
	}
	log.Printf("User %d logged in as %s at %v", result.ID, result.Email, t)
	return result, nil
}
