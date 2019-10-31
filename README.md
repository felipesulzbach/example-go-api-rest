# GoLang Web Service REST

API REST example in GoLang.

--------------------------------------------------------------------------------

## Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/) or other IDE (Integrated Development Environment);
- [Golang](https://golang.org/);

## Preparing the environment

The following technologies are critical for running/compiling application sources:

- Access the application directory from the terminal and execute:

  - Install the [gorilla/mux](https://github.com/gorilla/mux):

    > go get -u github.com/gorilla/mux

  - Install the [Postgres Driver](https://github.com/lib/pq):

    > go get github.com/lib/pq

  - Install the [go-i18n](https://github.com/nicksnyder/go-i18n)

    > go get -u github.com/nicksnyder/go-i18n/v2/goi18n

## PostgreSql Database configuration

### Connection Settings

- HOST: **localhost**
- PORT: **5435**
- USER: **postgres**
- PASSWORD: **postgres**
- DATABASE NAME: **go_rest_db**

### Create structure

```sql
DROP SCHEMA IF EXISTS GO_TST CASCADE;

CREATE SCHEMA GO_TST;

-- PERSON --
CREATE TABLE GO_TST.person
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
CREATE TABLE GO_TST.course
(id numeric NOT NULL
,name character varying(255)
,description character varying(255)
,registration_date timestamp
,CONSTRAINT pk_course PRIMARY KEY (id));

-- CLASS --
CREATE TABLE GO_TST.class
(id numeric NOT NULL
,course_id numeric NOT NULL
,start_date timestamp
,end_date timestamp
,registration_date timestamp
,CONSTRAINT pk_class PRIMARY KEY (id)
,CONSTRAINT fk_class_course FOREIGN KEY (course_id) REFERENCES GO_TST.course (id));
CREATE INDEX idx_class_course ON GO_TST.class(course_id);

-- STUDENT --
CREATE TABLE GO_TST.student
(person_id numeric NOT NULL
,class_id numeric NOT NULL
,CONSTRAINT pk_student PRIMARY KEY (person_id, class_id)
,CONSTRAINT fk_student_person FOREIGN KEY (person_id) REFERENCES GO_TST.person (id)
,CONSTRAINT fk_student_class FOREIGN KEY (class_id) REFERENCES GO_TST.class (id)
);

-- TEACHER --
CREATE TABLE GO_TST.teacher
(person_id numeric NOT NULL
,course_id numeric NOT NULL
,CONSTRAINT pk_teacher PRIMARY KEY (person_id, course_id)
,CONSTRAINT fk_teacher_person FOREIGN KEY (person_id) REFERENCES GO_TST.person (id)
,CONSTRAINT fk_teacher_course FOREIGN KEY (course_id) REFERENCES GO_TST.course (id)
);
```

### Mock registers

```sql
-- PERSON --
INSERT INTO GO_TST.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
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
INSERT INTO GO_TST.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
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
INSERT INTO GO_TST.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
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
INSERT INTO GO_TST.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
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
INSERT INTO GO_TST.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
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
INSERT INTO GO_TST.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
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
INSERT INTO GO_TST.course (id,name,description,registration_date)
VALUES (
    1 --id
   ,'Nome 1' --name
   ,'Descricao 1' --description
   ,now() --registration_date
);
------------
INSERT INTO GO_TST.course (id,name,description,registration_date)
VALUES (
    2 --id
   ,'Nome 2' --name
   ,'Descricao 2' --description
   ,now() --registration_date
);


-- CLASS --
INSERT INTO GO_TST.class (id,course_id,start_date,end_date,registration_date)
VALUES (
    1 --id
   ,1 --course_id
   ,to_timestamp('20-01-2019 15:00:00', 'dd-mm-yyyy hh24:mi:ss') --start_date
   ,to_timestamp('25-01-2019 18:30:00', 'dd-mm-yyyy hh24:mi:ss') --end_date
   ,now() --registration_date
);
-----------
INSERT INTO GO_TST.class (id,course_id,start_date,end_date,registration_date)
VALUES (
    2 --id
   ,2 --course_id
   ,to_timestamp('28-01-2019 15:10:00', 'dd-mm-yyyy hh24:mi:ss') --start_date
   ,to_timestamp('05-02-2019 18:25:00', 'dd-mm-yyyy hh24:mi:ss') --end_date
   ,now() --registration_date
);


-- TEACHER --
INSERT INTO GO_TST.teacher (person_id,course_id)
VALUES (
    1 --person_id
   ,1 --course_id
);
-------------
INSERT INTO GO_TST.teacher (person_id,course_id)
VALUES (
    2 --person_id
   ,2 --course_id
);


-- STUDENT --
INSERT INTO GO_TST.student (person_id,class_id)
VALUES (
    3 --person_id
   ,1 --class_id
);
-------------
INSERT INTO GO_TST.student (person_id,class_id)
VALUES (
    4 --person_id
   ,1 --class_id
);
-------------
INSERT INTO GO_TST.student (person_id,class_id)
VALUES (
    5 --person_id
   ,2 --class_id
);
-------------
INSERT INTO GO_TST.student (person_id,class_id)
VALUES (
    6 --person_id
   ,2 --class_id
);
```

## End points

### Course

- [GET] - Find All - <http://localhost:8080/course>
- [GET] - Find By ID - <http://localhost:8080/course/1>
- [POST] - Insert - <http://localhost:8080/course>

  ```json
  Body Request:
  {
    "id": 3,
    "name": "Name 3",
    "description": "Description 3",
    "registrationDate": "2019-10-30T19:54:20.060092Z"
  }
  ```

- [DELETE] - Delete - <http://localhost:8080/course/3>

### Class

- [GET] - Find All - <http://localhost:8080/class>
- [GET] - Find By ID - <http://localhost:8080/class/1>

### Person

- [GET] - Find All - <http://localhost:8080/person>

### Student

- [GET] - Find All - <http://localhost:8080/student>
- [GET] - Find By ID - <http://localhost:8080/student/1>

### Teacher

- [GET] - Find All - <http://localhost:8080/teacher>
- [GET] - Find By ID - <http://localhost:8080/teacher/1>
