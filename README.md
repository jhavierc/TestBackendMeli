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
 
## Implementación de la solución
Para el desarrollo de la siguiente prueba, se utilizó las siguientes tecnologías:

- Golang vgo1.15.2
- Docker 20.10.2
- AWS Cloud (Creación de infraestructura base con codigo usando CloudFormation)
- Docker Hub

## Arquitectura Cloud

![](https://matthcep.s3.amazonaws.com/mercadolibre_architecture.png)

Para este proyecto se seleccionó a AWS como proveedor cloud para despliegue del microservicio implemetado.
Una vez cargado la imagen del proyecto en el ECR (Amazon Elastic Container Registry), esta queda disponible para que la tarea configurada en el ECS (Elastic Container Service) realice el despliegue del servicio.
Para realizar el consumo del servicio se cuentan con tres opciones, eso depende de la implementación y la configuración misma de los servicios de AWS, pero para la prueba el cliente puede acceder al servicio utilizando el DNS del balanceador de carga.
La arquitectura propuesta según el Well-Architected Framework de AWS esta permite una excelencia operativa, seguridad, fiabilidad, eficiencia de rendimiento y optimización de costos debido a que es posible escalar horizontalmente de forma automatica dependiendo de la cantidad de carga operativa que se detecte en los microservicios.
