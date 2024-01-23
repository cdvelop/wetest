package wetest

func (h WeTest) DB_Insert_Data_Table(this string, t TestAction) (err string) {
	return h.CreateObjectsInDB(t.DB_Insert_Data_Table, false, t.Data)
}

func (h WeTest) DB_Delete_Data_Table(this string, t TestAction) (err string) {
	return h.DeleteObjectsInDB(t.DB_Delete_Data_Table, false, t.Data...)
}
