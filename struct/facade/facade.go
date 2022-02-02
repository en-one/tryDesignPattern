package facade

/*

 */

type IUser interface {
	Login(phone int, code int) (*User, error)
	Register(phone int, code int) (*User, error)
}

type User struct {
	Name string
}

// IUserFacade 简化流程， 门面模式
type IUserFacade interface {
	LoginOrRegister(phone int, code int) error
}

type UserService struct{}

// Login 登录
//func (*User) Login(phone int, code int) (*User, error) {
func (*UserService) Login(phone int, code int) (*User, error) {
	// 校验操作
	return &User{Name: "test login"}, nil
}

// Register 创建
//func (*User) Register(phone int, code int) (*User, error) {
func (*UserService) Register(phone int, code int) (*User, error) {
	// 校验操作
	// 创建用户
	return &User{Name: "test register"}, nil
}

// LoginOrRegister 登录或注册
func (u UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.Register(phone, code)
}
