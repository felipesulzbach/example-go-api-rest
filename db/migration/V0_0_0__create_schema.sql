--DROP SCHEMA IF EXISTS fs_auto CASCADE;
CREATE SCHEMA IF NOT EXISTS fs_auto;

---------- PERSON ----------
CREATE SEQUENCE fs_auto.person_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE fs_auto.person
(id NUMERIC NOT NULL
,name CHARACTER VARYING(255)
,cpf CHARACTER VARYING(20)
,cell_phone CHARACTER VARYING(20)
,city CHARACTER VARYING(255)
,zip_code CHARACTER VARYING(20)
,address CHARACTER VARYING(255)
,registration_date TIMESTAMP
,CONSTRAINT pk_person PRIMARY KEY (id));

ALTER TABLE fs_auto.person ALTER COLUMN id SET DEFAULT NEXTVAL('fs_auto.person_seq');
----------------------------

---------- COURSE ----------
CREATE SEQUENCE fs_auto.course_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE fs_auto.course
(id NUMERIC NOT NULL
,name CHARACTER VARYING(255)
,description CHARACTER VARYING(255)
,registration_date TIMESTAMP
,CONSTRAINT pk_course PRIMARY KEY (id));

ALTER TABLE fs_auto.course ALTER COLUMN id SET DEFAULT NEXTVAL('fs_auto.course_seq');
----------------------------

---------- SCHOOL_CLASS ----------
CREATE SEQUENCE fs_auto.school_class_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE fs_auto.school_class
(id NUMERIC NOT NULL
,course_id NUMERIC NOT NULL
,start_date TIMESTAMP
,end_date TIMESTAMP
,registration_date TIMESTAMP
,CONSTRAINT pk_school_class PRIMARY KEY (id)
,CONSTRAINT fk_school_class_course FOREIGN KEY (course_id) REFERENCES fs_auto.course (id));

CREATE INDEX idx_school_class_course ON fs_auto.school_class(course_id);

ALTER TABLE fs_auto.school_class ALTER COLUMN id SET DEFAULT NEXTVAL('fs_auto.school_class_seq');
---------------------------

---------- STUDENT ----------
CREATE TABLE fs_auto.student
(id NUMERIC NOT NULL
,school_class_id NUMERIC
,CONSTRAINT pk_student PRIMARY KEY (id)
,CONSTRAINT fk_student_person FOREIGN KEY (id) REFERENCES fs_auto.person (id)
,CONSTRAINT fk_student_school_class FOREIGN KEY (school_class_id) REFERENCES fs_auto.school_class (id));

CREATE INDEX idx_student_school_class ON fs_auto.student(school_class_id);
-----------------------------

---------- TEACHER ----------
CREATE TABLE fs_auto.teacher
(id NUMERIC NOT NULL
,course_id NUMERIC
,CONSTRAINT pk_teacher PRIMARY KEY (id)
,CONSTRAINT fk_teacher_person FOREIGN KEY (id) REFERENCES fs_auto.person (id)
,CONSTRAINT fk_teacher_course FOREIGN KEY (course_id) REFERENCES fs_auto.course (id));

CREATE INDEX idx_teacher_course ON fs_auto.teacher(course_id);
-----------------------------

---------- PROFILE ----------
CREATE SEQUENCE fs_auto.profile_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE fs_auto.profile
(id NUMERIC NOT NULL
,name CHARACTER VARYING(255)
,code CHARACTER VARYING(5)
,access_permission CHARACTER VARYING(50)
,CONSTRAINT pk_profile PRIMARY KEY (id));

ALTER TABLE fs_auto.profile ALTER COLUMN id SET DEFAULT NEXTVAL('fs_auto.profile_seq');
-----------------------------

---------- APP_USER ----------
CREATE SEQUENCE fs_auto.app_user_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE fs_auto.app_user
(id NUMERIC NOT NULL
,name CHARACTER VARYING(255)
,password CHARACTER VARYING(255)
,profile_id NUMERIC
,person_id NUMERIC
,start_date TIMESTAMP
,end_date TIMESTAMP
,registration_date TIMESTAMP
,CONSTRAINT pk_app_user PRIMARY KEY (id)
,CONSTRAINT fk_app_user_profile FOREIGN KEY (profile_id) REFERENCES fs_auto.profile (id)
,CONSTRAINT fk_app_user_person FOREIGN KEY (person_id) REFERENCES fs_auto.person (id));

CREATE INDEX idx_app_user_profile ON fs_auto.app_user(profile_id);
CREATE INDEX idx_app_user_person ON fs_auto.app_user(person_id);

ALTER TABLE fs_auto.app_user ALTER COLUMN id SET DEFAULT NEXTVAL('fs_auto.app_user_seq');
-----------------------------
