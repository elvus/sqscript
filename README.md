# sonarscript

Script for sonar

## Getting started
#### Instalacion
#### Windows
Descargar y mover el archivo *sq.exe* en la carperta bin del sonar donde se encuentra el archivo sonar-scanner.bat
Luego en las variables de entorno del windows buscar en variables del sistema la variable Path y agregar {path_del_sonar}\bin
## Mac/Linux
Descargar y mover el archivo *sq* en la carperta bin del sonar donde se encuentra el archivo sonar-scanner
Luego agregar al Path de las variables de entorno del sistema {path_del_sonar}\bin
### Configuracion
Ejectuar el comando sq y agregar la url y el shh key

## Uso
Una vez agregada la carpeta del sonar en los enviroment variables podremos ejecutar el sgte comando desde donde nos parezca conveniente
```
sq o sq {path_del_archivo_a_testear} {nombre_del_archivo_para_el_id}
Ejemplo: sq c:\sonar_proxy_v1 sonar_proxy_v1
```
Obs: el comando sq toma por defecto el nombre de la carpeta y el path donde estamos ubicados
