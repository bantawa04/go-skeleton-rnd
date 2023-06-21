ALTER TABLE posts
ADD COLUMN category_id INT NULL,
ADD CONSTRAINT fk_posts_categories
FOREIGN KEY (category_id)
REFERENCES categories (id)
ON DELETE CASCADE;