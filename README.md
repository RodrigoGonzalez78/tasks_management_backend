

# API REST para Gestión de Tareas/Notas

Esta API permite a los usuarios registrarse, iniciar sesión y gestionar sus tareas y notas. A continuación, se describen los endpoints disponibles y los formatos de solicitud y respuesta.

## Endpoints

### 1. Registro de Usuario

**Endpoint:** `/signup`  
**Método:** `POST`  
**Descripción:** Crea una nueva cuenta de usuario.

**Formato de solicitud:**
```json
{
  "email": "ejemplo@gmail.com",
  "password": "123456789"
}
```

**Requisitos:**
- El campo `email` debe ser un correo electrónico válido.
- La `password` debe tener al menos 8 caracteres.

**Respuestas:**

- **201 Created:** Usuario creado exitosamente.
- **400 Bad Request:** Error en el formato del correo, la contraseña no cumple los criterios, o el correo ya está en uso.
- **500 Internal Server Error:** Error interno al crear el usuario.

### 2. Inicio de Sesión

**Endpoint:** `/login`  
**Método:** `POST`  
**Descripción:** Inicia sesión en la aplicación con un correo electrónico y contraseña válidos.

**Formato de solicitud:**
```json
{
  "email": "ejemplo@gmail.com",
  "password": "123456789"
}
```

**Respuestas:**

- **200 OK:** Inicio de sesión exitoso.
  
  **Formato de respuesta:**
  ```json
  {
    "token": "djskdjskdjskjdksjjdksdksj..."
  }
  ```

- **400 Bad Request:** Error en el formato de la solicitud o usuario/contraseña incorrectos.
- **500 Internal Server Error:** Error interno al generar el token de autenticación.

---

### Ejemplo de Uso

#### Registro
```bash
curl -X POST http://tudominio.com/signup -H "Content-Type: application/json" -d '{
  "email": "ejemplo@gmail.com",
  "password": "123456789"
}'
```

#### Inicio de Sesión
```bash
curl -X POST http://tudominio.com/login -H "Content-Type: application/json" -d '{
  "email": "ejemplo@gmail.com",
  "password": "123456789"
}'
```


### 1. Eliminar cuenta de usuario

- **URL:** `/delete-account`
- **Método:** `DELETE`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Este endpoint elimina la cuenta de un usuario autenticado, incluyendo todas las tareas asociadas al usuario.

#### Cuerpo de la solicitud:
No se requiere un cuerpo para esta solicitud.

#### Respuestas:

| Código | Descripción                                            |
|--------|--------------------------------------------------------|
| 200    | La cuenta del usuario y todas sus tareas fueron eliminadas exitosamente. |
| 404    | El usuario no fue encontrado en la base de datos.      |

#### Ejemplo de Respuesta:
```json
{
  "message": "Cuenta eliminada exitosamente."
}
```



### Actualizar datos del usuario

- **URL:** `/update-user`
- **Método:** `PUT`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Permite a un usuario autenticado actualizar sus datos personales, como su nombre y apellido.

#### Cuerpo de la solicitud:

```json
{
  "first_name": "string",
  "last_name": "string"
}
```

- **first_name:** Nombre del usuario (obligatorio).
- **last_name:** Apellido del usuario (obligatorio).

#### Respuestas:

| Código | Descripción                                            |
|--------|--------------------------------------------------------|
| 200    | Los datos del usuario fueron actualizados exitosamente. |
| 400    | Faltan campos obligatorios o hay un error en la solicitud. |
| 404    | El usuario no fue encontrado en la base de datos.      |
| 500    | Error interno al actualizar los datos en la base de datos. |

#### Ejemplo de Cuerpo de la Solicitud:
```json
{
  "first_name": "Juan",
  "last_name": "Pérez"
}
```

#### Ejemplo de Respuesta:
```json
{
  "id": 1,
  "first_name": "Juan",
  "last_name": "Pérez",
  "email": "juan.perez@example.com",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-12-03T14:00:00Z"
}
```

