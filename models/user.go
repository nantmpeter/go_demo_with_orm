package models

import (
	"errors"
	_ "strconv"
	_ "time"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	//u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	//UserList["user_11111"] = &u
}

func AddUser(u User) *User {
	user := new(User)
	user.Username = u.Username
	user.Password = u.Password
	profile := new(Profile)
	if u.Profile != nil {
        profile.Age = u.Profile.Age
        profile.Gender = u.Profile.Gender
        profile.Email = u.Profile.Email
        profile.Address = u.Profile.Address
	}
	user.Profile = profile
	o := orm.NewOrm()
	o.Insert(profile)
	id, err := o.Insert(user)
	result := string(id)
	if err != nil {
		result += err.Error()
	}
	//UserList[id] = &u
	return user
	//return result
}

func GetUser(uid int) (u User, err error) {
	user := User{Id:int(uid)}
	orm.NewOrm().Read(&user)
	profile := &Profile{Id:user.Profile.Id}
	orm.NewOrm().Read(profile)
	user.Profile = profile
	return user, nil
	//if u, ok := UserList[uid]; ok {
	//	return u, nil
	//}
	//return nil, errors.New("User not exists")
}

func GetAllUsers() []*User {
	qs := orm.NewOrm().QueryTable("user")
	var users []*User
	qs.All(&users)
	return users
}

func UpdateUser(uid int, uu *User) (a User, err error) {
	u := User{Id:uid}
	err = orm.NewOrm().Read(&u)
	if err != orm.ErrNoRows {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		//profileId := int(u.Profile.Id)
		profile := &Profile{Id:1}
		orm.NewOrm().Read(profile)
		if uu.Profile != nil {
			if uu.Profile.Age != 0 {
				profile.Age = uu.Profile.Age
			}
			if uu.Profile.Address != "" {
				profile.Address = uu.Profile.Address
			}
			if uu.Profile.Gender != "" {
				profile.Gender = uu.Profile.Gender
			}
			if uu.Profile.Email != "" {
				profile.Email = uu.Profile.Email
			}
			orm.NewOrm().Update(profile)
		}
		u.Profile = profile
		_, err := orm.NewOrm().Update(&u)
		if err != nil {
			return u, nil
		}
		return u, nil
	}
	return u, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
