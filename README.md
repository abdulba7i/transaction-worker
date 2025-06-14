# 💸 RabbitMQ Transfer System

Тестовое задание: демонстрационный проект на Go с использованием RabbitMQ и PostgreSQL для обработки переводов средств между пользователями.

---

## 📌 Задача

Реализовать два сервиса:

- **Сервис A** — принимает HTTP-запросы на перевод средств и отправляет события в RabbitMQ.
- **Сервис B** — читает сообщения из RabbitMQ, проверяет баланс, создаёт заявки и обрабатывает их статус через 30 секунд.

---

## 🔧 Функциональность

### 🌐 Сервис A

- Принимает HTTP-запросы на перевод:
  - `user_id` — ID пользователя
  - `request_id` — ID заявки
  - `amount` — сумма
- Проверяет уникальность `request_id` для каждого `user_id`
- Ограничивает частоту запросов: не более **10 запросов в секунду** от одного пользователя
- Отправляет сообщение в RabbitMQ

### ⚙️ Сервис B

- Слушает сообщения из RabbitMQ
- Проверяет наличие средств у пользователя
- Создаёт заявку в PostgreSQL и списывает средства
- Через 30 секунд:
  - Случайным образом выбирается статус: `успех` или `неуспех`
  - В случае `успеха` — средства списаны окончательно
  - В случае `неуспеха` — средства возвращаются на баланс

---

## 🛠 Технологии и требования

| Компонент             | Использовано                        |
| --------------------- | ----------------------------------- |
| Язык программирования | Go                                  |
| Брокер сообщений      | RabbitMQ                               |
| База данных           | PostgreSQL                          |
| HTTP-сервер           | `net/http`                          |
| Хранилище заявок      | PostgreSQL                          |
| Обработка очередей    | RabbitMQ Consumer/Producer             |
| Ограничение запросов  | Rate Limiter на пользователя        |
| Тестирование          | Unit-тесты                          |
| Защита данных         | Идемпотентность и валидация         |
| Сборка                | Docker (опционально)                |
| Окружение             | Docker Compose (опционально)       |

---

## 📄 Пример запроса к сервису A

```http
POST /transfer HTTP/1.1
Content-Type: application/json

{
  "user_id": "123",
  "request_id": "req-456",
  "amount": 100
}
```