#### Lógica del Endpoint:
1. Decodifica el cuerpo de la solicitud para obtener los datos del usuario.
2. Verifica que los campos `first_name` y `last_name` no estén vacíos.
3. Busca al usuario autenticado en la base de datos usando `jwtMetods.IDUser`.
4. Si el usuario no existe, responde con un `404 Not Found`.
5. Actualiza los campos del usuario con los nuevos valores proporcionados.
6. Guarda los cambios en la base de datos.
7. Devuelve el usuario actualizado en caso de éxito (`200 OK`).

### Actualizar Contraseña del Usuario

- **URL:** `/update-password`
- **Método:** `PUT`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Este endpoint permite a un usuario autenticado actualizar su contraseña de forma segura.

#### Cuerpo de la solicitud:

```json
{
  "password": "string"
}
```

- **password:** Nueva contraseña del usuario (obligatoria).

#### Respuestas:

| Código | Descripción                                             |
|--------|---------------------------------------------------------|
| 200    | La contraseña fue actualizada exitosamente.             |
| 400    | Error en la solicitud (por ejemplo, el formato es inválido). |
| 404    | El usuario no fue encontrado en la base de datos.       |
| 500    | Error interno al generar el hash o actualizar la base de datos. |

#### Ejemplo de Cuerpo de la Solicitud:
```json
{
  "password": "nuevaContraseñaSegura123"
}
```

#### Ejemplo de Respuesta:
```json
{
  "message": "Contraseña actualizada correctamente"
}
```

#### Lógica del Endpoint:
1. Decodifica el cuerpo de la solicitud para obtener la nueva contraseña.
2. Verifica que el usuario autenticado existe en la base de datos usando `jwtMetods.IDUser`.
3. Genera un hash seguro para la nueva contraseña utilizando `bcrypt`.
4. Actualiza la contraseña en la base de datos.
5. Devuelve un mensaje de éxito si la operación es exitosa (`200 OK`).

---

### Obtener Todas las Tareas del Usuario

- **URL:** `/tasks`
- **Método:** `GET`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Devuelve una lista de todas las tareas asociadas al usuario autenticado.

#### Cuerpo de la solicitud:
No se requiere un cuerpo para esta solicitud.

#### Respuestas:

| Código | Descripción                                           |
|--------|-------------------------------------------------------|
| 200    | Lista de tareas recuperada exitosamente.              |
| 404    | No se encontraron tareas para el usuario autenticado. |

#### Ejemplo de Respuesta:
```json
[
  {
    "id": 1,
    "title": "Tarea 1",
    "description": "Descripción de la tarea 1",
    "done": false,
    "user_id": 1,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
]
```

---

### Obtener una Tarea Específica

- **URL:** `/tasks/{id}`
- **Método:** `GET`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Devuelve los detalles de una tarea específica perteneciente al usuario autenticado.

#### Parámetros de la URL:
- **id:** ID de la tarea a recuperar.

#### Cuerpo de la solicitud:
No se requiere un cuerpo para esta solicitud.

#### Respuestas:

| Código | Descripción                                            |
|--------|--------------------------------------------------------|
| 200    | Detalles de la tarea recuperados exitosamente.         |
| 404    | La tarea especificada no fue encontrada en la base de datos. |

#### Ejemplo de Respuesta:
```json
{
  "id": 1,
  "title": "Tarea 1",
  "description": "Descripción de la tarea 1",
  "done": false,
  "user_id": 1,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

#### Lógica del Endpoint:
1. Extrae el parámetro `id` de la URL.
2. Busca la tarea en la base de datos utilizando `db.GetTaskById`.
3. Si no se encuentra, responde con un `404 Not Found`.
4. Devuelve los detalles de la tarea si es encontrada (`200 OK`).

---



### Notas Adicionales

- Todos los endpoints de la API requieren que las solicitudes estén formateadas en JSON.
- Las respuestas exitosas están acompañadas de un token JWT, que se debe usar para autenticar solicitudes posteriores a otros endpoints de la API.

