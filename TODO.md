# TODOs

[x] No utilizar más go4it
[x] reemplazar el uso de struct app
[x] eliminar webcore y webcore_features: Ahora todas las librerías son tango_* lo que hace que sea más sencillo su customización.
[x] Usar librería de base de datos

[x] Tango ahora permite usar templ con htmlx
[x] En las opciones de configuracion debe estar si la app tiene para utilizar el sistema de plantillas
[] El cli debe poder crear vistas para esto
[x] Tango tiene un tipo de estructura que sirve para exportar datos a las vistas: ViewData
[x] Ejemplo de uso de vistas con template


CLI:
[x] Mejorar y limpiar el CLI
[x] Revisar y reescribir la parte de TEMPLATES. Quizás reemplazo por sistemas de submódulos ej: views, models, features, extras.
[x] Reemplazar el gomake por por el uso del CLI
[] Translate cli to english
[] Agregar creación de páginas o vistas por el CLI
[] Ahora puede construir los builds con nombre de app y version para varias plataformas (en proceso)
[] realiza las migraciones de los datos
[] realiza la carga de los seeds
[] puede revertir los estados de la migración (utiliza una tabla de sqlite)
[x] cli tiene modo de desarrollo
[] el modo de desarrollo debe poder levantar el frontend y el backend y mostrar que están corriendo (goroutines?)
[] El cli debe poner a la app en modo desarrollo para poder utilizar los templates
[] El cli debe poder reenderizar los templates (en proceso)
[] El cli debe poder crear templates


[] Comenzar a compartir la app
[] AUTH realizar sistema de autenticación
