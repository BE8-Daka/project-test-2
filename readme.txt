[POST] {url:8000}/users/register

{
    "name" : "testing",
    "username" : "testing",
    "no_hp" : "081234567890",
    "email" : "testing@gmail.com",
    "password" : "testing"
}
========================================
[POST] {url:8000}/users/login

{
    "username" : "testing",
    "password" : "testing"
}
========================================
[JWT][GET] {url:8000}/users/profile
========================================
[JWT][PUT] {url:8000}/users/profile

{
    "name" : "testing updated"
}
========================================