package dao

func ReturnPermission(name string) (string,error) {
	var permission string

	sqlStr := "select accessLevel from teacher where name = ?"
	rows,err := DB.Query(sqlStr,name)
	for rows.Next() {
		err = rows.Scan(&permission)
		if err != nil {
			return "",err
		}
	}
	return permission,nil
}