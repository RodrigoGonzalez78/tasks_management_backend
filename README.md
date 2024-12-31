

# API REST para Gestión de Tareas

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

| Código | Descripción                                            |
|--------|--------------------------------------------------------|
| 201    |  Usuario creado exitosamente. |
| 400    |  Error en el formato del correo, la contraseña no cumple los criterios, o el correo ya está en uso. |
| 500    | Error interno al crear el usuario. |


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

| Código | Descripción                                            |
|--------|--------------------------------------------------------|
| 400    | Error en el formato de la solicitud o usuario/contraseña incorrectos. |
| 500    | Error interno al generar el token de autenticación. |


---

### 1. Eliminar cuenta de usuario

- **Endpoint:** `/delete-account`
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

- **Endpoint:** `/update-user`
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
  "ID": 1,
  "CreatedAt": "2024-12-27T00:37:17.859405-03:00",
  "UpdatedAt": "2024-12-27T00:37:17.859405-03:00",
  "DeletedAt": null,
  "first_name": "Juan",
  "last_name": "Pérez",
  "email": "algo@gmail.com",
  "password": ""
}
```

### Actualizar Contraseña del Usuario

- **Endpoint:** `/update-password`
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

---

### Obtener Todas las Tareas del Usuario

- **Endpoint:** `/tasks`
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
    "ID": 1,
    "CreatedAt": "2024-12-28T00:41:07.98725-03:00",
    "UpdatedAt": "2024-12-28T00:41:07.98725-03:00",
    "DeletedAt": null,
    "title": "App Fix 1",
    "description": "Is boring",
    "done": false,
    "user_id": 1
  },
  {
    "ID": 2,
    "CreatedAt": "2024-12-28T00:42:00.127492-03:00",
    "UpdatedAt": "2024-12-28T00:42:00.127492-03:00",
    "DeletedAt": null,
    "title": "App Fix 2",
    "description": "Is boring",
    "done": false,
    "user_id": 1
  }
]
```

---

### Obtener una Tarea Específica

- **Endpoint:** `/tasks/{id}`
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
    "ID": 1,
    "CreatedAt": "2024-12-28T00:41:07.98725-03:00",
    "UpdatedAt": "2024-12-28T00:41:07.98725-03:00",
    "DeletedAt": null,
    "title": "App Fix 1",
    "description": "Is boring",
    "done": false,
    "user_id": 1
}
```

---

### Crear una Nueva Tarea

- **Endpoint:** `/tasks`
- **Método:** `POST`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Este endpoint permite crear una nueva tarea asociada al usuario autenticado.

#### Cuerpo de la Solicitud:
```json
{
  "title": "string",
  "description": "string"
}
```

- **title:** Título de la tarea (obligatorio).
- **description:** Descripción de la tarea (obligatorio).

#### Respuestas:

| Código | Descripción                                               |
|--------|-----------------------------------------------------------|
| 200    | Tarea creada exitosamente.                                |
| 400    | Error en la solicitud (formato inválido o campos vacíos). |

#### Ejemplo de Cuerpo de la Solicitud:
```json
{
  "title": "Comprar alimentos",
  "description": "Comprar frutas, verduras y leche"
}
```

#### Ejemplo de Respuesta:
```json
{
    "ID": 1,
    "CreatedAt": "2024-12-28T00:41:07.98725-03:00",
    "UpdatedAt": "2024-12-28T00:41:07.98725-03:00",
    "DeletedAt": null,
    "title": "Comprar alimentos",
    "description": "Comprar frutas, verduras y leche",
    "done": false,
    "user_id": 6
}
```

---

### Eliminar una Tarea

- **Endpoint:** `/tasks/{id}`
- **Método:** `DELETE`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Elimina una tarea específica perteneciente al usuario autenticado.

#### Parámetros de la URL:
- **id:** ID de la tarea a eliminar.

#### Respuestas:

| Código | Descripción                                 |
|--------|---------------------------------------------|
| 200    | Tarea eliminada exitosamente.              |
| 404    | La tarea especificada no fue encontrada.   |

#### Ejemplo de Respuesta:
```json
{
  "message": "Tarea eliminada exitosamente"
}
```

---

### Actualizar una Tarea

- **Endpoint:** `/tasks/{id}`
- **Método:** `PUT`
- **Middleware:** `CheckJwt` (Requiere autenticación JWT)

#### Descripción:
Permite actualizar los detalles de una tarea específica.

#### Parámetros de la URL:
- **id:** ID de la tarea a actualizar.

#### Cuerpo de la Solicitud:
```json
{
  "title": "string",
  "description": "string"
}
```

- **title:** Nuevo título de la tarea (obligatorio).
- **description:** Nueva descripción de la tarea (obligatorio).

#### Respuestas:

| Código | Descripción                                               |
|--------|-----------------------------------------------------------|
| 200    | Tarea actualizada exitosamente.                           |
| 400    | Error en la solicitud (formato inválido o campos vacíos). |
| 404    | La tarea especificada no fue encontrada.                  |
| 500    | Error al guardar la tarea en la base de datos.            |

#### Ejemplo de Cuerpo de la Solicitud:
```json
{
  "title": "Actualizar informe",
  "description": "Completar la sección de análisis de datos"
}
```

#### Ejemplo de Respuesta:
```json
{
    "ID": 1,
    "CreatedAt": "2024-12-28T00:41:07.98725-03:00",
    "UpdatedAt": "2024-12-28T00:41:07.98725-03:00",
    "DeletedAt": null,
    "title": "Actualizar informe",
    "description": "Completar la sección de análisis de datos",
    "done": false,
    "user_id": 1
}
```
