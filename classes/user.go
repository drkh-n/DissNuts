type UserObject struct{
  uin string
  name string
  surname string
  middlename string
  bloodtype string
  phone string
  email string
  emergencyContact string
  address string
  maritalStatus string
  birthdate string
  regDate string
}
// method returns true if user has no empty essential fields
func (u UserObject) isValidUser() bool{
  return u.uin.isEmpty() && u.name.isEmpty() && u.surname.isEmpty() && (u.phone.isEmpty() || u.email.isEmpty()) && u.birthdate.isEmpty()
}
