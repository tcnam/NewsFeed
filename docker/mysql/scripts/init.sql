use mysql_newsfeed;
go
create table if not exists t_user(
	user_id int not null auto_increment,
	hash_password binary(20) not null,
	salt char(20),
	first_name varchar(255),
	last_name varchar(255),
	dob datetime,
	email varchar(255) unique not null,
	user_name varchar(255),
	created_at datetime,
	updated_at datetime
	
	primary key(`user_id`)
)ENGINE=INNODB;
go
create table if not exists t_user_user(
	fk_user_id int,
	fk_folower_id int,
	INDEX fk_user_id_ind (fk_user_id),
	INDEX fk_folower_id_ind (fk_folower_id),
	FOREIGN KEY (fk_user_id) 
		REFERENCES t_user(user_id),
	FOREIGN KEY (fk_folower_id)
		REFERENCES t_user(user_id)
)ENGINE=INNODB;
go
create table if not exists t_post(
	post_id int not null auto_increment primary key,
	fk_user_id int,
	content_text blob, 
	content_image_path varchar(255),
	create_at datetime, 
	updated_at datetime,
	INDEX fk_user_id_ind (fk_user_id),
	FOREIGN KEY(fk_user_id) 
		REFERENCES t_user(user_id)
)ENGINE=INNODB;
go
create table if not exists t_like (
	fk_post_id int,
	fk_user_id int,
	created_at datetime,
	updated_at datetime,
	index fk_post_id_ind (fk_post_id),
	index fk_user_id_ind (fk_user_id),
	foreign key(fk_user_id) 
		references t_user(user_id),
	foreign key(fk_post_id) 
		references t_post(post_id),
	unique key (fk_post_id, fk_user_id)
)ENGINE=INNODB;
go
create table if not exists t_comment (
	comment_id int not null auto_increment primary key,
	fk_post_id int,
	fk_user_id int,
	content blob,
	created_at datetime,
	updated_at datetime,
	index fk_post_id_ind (fk_post_id),
	index fk_user_id_ind (fk_user_id),
	foreign key(fk_user_id) references t_user(user_id),
	foreign key(fk_post_id) references t_post(post_id)
)ENGINE=INNODB;
