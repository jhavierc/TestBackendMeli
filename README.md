# Operación Fuego de Quasar

Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un gran golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.
El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que transporta raciones y armamento para una legión entera.

![](https://matthcep.s3.amazonaws.com/Screen+Shot+2021-01-24+at+6.41.27+PM.png)

Ejecutar aplicación local
=============
Para ejecutar la aplicación en local, seguir los siguientes pasos:

Requerimientos:

* Tener instalado <abbr title="Docker">Docker</abbr>

Clonar el repositorio:

`$ git clone https://github.com/jhavierc/TestBackendMeli.git`

`$ cd TestBackendMeli`

Una vez dentro de la carpeta del proyecto ejecutar los siguientes comandos Docker:

`$ docker build . -t mercadolibre`

`$ docker run -p 4000:4000 mercadolibre`
 
Implementación de la solución
=============
Para el desarrollo de la siguiente prueba, se utilizó las siguientes tecnologías:

- Golang vgo1.15.2
- Docker 20.10.2
- AWS Cloud (Creación de infraestructura base con codigo usando CloudFormation)
- Docker Hub
