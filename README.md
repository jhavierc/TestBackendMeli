# Operación Fuego de Quasar

![](https://raw.githubusercontent.com/jhavierc/TestBackendMeli/master/Logo_ML.png?token=AB3AZQUAMMXOAI6S5YI66VLABYALK)

![](https://img.shields.io/github/stars/pandao/editor.md.svg) 

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
