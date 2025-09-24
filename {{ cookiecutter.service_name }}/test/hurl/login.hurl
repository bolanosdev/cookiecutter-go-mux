
POST {{API_URL}}/login
Content-Type: application/json
{
  "email": "test@email.com",
  "password": "123456"
}
HTTP 200
[Captures]
access_token: header "access-token"


