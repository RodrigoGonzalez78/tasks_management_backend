

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

### Notas Adicionales

- Todos los endpoints de la API requieren que las solicitudes estén formateadas en JSON.
- Las respuestas exitosas están acompañadas de un token JWT, que se debe usar para autenticar solicitudes posteriores a otros endpoints de la API.

--- 

Esta estructura mejora la legibilidad y proporciona una visión clara de los endpoints y sus usos, lo cual es útil para cualquier desarrollador que quiera interactuar con tu API.
