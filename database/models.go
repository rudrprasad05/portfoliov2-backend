package database

import "time"

type User struct {
	ID        *int       `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
}

const UserTable = `
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

type Image struct {
	ID        *int       `db:"id"`
	Link      string     `db:"link"`
	MediaID   *int       `db:"media_id"` // Foreign key to Media
	CreatedAt *time.Time `db:"created_at"`
}

const ImageTable = `
CREATE TABLE images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    link VARCHAR(255) NOT NULL UNIQUE,
    media_id INT UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE SET NULL
);`

type Media struct {
	ID        *int       `db:"id"`
	Name      string     `db:"name"`
	Alt       string     `db:"alt"`
	ImageID   *int       `db:"image_id"` // Foreign key to Image
	CreatedAt *time.Time `db:"created_at"`
}

const MediaTable = `
CREATE TABLE media (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    alt VARCHAR(255) NOT NULL,
    image_id INT UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE SET NULL
);`

type Tag struct {
	ID        *int       `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
}

const TagTable = `
CREATE TABLE tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

type Post struct {
	ID            *int       `db:"id"`
	Title         string     `db:"title"`
	Content       string     `db:"content"`
	CreatedAt     *time.Time `db:"created_at"`
	Tags          []*Tag
	FeaturedMedia *Media `db:"media_id"`
}

const PostTable = `
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    media_id INT, -- Foreign key to Media table
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE SET NULL
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
	ID          *int       `db:"id"`
	Type        *TypeInt   `db:"type"`
	Data        string     `db:"Data"`
	CreatedAt   *time.Time `db:"created_at"`
	IndexOnPage int        `db:"index_on_page"`
}

const content = `CREATE TABLE content (
    id INT AUTO_INCREMENT PRIMARY KEY,
    type INT NOT NULL CHECK (type IN (0, 1, 2, 3, 4, 5)),
    data TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    index_on_page INT NOT NULL
);`
