package events

const UserRegisteredEventType = "user.registered"

type UserRegistered struct {
	BaseEvent
	UserID uint
	Email  string
}

func NewUserRegistered(userID uint, email string) UserRegistered {
	return UserRegistered{
		BaseEvent: NewBaseEvent(UserRegisteredEventType),
		UserID:    userID,
		Email:     email,
	}
}
