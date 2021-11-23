package service

import (
	"fmt"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const (
	RoleClient = iota
	RoleAdmin
)

type Message interface {
	Data() string
}

type Store struct {
	db       *dbr.Connection
	Messages map[int]Message
}

type IPMessage struct {
	IP string
}

func (i *IPMessage) Data() string {
	return i.IP
}

type User struct {
	Username  string `db:"username"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	ID        int64  `db:"id"`
	ChatID    int64  `db:"chat_id"`
	Role      int    `db:"role"`
	Status    int    `db:"status"`
}

type Chat struct {
	Title  string `db:"title"`
	ID     int64  `db:"id"`
	ChatID int64  `db:"chat_id"`
	Role   int    `db:"role"`
}

func NewStore(dbname, user, password string) (*Store, error) {
	conn, err := dbr.Open("sqlite3", fmt.Sprintf("file:%s.db?_loc=auto&_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=sha512", dbname, user, password), nil)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	return &Store{db: conn, Messages: make(map[int]Message)}, nil
}

func (s *Store) CreateUser(u User) (int64, error) {
	result, err := s.db.NewSession(nil).
		InsertInto("users").
		Columns("chat_id, username, first_name, last_name, role, status").
		Values(u.ChatID, u.Username, u.FirstName, u.LastName, u.Role, u.Status).
		Exec()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Store) Admins() ([]*User, error) {
	var users []*User

	_, err := s.db.NewSession(nil).
		Select("*").
		From("users").
		Where(dbr.Eq("role", RoleAdmin)).
		Load(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) Clients() ([]*User, error) {
	var users []*User

	_, err := s.db.NewSession(nil).
		Select("*").
		From("users").
		Where(dbr.Eq("role", RoleClient)).
		Load(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) DeleteUser(id int64) error {
	_, err := s.db.NewSession(nil).
		DeleteFrom("users").
		Where(dbr.Eq("id", id)).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateChat(c Chat) error {
	_, err := s.db.NewSession(nil).
		InsertInto("chats").
		Columns("chat_id, role", "title").
		Values(c.ChatID, c.Role, c.Title).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ClientChats() ([]*Chat, error) {
	var chats []*Chat

	_, err := s.db.NewSession(nil).
		Select("*").
		From("chats").
		Where(dbr.Eq("role", RoleClient)).
		Load(&chats)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (s *Store) AdminChats() ([]*Chat, error) {
	var chats []*Chat

	_, err := s.db.NewSession(nil).
		Select("*").
		From("chats").
		Where(dbr.Eq("role", RoleAdmin)).
		Load(&chats)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (s *Store) DeleteChat(id int64) error {
	_, err := s.db.NewSession(nil).
		DeleteFrom("users").
		Where(dbr.Eq("id", id)).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ChatByID(chatID int64) (*Chat, error) {
	var chat *Chat

	_, err := s.db.NewSession(nil).
		Select("*").
		From("chats").
		Where(dbr.Eq("id", chatID)).
		Load(&chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
