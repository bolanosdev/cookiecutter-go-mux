package queries

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

var CREATE_ACCOUNT_QUERY = `
	insert into accounts (email, password) 
	values (@email, @password) 
	returning *;
`
