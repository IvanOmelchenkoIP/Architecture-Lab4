# Лабораторна робота 4 з Архітектури Програмного Забезпечення

## Вибір завдання за варіантом

Варіант згідно команди */team-variant 2* - 3.

![variant](https://github.com/IvanOmelchenkoIP/Architecture-Lab4/blob/main/markdown_files/variant.PNG)

Завдання згідно варіанту - як другу команду, імплементувати команду:
    
    reverse <arg>

Команда має перевертати вхідний рядок.

## Залежності

[Посилання на gocheck](https://github.com/go-check/check.git)

## Запуск програми

Для запуску програми можна скористатися командою:

    go run main.go <filename>

Де *filename* - ім'я файлу, з якого потрібно прочитати команди.

## Запуск тестів

Запуск тестів парсера:

    go test -v ./parser

Запуск тестів цикла подій:

    go test -v ./engine
