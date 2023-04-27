package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// User represents a user in the system
type User struct {
	ID   int
	Name string
}

// Friend represents a friend of a user
type Friend struct {
	ID   int
	Name string
}

// UserService represents the user service
type UserService struct {
	Users        map[int]*User
	Friends      map[int]*Friend
	UserToFriend map[int]int
	nextUserID   int
	nextFriendID int
}

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Name string
}

// CreateUserResponse represents the response after creating a user
type CreateUserResponse struct {
	ID int
}

// GetUserRequest represents the request to get a user
type GetUserRequest struct {
	ID int
}

// GetUserResponse represents the response after getting a user
type GetUserResponse struct {
	Name string
}

// CreateFriendRequest represents the request to create a friend
type CreateFriendRequest struct {
	Name string
}

// CreateFriendResponse represents the response after creating a friend
type CreateFriendResponse struct {
	ID int
}

// GetFriendRequest represents the request to get a friend
type GetFriendRequest struct {
	ID int
}

// GetFriendResponse represents the response after getting a friend
type GetFriendResponse struct {
	Name string
}

// AssignUserToFriendRequest represents the request to assign a user to a friend
type AssignUserToFriendRequest struct {
	UserID   int
	FriendID int
}

// AssignUserToFriendResponse represents the response after assigning a user to a friend
type AssignUserToFriendResponse struct {
}

// GetUserFriendRequest represents the request to get a user's friend
type GetUserFriendRequest struct {
	UserID int
}

// GetUserFriendResponse represents the response after getting a user's friend
type GetUserFriendResponse struct {
	Name string
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	user := &User{
		ID:   s.nextUserID,
		Name: req.Name,
	}

	s.Users[user.ID] = user
	s.nextUserID++

	return &CreateUserResponse{ID: user.ID}, nil
}

// GetUser returns the user details for a given user ID
func (s *UserService) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	user, ok := s.Users[req.ID]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return &GetUserResponse{Name: user.Name}, nil
}

// CreateFriend creates a new friend
func (s *UserService) CreateFriend(ctx context.Context, req *CreateFriendRequest) (*CreateFriendResponse, error) {
	friend := &Friend{
		ID:   s.nextFriendID,
		Name: req.Name,
	}

	s.Friends[friend.ID] = friend
	s.nextFriendID++

	return &CreateFriendResponse{ID: friend.ID}, nil
}

// GetFriend returns the friend details for a given friend ID
func (s *UserService) GetFriend(ctx context.Context, req *GetFriendRequest) (*GetFriendResponse, error) {
	friend, ok := s.Friends[req.ID]
	if !ok {
		return nil, fmt.Errorf("friend not found")
	}

	return &GetFriendResponse{Name: friend.Name}, nil
}

// AssignUserToFriend assigns a user to a friend
func (s *UserService) AssignUserToFriend(ctx context.Context, req *AssignUserToFriendRequest) (*AssignUserToFriendResponse, error) {
	_, userExists := s.Users[req.UserId]
	_, friendExists := s.Friends[req.FriendId]
	if !userExists || !friendExists {
		return nil, fmt.Errorf("invalid user or friend")
	}

	s.UserToFriend[req.UserId] = req.FriendId

	return &AssignUserToFriendResponse{}, nil
}

// GetUserFriend returns the friend of a user
func (s *UserService) GetUserFriend(ctx context.Context, req *GetUserFriendRequest) (*GetUserFriendResponse, error) {
	friendID, ok := s.UserToFriend[req.UserId]
	if !ok {
		return nil, fmt.Errorf("user does not have a friend")
	}

	friend, ok := s.Friends[friendID]
	if !ok {
		return nil, fmt.Errorf("friend not found")
	}

	return &GetUserFriendResponse{Name: friend.Name}, nil
}

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Initialize the user service
	userService := &UserService{
		Users:        make(map[int]*User),
		Friends:      make(map[int]*Friend),
		UserToFriend: make(map[int]int),
		nextUserID:   1,
		nextFriendID: 1,
	}

	// Register the user service with the gRPC server
	RegisterUserServiceServer(server, userService)

	// Create a TCP listener on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}

	// Start the gRPC server
	log.Println("Starting gRPC server...")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
