package consts

const QueryInit = `{"selector":{"type":"initialdata"}}`

func BodySecurity(dbName string) string {
	return `{"admins": {"names": ["` + dbName + `"],"roles": ["admin_role"]},"members": {"names": [],"roles": []}}`
}
func QueryCompanyAlias(alias string) string {
	return `{"selector": {"table":"company","alias":"` + alias + `"},"limit":1}`
}
