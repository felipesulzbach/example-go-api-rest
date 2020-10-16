--------------------------------------------------------------------------------------------------------------------------------

INSERT INTO fs_auto.profile (id,name,code,access_permission)
VALUES (
    nextval('fs_auto.profile_seq') --id
   ,'NOPERMISSION' --name
   ,'DF001' --code
   ,'NO_ACCESS' --access_permission
);
---------------------------------------------------------------
INSERT INTO fs_auto.profile (id,name,code,access_permission)
VALUES (
    nextval('fs_auto.profile_seq') --id
   ,'DEFAULT' --name
   ,'DF002' --code
   ,'READING' --access_permission
);
---------------------------------------------------------------
INSERT INTO fs_auto.profile (id,name,code,access_permission)
VALUES (
    nextval('fs_auto.profile_seq') --id
   ,'ADMIN' --name
   ,'DF003' --code
   ,'WRITING' --access_permission
);

--------------------------------------------------------------------------------------------------------------------------------

INSERT INTO fs_auto.person (id,name,cpf,cell_phone,city,zip_code,address,registration_date)
VALUES (
    nextval('fs_auto.person_seq') --id
   ,'ADMIN' --name
   ,null --cpf
   ,null --cell_phone
   ,null --city
   ,null --zip_code
   ,null --address
   ,now() --registration_date
);

--------------------------------------------------------------------------------------------------------------------------------

INSERT INTO fs_auto.app_user (id,name,password,person_id,profile_id,start_date,end_date,registration_date)
VALUES (
    nextval('fs_auto.app_user_seq') --id
   ,'admin' --name
   ,'240BE518FABD2724DDB6F04EEB1DA5967448D7E831C08C8FA822809F74C720A9' --password (admin123)
   ,(SELECT id FROM fs_auto.person WHERE name = 'ADMIN') --person_id
   ,(SELECT id FROM fs_auto.profile WHERE code = 'DF003') --profile_id
   ,now() --start_date
   ,(now() + interval '30 day') --end_date
   ,now() --registration_date
);
---------------------------------------------------------------
INSERT INTO fs_auto.app_user (id,name,password,person_id,profile_id,start_date,end_date,registration_date)
VALUES (
    nextval('fs_auto.app_user_seq') --id
   ,'outher' --name
   ,'B62D5C4E8F6880A95882D373B8B96C90827A2C3F656842205CF7353B3FBC2F4F' --password (outher123)
   ,null --person_id
   ,(SELECT id FROM fs_auto.profile WHERE code = 'DF002') --profile_id
   ,now() --start_date
   ,(now() + interval '20 day') --end_date
   ,now() --registration_date
);

--------------------------------------------------------------------------------------------------------------------------------
