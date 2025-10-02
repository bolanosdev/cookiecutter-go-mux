package queries

import "fmt"

var GET_ACCOUNTS_QUERY = `
  select 
    a.id, 
    a.email, 
    a.password,
    a.role_id,
    r.name as role_name,
    a.is_active,
    a.created_at,
    a.updated_at
  from accounts a
  join roles r on r.id = a.role_id`

var GET_ACCOUNTS_BY_ID_QUERY = fmt.Sprintf(`
  %s 
  where a.id = $1`, GET_ACCOUNTS_QUERY)

var GET_ACCOUNTS_BY_EMAIL_QUERY = fmt.Sprintf(`
  %s 
  where a.email = $1`, GET_ACCOUNTS_QUERY)
