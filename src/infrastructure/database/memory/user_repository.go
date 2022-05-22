package memory

import (
  "fmt"
  "github.com/google/uuid"
  "github.com/a2-ito/go-onion/src/domain/model"
  "github.com/a2-ito/go-onion/src/domain/repository"
)

type userRepository struct {
  users  map[string]*model.User
}

func NewUserRepository() repository.UserRepository {
  return &userRepository{
    users: map[string]*model.User{},
  }
}

func newUser(id uuid.UUID, name string) *model.User {
  return &model.User{
    ID: id,
    Name: name,
  }
}

func (r *userRepository) Hello() error {
  fmt.Println("UserRepository Hello")
  return nil
}

func (u *userRepository) FindById(id uuid.UUID) (*model.User, error) {
  fmt.Println("UserRepository FindById")
  for _, ml := range u.users {
    if ml.ID == id {
      return ml, nil
    }
  }
  //return nil, fmt.Errorf("Error: %s", "no memo data")
  return nil, nil
}

func (u *userRepository) FindAll() ([]*model.User, error) {
  fmt.Println("UserRepository FindAll")
  users := make([]*model.User, len(u.users))
  //fmt.Println(u.users[0].ID)
  i := 0
  for _, user := range u.users {
    users[i] = newUser(user.ID, user.Name)
    i++
  }
  //return u.users, nil
  return users, nil
}

func (u *userRepository) Create(name string) (*model.User, error) {
  id, err := uuid.NewUUID()
  if err != nil {
    return nil, err
  }

  user := &model.User{
    ID:   id,
    Name: name,
  }

  //u.users = append(u.users, user)
  return user, nil
}

/*
func (u *userRepository) Update(user model.User) (*model.User, error) {
  return nil, nil
}

func (u *userRepository) Delete(id uuid.UUID) (error) {
  result := []*domain.User{}
	for _, ul := range u.userList {
		if ul.ID != id {
      result = append(result, ul)
    }
  }
  return nil
}
*/
