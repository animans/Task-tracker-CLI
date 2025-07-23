# Task Tracker CLI

📌 **Task-tracker-CLI** — это консольное приложение для управления задачами, написанное на Go. Простой и удобный способ вести список дел из терминала.

---

## 🚀 Возможности

- Добавление новых задач
- Просмотр всех текущих задач
- Отметка задачи как выполненной
- Удаление задачи
- Сохранение задач в файл

---

## 📦 Установка

### 1. Клонирование репозитория

```bash
git clone https://github.com/animans/Task-tracker-CLI.git
cd Task-tracker-CLI
```

### 2. Убедитесь, что Go установлен

```bash
go version
```

Если Go не установлен, скачайте его с [официального сайта](https://golang.org/dl).

### 3. Сборка проекта

```bash
go build -o task-tracker
```

После сборки появится исполняемый файл `task-tracker`.

---

## 🧪 Как использовать проект
```bash
#Добавление новой задачи
./task-tracker add "Название задачи"

#Обавление и удаление задачи
./task-tracker update 1 "Новое название задачи"
./task-tracker delete 1

#Отметить задачу как in progress или done
./task-tracker mark-in-progress 1
./task-tracker mark-done 1

#Просмотр всех задач
./task-tracker list

#Просмотр задач по статусу
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```
