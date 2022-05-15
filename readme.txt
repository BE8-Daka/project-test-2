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

[JWT][DELETE] {url:8000}/users/profile

========================================

[JWT][POST] {url:8000}/tasks

{
    "name" : "project pertama"
    "project_id" : 1 (opsional)
}

========================================

[JWT][POST] {url:8000}/projects

{
    "name" : "project pertama"
}

========================================

[JWT][PUT] {url:8000}/projects

{
    "name" : "project updated"
}

========================================

[JWT][GET] {url:8000}/projects

========================================