package queries

import "fmt"

var GET_ALL_ROLES_QUERY = `
 select 
    r.id,
    r.name ,
    (
      select array_agg(row(p.id, p.name)) 
      from role_permissions as rp 
      join permissions as p on rp.permission_id = p.id
      where rp.role_id = r.id 
    )
  from roles as r
`

var GET_ROLES_BY_ID_QUERY = fmt.Sprintf(`
  %s
  where r.id = $1
`, GET_ALL_ROLES_QUERY)
