# Project Test 2 - Todo List

Salah satu tugas individu untuk menyelesaikan kursus di alterra.id (Alterra Academy)

## Overview

    - Fitur Register dan Login
    - Fitur Task (Membuat todo)
    - Fitur Project (mengelompokkan task yang sudah dibuat)

## Endpoint

    [POST] {url:8080}/users/register

    {
        "name" : "testing",
        "username" : "testing",
        "no_hp" : "081234567890",
        "email" : "testing@gmail.com",
        "password" : "testing"
    }

    ========================================

    [POST] {url:8080}/users/login

    {
        "username" : "testing",
        "password" : "testing"
    }

    ========================================

    [JWT][GET] {url:8080}/users/profile

    ========================================

    [JWT][PUT] {url:8080}/users/profile

    {
        "name" : "testing updated"
    }

    ========================================

    [JWT][DELETE] {url:8080}/users/profile

    ========================================

    [JWT][POST] {url:8080}/tasks

    {
        "name" : "project pertama"
        "project_id" : 1 (opsional)
    }

    ========================================

    [JWT][GET] {url:8080}/tasks

    ========================================

    [JWT][PUT] {url:8080}/tasks/{id}

    {
        "name" : "project updated"
    }

    ========================================

    [JWT][POST] {url:8080}/tasks/{id}/completed

    ========================================

    [JWT][POST] {url:8080}/tasks/{id}/reopen

    ========================================

    [JWT][POST] {url:8080}/projects

    {
        "name" : "project pertama"
    }

    ========================================

    [JWT][GET] {url:8080}/projects

    ========================================

    [JWT][PUT] {url:8080}/projects/{id}

    {
        "name" : "project updated"
    }

    ========================================

    [JWT][GET] {url:8080}/projects/{id}/tasks

    ========================================

    [JWT][DELETE] {url:8080}/projects

    ========================================

## Contributing

    - Mahmuda Karima (DAKA)

## Copyrights

    - Mei 2022
