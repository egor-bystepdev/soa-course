# Документация GraphQL

## Типы

**Player**: Игрок в игре.
- `username: String!` - Уникальное имя игрока.
- `role: String!` - Роль игрока в игре.
- `alive: Boolean!` - Жив ли игрок.

**Comment**: Комментарий к игре.
- `source: String!` - Источник комментария.
- `text: String!` - Текст комментария.
- `date: String!` - Дата комментария.

**GameState**: Состояние игры (может быть в процессе или окончено).

**Game**: Игра.
- `id: ID!` - Уникальный идентификатор игры.
- `state: GameState!` - Текущее состояние игры.
- `date: String!` - Дата начала игры.
- `players: [Player!]!` - Список игроков в игре.
- `comments: [Comment!]!` - Список комментариев к игре.

## Входные типы

**GameInfo**: Информация об игре при ее создании.
- `id: ID!` - Уникальный идентификатор для новой игры.
- `state: GameState!` - Начальное состояние игры.

**PlayerStatus**: Состояние игрока в игре.
- `username: String!` - Имя игрока.
- `alive: Boolean!` - Жив ли игрок.

**PlayerData**: Информация об игроке при его добавлении в игру.
- `username: String!` - Имя игрока.
- `role: String!` - Роль игрока.

**CommentData**: Информация о комментарии при его добавлении к игре.
- `source: String!` - Источник комментария.
- `text: String!` - Текст комментария.

## Мутации

**CreateGame(game_info: GameInfo!, players: [PlayerData!]!): Boolean!** - Создает новую игру.

**UpdateGame(game_id: ID!, new_state: GameState!, player_statuses: [PlayerStatus!]!): Boolean!** - Обновляет состояние игры и статус игроков.

**AddComment(game_id: ID!, data: CommentData!): Boolean!** - Добавляет комментарий к игре.

## Запросы

**GetGamesInfo(ids: [ID!]!): [Game!]!** - Возвращает информацию о специфичных играх.

**ListGames(max_count: Int!, state: GameState!): [ID!]!** - Возвращает список игр в определенном состоянии.

## Примеры запросов

#### Создание игры:
```graphql
mutation {
  CreateGame(
    game_info: { id: "1", state: IN_PROGRESS }, 
    players: [
      { username: "Player1", role: "Survivor" },
      { username: "Player2", role: "Zombie" }
    ]
  )
}
```

#### Обновление игры:
```graphql
mutation {
  UpdateGame(
   

 game_id: "1", 
    new_state: END, 
    player_statuses: [
      { username: "Player1", alive: false },
      { username: "Player2", alive: true }
    ]
  )
}
```

#### Добавление комментария:
```graphql
mutation {
  AddComment(
    game_id: "1", 
    data: { source: "Player1", text: "This game is intense!" }
  )
}
```

#### Получение информации об играх:
```graphql
query {
  GetGamesInfo(ids: ["1", "2"])
}
```

#### Получение списка игр:
```graphql
query {
  ListGames(max_count: 10, state: IN_PROGRESS)
}
```
