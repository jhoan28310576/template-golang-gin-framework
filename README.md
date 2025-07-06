# template-golang-gin framework
template goland con estructura de carpetas listo para adactar a un comienzo de cualquier app

El framework Gin se ha instalado correctamente y he creado un servidor web de ejemplo. Aquí está el resumen de lo que se ha configurado:

✅ Gin Framework Instalado
Dependencias agregadas al go.mod:
github.com/gin-gonic/gin v1.10.1 - Framework principal
Todas las dependencias necesarias automáticamente

�� Servidor Web de Ejemplo
He creado un servidor web completo en cmd/webserver/main.go con las siguientes funcionalidades:
Rutas API disponibles:
GET / - Mensaje de bienvenida
GET /api/info - Información del proyecto
GET /api/saludo/:nombre - Saludo personalizado
POST /api/datos - Recibir datos JSON
GET /html - Página web interactiva

Características incluidas:
✅ Servir archivos estáticos desde /assets
✅ Renderizar templates HTML desde /templates
✅ Manejo de parámetros de URL
✅ Procesamiento de JSON
✅ Respuestas HTTP apropiadas
🎨 Página Web Interactiva
He creado templates/index.html con:
Interfaz moderna y responsive
Botones para probar todas las APIs
Manejo de errores en JavaScript
Diseño limpio y profesional

🏃‍♂️ Cómo usar:
Ejecutar el servidor:
 go run cmd/webserver/main.go
 

Acceder a las rutas:
http://localhost:8080/ - Mensaje de bienvenida
http://localhost:8080/html - Página web interactiva
http://localhost:8080/api/info - Información del proyecto
 Puedes acceder a http://localhost:8080/html para ver la página web interactiva y probar todas las funcionalidades de Gin.
