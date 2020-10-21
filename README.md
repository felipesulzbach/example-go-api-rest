# GoLang Web Service REST

API REST example in GoLang.

--------------------------------------------------------------------------------

## About Golang

![](https://raw.githubusercontent.com/felipesulzbach/grpc-go-example/master/things/go.png)

### What is?

Golang, or simply Go, is an open source language created in 2009 by [Google](https://about.google/intl/en_US/) (by engineers [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike) and [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson)). The Go language was created with the goal of having **C** language performance but also focusing more readable and easier to program from more robust languages like **Java**.

### Some advantages of language

- Incredibly light in terms of memory usage;
- Suppose several concurrent processing because it uses goroutines instead of threads that are found in most programming languages. Competition is one of the language's strengths;
- Compiles very fast;
- Has garbage collector (has been incorporated into its core in order to prioritize performance);
- It is strongly typed.

GoLang intentionally leaves out many features of modern _OOP_ languages. Everything is divided into packages. [Google](https://about.google/intl/en_US/) technology has only _structs_ instead of _classes_.

### Some companies that have adopted Golang:

- Netflix
- The Economist
- IBM
- GitHub
- Uber
- Docker
- Dropbox
- OpenShift
- Twitter
- [Complete list by country (link here)](https://github.com/golang/go/wiki/GoUsers)

## About this project

### Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/) or other IDE (Integrated Development Environment)
- [Golang](https://golang.org/)
- [Docker](https://hub.docker.com/search?offering=community&type=edition&operating_system=linux%2Cwindows)

### Preparing the environment

The following technologies are critical for running/compiling application sources:

- Access the application directory from the terminal and execute:

  - Install the [gorilla/mux](https://github.com/gorilla/mux):

    > go get -u github.com/gorilla/mux

  - Install the [Postgres Driver](https://github.com/lib/pq):

    > go get github.com/lib/pq

  - Install the [go-i18n](https://github.com/nicksnyder/go-i18n)

    > go get -u github.com/nicksnyder/go-i18n/v2/goi18n

  - Install the [Validator](https://github.com/go-playground/validator)

    > go get gopkg.in/go-playground/validator.v9

### PostgreSql Database configuration

#### Connection Settings

- HOST: **localhost**
- PORT: **5435**
- USER: **postgres**
- PASSWORD: **postgres**
- DATABASE NAME: **go_rest_db**

#### Create structure

```sql
DROP SCHEMA IF EXISTS fs_auto CASCADE;

CREATE SCHEMA fs_auto;

-- PERSON --
CREATE TABLE fs_auto.person
(id numeric NOT NULL
,name character varying(255)
,cpf character varying(20)
,cell_phone character varying(20)
,city character varying(255)
,zip_code character varying(20)
,address character varying(255)
,registration_date timestamp
,CONSTRAINT pk_person PRIMARY KEY (id));

-- COURSE --
CREATE TABLE fs_auto.course
(id numeric NOT NULL
,name character varying(255)
,description character varying(255)
,registration_date timestamp
,CONSTRAINT pk_course PRIMARY KEY (id));

-- CLASS --
CREATE TABLE fs_auto.class
(id numeric NOT NULL
,course_id numeric NOT NULL
,start_date timestamp
,end_date timestamp
,registration_date timestamp
,CONSTRAINT pk_class PRIMARY KEY (id)
,CONSTRAINT fk_class_course FOREIGN KEY (course_id) REFERENCES fs_auto.course (id));
CREATE INDEX idx_class_course ON fs_auto.class(course_id);

-- STUDENT --
CREATE TABLE fs_auto.student
(person_id numeric NOT NULL
,class_id numeric NOT NULL
,CONSTRAINT pk_student PRIMARY KEY (person_id, class_id)
,CONSTRAINT fk_student_person FOREIGN KEY (person_id) REFERENCES fs_auto.person (id)
,CONSTRAINT fk_student_class FOREIGN KEY (class_id) REFERENCES fs_auto.class (id)
);

-- TEACHER --
CREATE TABLE fs_auto.teacher
(person_id numeric NOT NULL
,course_id numeric NOT NULL
,CONSTRAINT pk_teacher PRIMARY KEY (person_id, course_id)
,CONSTRAINT fk_teacher_person FOREIGN KEY (person_id) REFERENCES fs_auto.person (id)
,CONSTRAINT fk_teacher_course FOREIGN KEY (course_id) REFERENCES fs_auto.course (id)
);
```

#### Mock registers

```sql
-- PERSON --
INSERT INTO fs_auto.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
VALUES (
    1 --id
   ,'Pessoa 1' --name
   ,01234567890 --cpf
   ,012345678 --cell_phone
   ,'Cidade 1' --city
   ,01234567 --zip_code
   ,'Endereco 1' --address
   ,now() --registration_date
);
------------
INSERT INTO fs_auto.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
VALUES (
    2 --id
   ,'Pessoa 2' --name
   ,12345678901 --cpf
   ,123456789 --cell_phone
   ,'Cidade 2' --city
   ,12345678 --zip_code
   ,'Endereco 2' --address
   ,now() --registration_date
);
------------
INSERT INTO fs_auto.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
VALUES (
    3 --id
   ,'Pessoa 3' --name
   ,23456789012 --cpf
   ,234567890 --cell_phone
   ,'Cidade 3' --city
   ,23456789 --zip_code
   ,'Endereco 3' --address
   ,now() --registration_date
);
------------
INSERT INTO fs_auto.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
VALUES (
    4 --id
   ,'Pessoa 4' --name
   ,34567890123 --cpf
   ,345678901 --cell_phone
   ,'Cidade 4' --city
   ,34567890 --zip_code
   ,'Endereco 4' --address
   ,now() --registration_date
);
------------
INSERT INTO fs_auto.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
VALUES (
    5 --id
   ,'Pessoa 5' --name
   ,45678901234 --cpf
   ,456789012 --cell_phone
   ,'Cidade 5' --city
   ,45678901 --zip_code
   ,'Endereco 5' --address
   ,now() --registration_date
);
------------
INSERT INTO fs_auto.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
VALUES (
    6 --id
   ,'Pessoa 6' --name
   ,56789012345 --cpf
   ,567890123 --cell_phone
   ,'Cidade 6' --city
   ,56789012 --zip_code
   ,'Endereco 6' --address
   ,now() --registration_date
);


-- COURSE --
INSERT INTO fs_auto.course (id,name,description,registration_date)
VALUES (
    1 --id
   ,'Nome 1' --name
   ,'Descricao 1' --description
   ,now() --registration_date
);
------------
INSERT INTO fs_auto.course (id,name,description,registration_date)
VALUES (
    2 --id
   ,'Nome 2' --name
   ,'Descricao 2' --description
   ,now() --registration_date
);


-- CLASS --
INSERT INTO fs_auto.class (id,course_id,start_date,end_date,registration_date)
VALUES (
    1 --id
   ,1 --course_id
   ,to_timestamp('20-01-2019 15:00:00', 'dd-mm-yyyy hh24:mi:ss') --start_date
   ,to_timestamp('25-01-2019 18:30:00', 'dd-mm-yyyy hh24:mi:ss') --end_date
   ,now() --registration_date
);
-----------
INSERT INTO fs_auto.class (id,course_id,start_date,end_date,registration_date)
VALUES (
    2 --id
   ,2 --course_id
   ,to_timestamp('28-01-2019 15:10:00', 'dd-mm-yyyy hh24:mi:ss') --start_date
   ,to_timestamp('05-02-2019 18:25:00', 'dd-mm-yyyy hh24:mi:ss') --end_date
   ,now() --registration_date
);


-- TEACHER --
INSERT INTO fs_auto.teacher (person_id,course_id)
VALUES (
    1 --person_id
   ,1 --course_id
);
-------------
INSERT INTO fs_auto.teacher (person_id,course_id)
VALUES (
    2 --person_id
   ,2 --course_id
);


-- STUDENT --
INSERT INTO fs_auto.student (person_id,class_id)
VALUES (
    3 --person_id
   ,1 --class_id
);
-------------
INSERT INTO fs_auto.student (person_id,class_id)
VALUES (
    4 --person_id
   ,1 --class_id
);
-------------
INSERT INTO fs_auto.student (person_id,class_id)
VALUES (
    5 --person_id
   ,2 --class_id
);
-------------
INSERT INTO fs_auto.student (person_id,class_id)
VALUES (
    6 --person_id
   ,2 --class_id
);
```

### End points

#### Course

- [GET] - Find All - <http://localhost:8080/course>
- [GET] - Find By ID - <http://localhost:8080/course/1>
- [POST] - Insert - <http://localhost:8080/course>

  ```json
  Body Request:
  {
    "name": "Name 3",
    "description": "Description 3"
  }
  ```

- [PUT] - Update - <http://localhost:8080/course>

  ```json
  Body Request:
  {
    "id": 3,
    "name": "New Name 3",
    "description": "New Description 3"
  }
  ```

- [DELETE] - Delete - <http://localhost:8080/course/3>

#### Class

- [GET] - Find All - <http://localhost:8080/class>
- [GET] - Find By ID - <http://localhost:8080/class/1>
- [POST] - Insert - <http://localhost:8080/class>

  ```json
  Body Request:
  {
    "course": {
        "id": 3,
        "registrationDate": "2019-10-20T15:00:00Z"
    },
    "startDate": "2019-10-20T15:00:00Z",
    "endDate": "2019-10-25T18:30:00Z"
  }
  ```

- [PUT] - Update - <http://localhost:8080/class>

  ```json
  Body Request:
  {
    "id": 3,
    "course": {
        "id": 2,
        "registrationDate": "2019-10-26T19:54:20.060092Z"
    },
    "startDate": "2019-10-20T15:00:00Z",
    "endDate": "2019-10-25T18:30:00Z"
  }
  ```

- [DELETE] - Delete - <http://localhost:8080/class/3>

#### Person

- [GET] - Find All - <http://localhost:8080/person>

#### Student

- [GET] - Find All - <http://localhost:8080/student>
- [GET] - Find By ID - <http://localhost:8080/student/1>
- [POST] - Insert - <http://localhost:8080/student>

  ```json
  Body Request:
  {
      "person": {
          "name": "Pessoa 7",
          "cpf": "23456789888",
          "cellPhone": "234567333",
          "city": "Cidade 7",
          "zipCode": "234566666",
          "address": "Endereco 7",
          "registrationDate": "2019-10-26T19:54:20.060092Z"
      },
      "class": {
          "id": 1,
          "course": {
              "registrationDate": "0001-01-01T00:00:00Z"
          },
          "startDate": "0001-01-01T00:00:00Z",
          "endDate": "0001-01-01T00:00:00Z",
          "registrationDate": "0001-01-01T00:00:00Z"
      }
  }
  ```

- [PUT] - Update - <http://localhost:8080/student>

  ```json
  Body Request:
  {
      "person": {
          "id": 3,
          "name": "Pessoa 3",
          "cpf": "23456789012",
          "cellPhone": "234567890",
          "city": "Cidade 3",
          "zipCode": "23456789",
          "address": "Endereco 3",
          "registrationDate": "2019-10-26T19:54:20.060092Z"
      },
      "class": {
          "id": 1,
          "course": {
              "registrationDate": "0001-01-01T00:00:00Z"
          },
          "startDate": "0001-01-01T00:00:00Z",
          "endDate": "0001-01-01T00:00:00Z",
          "registrationDate": "0001-01-01T00:00:00Z"
      }
  }
  ```

- [DELETE] - Delete - <http://localhost:8080/student/7>

#### Teacher

- [GET] - Find All - <http://localhost:8080/teacher>
- [GET] - Find By ID - <http://localhost:8080/teacher/1>
- [POST] - Insert - <http://localhost:8080/teacher>

  ```json
  Body Request:
  {
      "person": {
          "name": "Pessoa 7",
          "cpf": "56789012111",
          "cellPhone": "567890888",
          "city": "Cidade 7",
          "zipCode": "56789000",
          "address": "Endereco 7",
          "registrationDate": "2019-10-26T19:54:20.060092Z"
      },
      "class": {
          "id": 2,
          "course": {
              "registrationDate": "0001-01-01T00:00:00Z"
          },
          "startDate": "0001-01-01T00:00:00Z",
          "endDate": "0001-01-01T00:00:00Z",
          "registrationDate": "0001-01-01T00:00:00Z"
      }
  }
  ```

- [PUT] - Insert - <http://localhost:8080/teacher>

  ```json
  Body Request:
  {
      "person": {
          "id": 7
          "name": "New Pessoa 7",
          "cpf": "56789012111",
          "cellPhone": "567890888",
          "city": "New Cidade 7",
          "zipCode": "56789000",
          "address": "New Endereco 7",
          "registrationDate": "2019-10-30T19:54:20.060092Z"
      },
      "class": {
          "id": 2,
          "course": {
              "registrationDate": "0001-01-01T00:00:00Z"
          },
          "startDate": "0001-01-01T00:00:00Z",
          "endDate": "0001-01-01T00:00:00Z",
          "registrationDate": "0001-01-01T00:00:00Z"
      }
  }
  ```

- [DELETE] - Delete - <http://localhost:8080/student/7>
