create table writer(
            author_id int not null,
            name text not null,
            nationality text not null,
            age int ,
            address text,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            update_at TIMESTAMP default null,
            deleted_at TIMESTAMP default null,
            created_by uuid default null,
            updated_by uuid default null,
            primary key (author_id)
);