package database

import "time"

type User struct {
	ID        *int       `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
}

const user = `
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255)
);
`

type Image struct {
	ID        *int   `db:"id"`
	Link      string `db:"link"`
	Media     *Media
	CreatedAt *time.Time `db:"created_at"`
}

const image = `CREATE TABLE images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    link VARCHAR(100) NOT NULL UNIQUE,
    media_id INT UNIQUE, -- Foreign key to Media table
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE SET NULL
`

type Media struct {
	ID        *int   `db:"id"`
	Name      string `db:"name"`
	Alt       string `db:"alt"`
	Image     *Image
	CreatedAt *time.Time `db:"created_at"`
}

const media = `CREATE TABLE media (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    alt VARCHAR(100) NOT NULL,
    image_id INT UNIQUE, -- Foreign key to Image table
    FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE SET NULL
`

type Tag struct {
	ID        *int       `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
}

const tag = `CREATE TABLE tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

type Post struct {
	ID            *int       `db:"id"`
	Title         string     `db:"title"`
	Content       string     `db:"content"`
	CreatedAt     *time.Time `db:"created_at"`
	Tags          []*Tag
	FeaturedImage *Image `db:"image"`
}

const post = `CREATE TABLE posts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

type PostTag struct {
	PostID int `db:"post_id"`
	TagID  int `db:"tag_id"`
}

const postTags = `CREATE TABLE post_tags (
    post_id INT,
    tag_id INT,
    PRIMARY KEY (post_id, tag_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);`

type Content struct {
	ID        *int       `db:"id"`
	Type      *TypeInt   `db:"type"`
	Content   string     `db:"content"`
	CreatedAt *time.Time `db:"created_at"`
	Index     int        `db:"index"`
}
