ALTER TABLE `posts`
ADD CONSTRAINT `fk_posts_categories`
FOREIGN KEY (`category_id`)
REFERENCES `categories` (`id`)
ON DELETE SET NULL;