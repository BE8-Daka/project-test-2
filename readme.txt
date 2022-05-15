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

[JWT][GET] {url:8000}/tasks

========================================

[JWT][PUT] {url:8000}/tasks/{id}

{
    "name" : "project updated"
}

========================================

[JWT][POST] {url:8000}/tasks/{id}/completed

========================================

[JWT][POST] {url:8000}/tasks/{id}/reopen

========================================

[JWT][POST] {url:8000}/projects

{
    "name" : "project pertama"
}

========================================

[JWT][GET] {url:8000}/projects

========================================

[JWT][PUT] {url:8000}/projects/{id}

{
    "name" : "project updated"
}

========================================

[JWT][GET] {url:8000}/projects/{id}/tasks

========================================

[JWT][DELETE] {url:8000}/projects

========================================