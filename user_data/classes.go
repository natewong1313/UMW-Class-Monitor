package user_data

func CheckClassInConfig(newClass Class, userData *UserData) bool{
	for _, class := range userData.Classes{
		if class.CRN == newClass.CRN{
			return true
		}
	}
	return false
}

func AddNewClassToConfig(class Class, userData *UserData) {
	userData.Classes = append(userData.Classes, class)

	UpdateConfigData(userData)
}

func RemoveClassFromConfig(CRN string, userData *UserData) {
	var newClasses []Class
	for _, class := range userData.Classes{
		if class.CRN != CRN{
			newClasses = append(newClasses, class)
		}
	}
	userData.Classes = newClasses

	UpdateConfigData(userData)
}